package service

import "github.com/spf13/viper"

type ConnectConf struct {
	adbPath string
	Address string `json:"address"`
	config  string
}

func init() {
	// 初始化默认配置
	viper.SetDefault("adbPath", "adb")
	viper.SetDefault("address", "192.168.123.234:5555")
	viper.SetDefault("config", "General")
}

func GetMaaConnectConf() ConnectConf {
	var connectConf ConnectConf
	connectConf.adbPath = viper.GetString("adbPath")
	connectConf.Address = viper.GetString("address")
	connectConf.config = viper.GetString("config")
	return connectConf
}
