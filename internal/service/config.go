package service

import "github.com/spf13/viper"

type ConnectConf struct {
	AdbPath string `json:"adbPath"`
	Address string `json:"address"`
	// Config  string `json:"config"`
}

func init() {
	// 初始化默认配置
	viper.SetDefault("adbPath", "adb")
	viper.SetDefault("address", "192.168.123.234:5555")
	// viper.SetDefault("config", "General")
}

func GetMaaConnectConf() ConnectConf {
	var connectConf ConnectConf
	connectConf.AdbPath = viper.GetString("adbPath")
	connectConf.Address = viper.GetString("address")
	// connectConf.Config = viper.GetString("config")
	return connectConf
}

func SetMaaConnectConf(adbPath string, address string) ConnectConf {
	viper.Set("address", address)
	viper.Set("adbPath", adbPath)
	return GetMaaConnectConf()
}
