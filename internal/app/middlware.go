package app

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"time"
)

func loggingMiddleware(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {

		start := time.Now()

		err := c.Next()

		latency := time.Since(start)

		logger.Info("http request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("latency", latency),
			zap.String("ip", c.IP()),
		)

		return err
	}
}
