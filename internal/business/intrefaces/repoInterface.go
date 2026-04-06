package interfaces

import (
	"NumbersManagmentService/internal/domain"
	"context"
	"io"
)

type PhoneRepository interface {
	Exists(ctx context.Context, number string) (bool, error)
	InsertBatch(ctx context.Context, numbers []domain.PhoneNumber) error
	Search(ctx context.Context, q domain.SearchQuery) ([]domain.PhoneNumber, int, error)
	io.Closer
}
