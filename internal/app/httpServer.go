package app

import (
	"NumbersManagmentService/internal/config"
	"NumbersManagmentService/internal/transport"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type HTTPServer struct {
	app    *fiber.App
	host   string
	port   int
	logger *zap.Logger
}

func NewHTTPServer(handler *transport.Handler, cfg config.HTTPConfig, logger *zap.Logger) *HTTPServer {

	app := fiber.New()

	app.Use(loggingMiddleware(logger))

	api := app.Group("/api")
	numbers := api.Group("/numbers")

	numbers.Post("/import", handler.Import)
	numbers.Get("/search", handler.Search)

	return &HTTPServer{
		app:    app,
		host:   cfg.Host,
		port:   cfg.Port,
		logger: logger,
	}
}

func (s *HTTPServer) Run() error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	return s.app.Listen(addr)
}
