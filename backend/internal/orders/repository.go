package orders

import (
	"context"

	"github.com/github/gapsi-order-management-dashboard/backend/internal/database"
	"gorm.io/gorm"
)

type Repository struct {
	db        *gorm.DB
	tableName string
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db:        db,
		tableName: database.TableNameOrders,
	}
}

func (r *Repository) List(ctx context.Context, filters ListOrdersFilters) ([]Order, int64, error) {
	var entities []OrderEntity
	var total int64

	query := r.db.WithContext(ctx).Model(&OrderEntity{})

	if filters.Canal != "" {
		query = query.Where("canal = ?", filters.Canal)
	}

	if filters.Company != "" {
		query = query.Where("company = ?", filters.Company)
	}

	if filters.FulfillmentType != "" {
		query = query.Where("fulfillment_type = ?", filters.FulfillmentType)
	}

	if filters.ProductType != "" {
		query = query.Where("product_type = ?", filters.ProductType)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (filters.Page - 1) * filters.PageSize

	err := query.
		Order("created_at DESC NULLS LAST").
		Limit(filters.PageSize).
		Offset(offset).
		Find(&entities).
		Error

	if err != nil {
		return nil, 0, err
	}

	return EntitiesToOrders(entities), total, nil
}

func (r *Repository) GetFilters(ctx context.Context) (*OrderFiltersOptions, error) {
	channels, err := r.getDistinctValues(ctx, "canal")
	if err != nil {
		return nil, err
	}

	companies, err := r.getDistinctValues(ctx, "company")
	if err != nil {
		return nil, err
	}

	fulfillmentTypes, err := r.getDistinctValues(ctx, "fulfillment_type")
	if err != nil {
		return nil, err
	}

	productTypes, err := r.getDistinctValues(ctx, "product_type")
	if err != nil {
		return nil, err
	}

	return &OrderFiltersOptions{
		Channels:         channels,
		Companies:        companies,
		FulfillmentTypes: fulfillmentTypes,
		ProductTypes:     productTypes,
	}, nil
}

func (r *Repository) getDistinctValues(ctx context.Context, column string) ([]string, error) {
	var values []string

	err := r.db.WithContext(ctx).
		Table(r.tableName).
		Select(column).
		Where(column+" IS NOT NULL").
		Where("TRIM("+column+") <> ''").
		Group(column).
		Order(column+" ASC").
		Pluck(column, &values).
		Error

	if err != nil {
		return nil, err
	}

	return values, nil
}
