package main

import (
	"fmt"

	"github.com/JoelInman-Dev/tickwars/configs"
	"github.com/JoelInman-Dev/tickwars/pkg/utils"
)

func main() {
	// load configs
	cfg := configs.Load()

	// load logging
	LogFiles, err := utils.CreateLogFiles()
	if err != nil {
		panic(err)
	}
	// make sure to close them again after when we are done
	defer utils.CloseLogFiles(LogFiles)
	utils.InitLoggers(LogFiles)
	utils.PlayerLogger.Info("hello!", "foo", "bar", "fee", 0)
	//server := web.NewServer(cfg)
	fmt.Printf("App URL: %s \n", cfg.AppUrl)
	fmt.Printf("App Name: %s \n", cfg.AppName)
	fmt.Printf("DB Connection: %s \n", cfg.DBConn)
	fmt.Printf("Port: %s \n", cfg.Port)
	//game.StartGlobalLoop() // global game loop
	//server.Run()
}
