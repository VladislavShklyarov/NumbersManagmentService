package mapper

import (
	"NumbersManagmentService/internal/domain"
	"NumbersManagmentService/internal/transport/requests"
	"NumbersManagmentService/internal/transport/responses"
)

func ToImportCommand(req requests.ImportRequest) domain.ImportCommand {
	return domain.ImportCommand{
		Numbers: req.Numbers,
		Source:  req.Source,
	}
}

func ToSearchQuery(req requests.SearchRequest) domain.SearchQuery {
	limit := req.Limit
	if limit <= 0 || limit > 100 {
		limit = 50
	}

	offset := req.Offset
	if offset < 0 {
		offset = 0
	}

	return domain.SearchQuery{
		Number:   req.Number,
		Country:  req.Country,
		Region:   req.Region,
		Provider: req.Provider,
		Limit:    limit,
		Offset:   offset,
	}
}

func ToSearchResponse(res domain.SearchResult) responses.SearchResponse {
	dto := make([]responses.PhoneDTO, len(res.Numbers))

	for i, n := range res.Numbers {
		dto[i] = responses.PhoneDTO{
			Number:   n.Number,
			Country:  n.Country,
			Region:   n.Region,
			Provider: n.Provider,
			Source:   n.Source,
		}
	}

	return responses.SearchResponse{
		Total:   res.Total,
		Results: dto,
	}
}

func ToImportResponse(res domain.ImportResult) responses.ImportResponse {
	return responses.ImportResponse{
		Accepted:   res.Accepted,
		Duplicates: res.Duplicates,
		Errors:     res.Errors,
	}
}
