package postgres

import (
	interfaces "NumbersManagmentService/internal/business/intrefaces"
	"NumbersManagmentService/internal/config"
	"NumbersManagmentService/internal/domain"
	"NumbersManagmentService/internal/migrations"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

func buildDSN(cfg config.DBConfig) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)
}

func NewPostgresRepo(cfg *config.PostgresConfig, logger *zap.Logger) interfaces.PhoneRepository {
	logger.Info("MASTER DB CONFIG",
		zap.String("host", cfg.Master.Host),
		zap.Int("port", cfg.Master.Port),
		zap.String("user", cfg.Master.User),
	)

	masterDSN := buildDSN(cfg.Master)
	slaveDSN := buildDSN(cfg.Slave)

	masterDB, err := sql.Open("postgres", masterDSN)
	if err != nil {
		logger.Fatal("failed to open master db", zap.Error(err))
	}

	if err := masterDB.Ping(); err != nil {
		logger.Fatal("failed to ping master db", zap.Error(err))
	}

	slaveDB, err := sql.Open("postgres", slaveDSN)
	if err != nil {
		logger.Fatal("failed to open slave db", zap.Error(err))
	}

	if err := slaveDB.Ping(); err != nil {
		logger.Fatal("failed to ping slave db", zap.Error(err))
	}

	logger.Info("connected to postgres",
		zap.String("master", cfg.Master.Host),
		zap.String("slave", cfg.Slave.Host),
	)

	err = migrations.Run(masterDB, cfg, logger)
	if err != nil {
		logger.Fatal("failed to apply migrations", zap.Error(err))
	}

	return &PostgresRepo{
		master: masterDB,
		slave:  slaveDB,
	}
}

type PostgresRepo struct {
	master *sql.DB
	slave  *sql.DB
}

func (pr *PostgresRepo) Exists(ctx context.Context, number string) (bool, error) {
	const query = `SELECT 1 FROM phone_numbers WHERE number = $1 LIMIT 1`

	var tmp int
	err := pr.master.QueryRowContext(ctx, query, number).Scan(&tmp)

	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
func (pr *PostgresRepo) InsertBatch(ctx context.Context, numbers []domain.PhoneNumber) error {

	if len(numbers) == 0 {
		return nil
	}

	tx, err := pr.master.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	const query = `
	INSERT INTO phone_numbers (number, country, region, provider, source)
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (number) DO NOTHING
	`

	for _, n := range numbers {
		_, err := tx.ExecContext(ctx, query,
			n.Number,
			n.Country,
			n.Region,
			n.Provider,
			n.Source,
		)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (pr *PostgresRepo) Search(
	ctx context.Context,
	q domain.SearchQuery,
) ([]domain.PhoneNumber, int, error) {

	args := make([]interface{}, 0)
	where := "WHERE 1=1"
	argPos := 1

	if q.Number != "" {
		where += fmt.Sprintf(" AND number LIKE $%d", argPos)
		args = append(args, "%"+q.Number+"%")
		argPos++
	}

	if q.Country != "" {
		where += fmt.Sprintf(" AND country = $%d", argPos)
		args = append(args, q.Country)
		argPos++
	}

	if q.Region != "" {
		where += fmt.Sprintf(" AND region = $%d", argPos)
		args = append(args, q.Region)
		argPos++
	}

	if q.Provider != "" {
		where += fmt.Sprintf(" AND provider = $%d", argPos)
		args = append(args, q.Provider)
		argPos++
	}

	countQuery := "SELECT COUNT(*) FROM phone_numbers " + where

	var total int
	err := pr.slave.QueryRowContext(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	dataQuery := fmt.Sprintf(`
		SELECT number, country, region, provider, source
		FROM phone_numbers
		%s
		ORDER BY number
		LIMIT $%d OFFSET $%d
	`, where, argPos, argPos+1)

	args = append(args, q.Limit, q.Offset)

	rows, err := pr.slave.QueryContext(ctx, dataQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	result := make([]domain.PhoneNumber, 0)

	for rows.Next() {
		var n domain.PhoneNumber

		if err := rows.Scan(
			&n.Number,
			&n.Country,
			&n.Region,
			&n.Provider,
			&n.Source,
		); err != nil {
			return nil, 0, err
		}

		result = append(result, n)
	}

	return result, total, nil
}

func (pr *PostgresRepo) Close() error {
	if err := pr.master.Close(); err != nil {
		return err
	}

	if pr.slave != pr.master {
		return pr.slave.Close()
	}

	return nil
}
