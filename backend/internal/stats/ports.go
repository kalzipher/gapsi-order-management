package stats

import "context"

type ServicePort interface {
	Get(ctx context.Context, filters StatsFilters) (*StatsResponse, error)
}

type RepositoryPort interface {
	CountTotalOrders(ctx context.Context, filters StatsFilters) (int64, error)
	CountErrorOrders(ctx context.Context, filters StatsFilters) (int64, error)
	CountByChannel(ctx context.Context, filters StatsFilters) ([]StatItem, error)
	CountByFulfillmentType(ctx context.Context, filters StatsFilters) ([]StatItem, error)
	CountByProductType(ctx context.Context, filters StatsFilters) ([]StatItem, error)
}
