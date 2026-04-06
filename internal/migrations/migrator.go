package migrations

import (
	"NumbersManagmentService/internal/config"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

func Run(
	db *sql.DB,
	cfg *config.PostgresConfig,
	logger *zap.Logger,
) error {

	if !cfg.Migrations.Enabled {
		logger.Info("migrations are disabled")
		return nil
	}

	_, err := os.ReadDir(cfg.Migrations.Path)
	if err != nil {
		return fmt.Errorf("failed to read migrations folder: %w", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{MigrationsTable: "schema_migrations"})
	if err != nil {
		return fmt.Errorf("create postgres migration driver failed: %w", err)
	}

	absPath, err := filepath.Abs(cfg.Migrations.Path)
	if err != nil {
		return fmt.Errorf("failed to get absolute path for migrations: %w", err)
	}
	sourceURL := fmt.Sprintf("file://%s", absPath)

	logger.Info("using migrations path", zap.String("absPath", absPath), zap.String("sourceURL", sourceURL))

	m, err := migrate.NewWithDatabaseInstance(
		sourceURL,
		"postgres",
		driver)
	if err != nil {
		return fmt.Errorf("create migrate instance failed: %w", err)
	}

	logger.Info("running database migrations", zap.String("path", cfg.Migrations.Path))

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logger.Info("no new migrations to apply")
			return nil
		}
		return fmt.Errorf("apply migrations failed: %w", err)
	}

	logger.Info("database migrations applied successfully")
	return nil
}
