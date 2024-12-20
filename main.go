package main

import (
	"fmt"
	"go_server/config"
	"go_server/router"
)

func main() {
	fmt.Println("Initializing application...") // 提前打印初始化信息
	config.InitConfig()
	fmt.Println("App start...", config.AppConfig.App.Port)
	router.InitRouter()
}
