package app

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"delivery/internal/config"
	"delivery/internal/data"
	"delivery/internal/handler/grpc"
	httphandler "delivery/internal/handler/http"
	"delivery/internal/model/proto"
	"delivery/internal/repo"
	"delivery/internal/service"
	"delivery/pkg/database/postgres"
	"delivery/pkg/grpc/server"
	"delivery/pkg/httpserver"
	"delivery/pkg/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	go data.Init(pg.DB, cfg)

	boil.DebugMode = cfg.Debug

	// Repository
	repos := repo.NewRepositories(pg.DB)
	// Service
	services := service.NewServices(&service.Deps{
		Repos: repos,
		Cfg:   cfg,
	})

	// grpc
	grpcServer, err := server.New(cfg.GRPC.Port, l)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - grpcServer - server.New: %w", err))
	}
	deliveryServer := grpc.NewDeliveryServer(services)
	proto.RegisterDeliveryServer(grpcServer.GRPCServer, deliveryServer)

	// HTTP Server
	e := echo.New()
	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	}))
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{"*"},
	}))
	e.Use(middleware.Recover())

	httphandler.Init(services, cfg, e)

	// init route
	httpServer := httpserver.New(e, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	case err = <-grpcServer.Notify():
		l.Error(fmt.Errorf("app - Run - grpcServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	err = grpcServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - grpcServer.Shutdown: %w", err))
	}
}
