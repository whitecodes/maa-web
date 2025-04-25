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
	maaConnected := service.GetMaaConnected(maaHandle)
	return c.JSON(http.StatusOK, map[string]bool{"connected": maaConnected})
}

func GetMaaConnectConf(c echo.Context) error {

	return c.JSON(http.StatusOK, service.GetMaaConnectConf())
}

func SetMaaConnectConf(c echo.Context) error {
	connectConf := new(service.ConnectConf)
	err := c.Bind(connectConf)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, service.SetMaaConnectConf(connectConf.AdbPath, connectConf.Address))
}

func TestConnect(c echo.Context) error {
	connectConf := new(service.ConnectConf)
	err := c.Bind(connectConf)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// todo 资源路径应该是配置文件中配置的
	err = service.LoadMaaResource("/home/white/code/maa-web/lib")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}

	maaHandle := service.GetMaaHandle()

	// err = service.ConnectDevice(maaHandle, connectConf.AdbPath, connectConf.Address, connectConf.Config)
	err = service.ConnectDevice(maaHandle, connectConf.AdbPath, connectConf.Address, "General")
	defer service.DestroyMaaHandle(maaHandle)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": err.Error()})
	}
	// 怎么判断连接成功？
	// if !service.BackToHome(maaHandle) {
	// 	return c.JSON(http.StatusOK, map[string]string{"status": "fail"})
	// }

	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}
