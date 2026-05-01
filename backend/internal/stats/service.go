package stats

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

func (s *Service) Get(ctx context.Context, filters StatsFilters) (*StatsResponse, error) {
	totalOrders, err := s.repo.CountTotalOrders(ctx, filters)
	if err != nil {
		return nil, err
	}

	errorOrders, err := s.repo.CountErrorOrders(ctx, filters)
	if err != nil {
		return nil, err
	}

	byChannel, err := s.repo.CountByChannel(ctx, filters)
	if err != nil {
		return nil, err
	}

	byFulfillmentType, err := s.repo.CountByFulfillmentType(ctx, filters)
	if err != nil {
		return nil, err
	}

	byProductType, err := s.repo.CountByProductType(ctx, filters)
	if err != nil {
		return nil, err
	}

	errorPercentage := calculateErrorPercentage(errorOrders, totalOrders)

	return &StatsResponse{
		TotalOrders:       totalOrders,
		ErrorPercentage:   errorPercentage,
		ByChannel:         byChannel,
		ByFulfillmentType: byFulfillmentType,
		ByProductType:     byProductType,
	}, nil
}

func calculateErrorPercentage(errorOrders int64, totalOrders int64) float64 {
	if totalOrders == 0 {
		return 0
	}

	value := (float64(errorOrders) / float64(totalOrders)) * 100

	return math.Round(value*100) / 100
}
