package app

import (
	"NumbersManagmentService/internal/business"
	interfaces "NumbersManagmentService/internal/business/intrefaces"
	"NumbersManagmentService/internal/config"
	"NumbersManagmentService/internal/repository/inmemory"
	"NumbersManagmentService/internal/repository/postgres"
	"NumbersManagmentService/internal/transport"
	"go.uber.org/zap"
)

type App struct {
	httpServer *HTTPServer
	logger     *zap.Logger
}

func NewApp(cfg *config.Config, logger *zap.Logger) *App {

	// repo
	repo := buildRepo(cfg, logger)

	// services
	phoneService := business.NewPhoneService(repo, logger)

	// handler
	handler := transport.NewHandler(phoneService, phoneService, logger)

	// server
	server := NewHTTPServer(handler, cfg.HTTP, logger)

	return &App{
		httpServer: server,
	}
}

func (a *App) Run() error {
	return a.httpServer.Run()
}

func buildRepo(cfg *config.Config, logger *zap.Logger) interfaces.PhoneRepository {

	if cfg.Storage.Type == "postgres" {
		return postgres.NewPostgresRepo(cfg.Storage.Postgres, logger)
	}

	return inmemory.NewInMemoryRepo(logger)
}
