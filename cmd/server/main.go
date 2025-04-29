package main

import (
	"github.com/JoelInman-Dev/tickwars/configs"
	"github.com/JoelInman-Dev/tickwars/internal/web"
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

	// initial boot logs
	utils.PlayerLogger.Info("hello!", "foo", "bar", "fee", 0)
	utils.EventLogger.Info("App URL Loaded", "APP_URL", cfg.AppUrl)
	utils.EventLogger.Info("App Name Loaded", "APP_NAME", cfg.AppName)
	utils.EventLogger.Info("DB Connection established", "DBCONN", cfg.DBConn)
	utils.EventLogger.Info("Port Loaded", "PORT", cfg.Port)

	// Load the API Server and MUX
	server := web.NewServer()

	// start the game ticker
	//game.StartGlobalLoop() // global game loop

	// Go Motherfucker!
	server.Run(cfg.AppUrl, cfg.Port)
}
