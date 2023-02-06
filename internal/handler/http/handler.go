package http

import (
	"github.com/labstack/echo/v4"

	"delivery/internal/config"
	v1 "delivery/internal/handler/http/v1"
	"delivery/internal/service"
)

func Init(services *service.Services, cfg *config.Config, e *echo.Echo) {
	v1.Init(services, cfg, e)
}
