package main

import (
	"fmt"
	"go_server/config"
	"go_server/router"
)

func main() {
	config.InitConfig()
	router.InitRouter()
	fmt.Print(config.AppConfig.App.Port)
}
