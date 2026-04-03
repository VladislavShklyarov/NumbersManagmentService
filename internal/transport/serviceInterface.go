package transport

import (
	"NumbersManagmentService/internal/domain"
	"context"
)

type CommandService interface {
	Import(ctx context.Context, cmd domain.ImportCommand) (domain.ImportResult, error)
}

type QueryService interface {
	Search(ctx context.Context, q domain.SearchQuery) (domain.SearchResult, error)
}
