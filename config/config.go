package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")   //指定配置文件名
	viper.SetConfigType("yml")      //指定文件类型
	viper.AddConfigPath("./config") //指定路径

	//ReadInConfig 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	AppConfig = &Config{} //执行config实例，如果没有这行AppConfig则会打印nil

	//将读取的配置数据解析到 AppConfig 结构体中
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decode into  struct: %v", err)
	}
}
