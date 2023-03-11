package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var Config *viper.Viper

func init() {
	Config = viper.New()
	//获取当前工作目录
	dir, err := os.Getwd()
	fmt.Println(`读取当前工作目录下的配置文件：` + dir)
	if err != nil {
		panic(err)
	}
	Config.AddConfigPath(dir)
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	if err = Config.ReadInConfig(); err != nil {
		panic(err)
	}
}
