package v1

import (
	"github.com/labstack/echo/v4"

	"delivery/internal/config"
	"delivery/internal/service"
)

func Init(services *service.Services, cfg *config.Config, e *echo.Echo) {
	g := e.Group("/v1")

	NewWebhook(services.Delivery).init(g)
}
