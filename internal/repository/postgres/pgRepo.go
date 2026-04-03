package postgres

import (
	interfaces "NumbersManagmentService/internal/business/intrefaces"
	"NumbersManagmentService/internal/config"
	"NumbersManagmentService/internal/domain"
	"context"
	"go.uber.org/zap"
)

func NewPostgresRepo(config config.PostgresConfig, logger *zap.Logger) interfaces.PhoneRepository {
	return &PgRepo{}
}

type PgRepo struct{}

func (pr *PgRepo) Exists(ctx context.Context, number string) (bool, error) {
	return false, nil
}

func (pr *PgRepo) InsertBatch(ctx context.Context, numbers []domain.PhoneNumber) error {
	return nil
}

func (pr *PgRepo) Search(ctx context.Context, q domain.SearchQuery) ([]domain.PhoneNumber, int, error) {
	return make([]domain.PhoneNumber, 0), 0, nil
}
