package handler

import (
	"maa-web/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetVersion(c echo.Context) error {
	maa_version := service.GetMaaVersion()
	return c.JSON(http.StatusOK, map[string]string{"version": maa_version})
}

func GetMaaConnected(c echo.Context) error {
	maaHandle := service.GetMaaHandle()
	maa_connected := service.GetMaaConnected(maaHandle)
	return c.JSON(http.StatusOK, map[string]bool{"connected": maa_connected})
}
