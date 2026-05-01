package orders

import "context"

type ServicePort interface {
	List(ctx context.Context, filters ListOrdersFilters) (*ListOrdersResponse, error)
	GetFilters(ctx context.Context) (*OrderFiltersOptions, error)
}

type RepositoryPort interface {
	List(ctx context.Context, filters ListOrdersFilters) ([]Order, int64, error)
	GetFilters(ctx context.Context) (*OrderFiltersOptions, error)
}
