package orders

import (
	"context"
	"math"
)

type Service struct {
	repo RepositoryPort
}

func NewService(repo RepositoryPort) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) List(ctx context.Context, filters ListOrdersFilters) (*ListOrdersResponse, error) {
	if filters.Page <= 0 {
		filters.Page = 1
	}

	if filters.PageSize <= 0 {
		filters.PageSize = 20
	}

	if filters.PageSize > 100 {
		filters.PageSize = 100
	}

	result, total, err := s.repo.List(ctx, filters)
	if err != nil {
		return nil, err
	}

	data := OrdersToDTOs(result)

	totalPages := 0
	if total > 0 {
		totalPages = int(math.Ceil(float64(total) / float64(filters.PageSize)))
	}

	return &ListOrdersResponse{
		Data: data,
		Pagination: Pagination{
			Page:       filters.Page,
			PageSize:   filters.PageSize,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}

func (s *Service) GetFilters(ctx context.Context) (*OrderFiltersOptions, error) {
	return s.repo.GetFilters(ctx)
}