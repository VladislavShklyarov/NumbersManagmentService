package main

import (
	"NumbersManagmentService/internal/app"
	"NumbersManagmentService/internal/config"
	"NumbersManagmentService/internal/pkg/logger"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	zapLogger, err := logger.New(cfg.Env)
	if err != nil {
		log.Fatal(err)
	}

	app := app.NewApp(cfg, zapLogger)

	zapLogger.Info("service started",
		zap.String("env", cfg.Env),
	)

	log.Fatal(app.Run())
}
