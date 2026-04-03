package business

import (
	interfaces "NumbersManagmentService/internal/business/intrefaces"
	"NumbersManagmentService/internal/domain"
	"context"
	"go.uber.org/zap"
)

type PhoneService struct {
	repo   interfaces.PhoneRepository
	logger *zap.Logger
}

func NewPhoneService(repo interfaces.PhoneRepository, logger *zap.Logger) *PhoneService {
	return &PhoneService{
		repo:   repo,
		logger: logger,
	}
}

func (s *PhoneService) Import(
	ctx context.Context,
	cmd domain.ImportCommand,
) (domain.ImportResult, error) {

	var result domain.ImportResult

	var toInsert []domain.PhoneNumber

	seen := make(map[string]struct{})

	for _, raw := range cmd.Numbers {

		number, err := NormalizePhone(raw)
		if err != nil {
			result.Errors++
			continue
		}

		if _, ok := seen[number]; ok {
			result.Duplicates++
			continue
		}

		exists, err := s.repo.Exists(ctx, number)
		if err != nil {
			return result, err
		}

		if exists {
			result.Duplicates++
			continue
		}

		seen[number] = struct{}{}

		country := DetectCountry(number)

		region, provider := "", ""

		if country == "RU" {
			region, provider = DetectRUProvider(number)
		}

		toInsert = append(toInsert, domain.PhoneNumber{
			Number:   number,
			Country:  country,
			Region:   region,
			Provider: provider,
			Source:   cmd.Source,
		})

		result.Accepted++
	}

	if len(toInsert) > 0 {
		err := s.repo.InsertBatch(ctx, toInsert)
		if err != nil {
			return result, err
		}
	}

	return result, nil
}

func (s *PhoneService) Search(
	ctx context.Context,
	q domain.SearchQuery,
) (domain.SearchResult, error) {

	items, total, err := s.repo.Search(ctx, q)
	if err != nil {
		return domain.SearchResult{}, err
	}

	return domain.SearchResult{
		Numbers: items,
		Total:   total,
	}, nil
}
