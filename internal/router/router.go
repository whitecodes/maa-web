package router

import (
	"maa-web/internal/handler"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/version", handler.GetVersion)
}
