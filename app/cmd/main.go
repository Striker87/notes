package main

import (
	"fmt"

	"github.com/Striker87/notes/internal/app"
	"github.com/Striker87/notes/internal/config"
	"github.com/Striker87/notes/pkg/logging"
)

// swag init -g ./app/cmd/main.go -o ./app/docs
func main() {
	fmt.Println("config init")
	cfg := config.NewConfig()

	fmt.Println("logger init")
	logging.Init(cfg.AppConfig.LogLevel)
	logger := logging.GetLogger()

	a, err := app.NewApp(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("Running application")
	a.Run()
}
