package inmemory

import (
	interfaces "NumbersManagmentService/internal/business/intrefaces"
	"NumbersManagmentService/internal/domain"
	"context"
	"go.uber.org/zap"
	"sync"
)

type InMemoryRepo struct {
	mu     sync.RWMutex
	db     map[string]domain.PhoneNumber
	logger *zap.Logger
}

func NewInMemoryRepo(logger *zap.Logger) interfaces.PhoneRepository {
	return &InMemoryRepo{
		db:     make(map[string]domain.PhoneNumber),
		logger: logger,
	}
}

func (r *InMemoryRepo) Exists(ctx context.Context, number string) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, ok := r.db[number]
	return ok, nil
}

func (r *InMemoryRepo) InsertBatch(ctx context.Context, numbers []domain.PhoneNumber) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, n := range numbers {
		r.db[n.Number] = n
	}

	return nil
}

func (r *InMemoryRepo) Search(
	ctx context.Context,
	q domain.SearchQuery,
) ([]domain.PhoneNumber, int, error) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []domain.PhoneNumber

	for _, n := range r.db {

		if q.Number != "" && !contains(n.Number, q.Number) {
			continue
		}

		if q.Country != "" && n.Country != q.Country {
			continue
		}

		if q.Region != "" && n.Region != q.Region {
			continue
		}

		if q.Provider != "" && n.Provider != q.Provider {
			continue
		}

		result = append(result, n)
	}

	total := len(result)

	start := q.Offset
	if start > total {
		return []domain.PhoneNumber{}, total, nil
	}

	end := start + q.Limit
	if end > total {
		end = total
	}

	return result[start:end], total, nil
}

func contains(str, substr string) bool {
	return substr == "" || (len(substr) <= len(str) &&
		containsSimple(str, substr))
}

func containsSimple(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

func (r *InMemoryRepo) Close() error {
	return nil
}
