package main

import (
	"context"

	"applicationDesignTest/internal/app"
	"applicationDesignTest/internal/logger"
)

func main() {
	log := logger.NewLogger()

	server := app.NewApp(log)

	if err := server.Run(context.Background()); err != nil {
		log.Errorf("server: error:", err)
	}
	log.Info("server: gracefully stopped")
}
