package server

import (
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"

	"delivery/pkg/logger"
)

const (
	_defaultTimeout = 2 * time.Second
)

type Server struct {
	GRPCServer      *grpc.Server
	notify          chan error
	shutdownTimeout time.Duration
	logger          logger.Interface
}

func (s *Server) Shutdown() error {
	s.GRPCServer.GracefulStop()
	return nil
}

func (s *Server) Notify() chan error {
	return s.notify
}

func New(addr string, l logger.Interface) (*Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", addr))
	if err != nil {
		l.Error("failed to listen: %v", err)
		return nil, err
	}
	s := grpc.NewServer()
	serv := &Server{
		GRPCServer:      s,
		notify:          make(chan error),
		shutdownTimeout: _defaultTimeout,
		logger:          l,
	}
	l.Info("grpc server listening at %v", lis.Addr())
	go func() {
		if err = s.Serve(lis); err != nil {
			serv.notify <- err
			close(serv.notify)
			l.Fatal("grpc failed to serve: %v", err)
		}
	}()

	return serv, nil
}
