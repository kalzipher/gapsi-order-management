package stats

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

func (r *Repository) CountTotalOrders(ctx context.Context, filters StatsFilters) (int64, error) {
	var total int64

	query := r.baseQuery(ctx, filters)

	if err := query.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (r *Repository) CountErrorOrders(ctx context.Context, filters StatsFilters) (int64, error) {
	var total int64

	query := r.baseQuery(ctx, filters).
		Where("(NULLIF(TRIM(error_code), '') IS NOT NULL OR NULLIF(TRIM(error_message), '') IS NOT NULL)")

	if err := query.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (r *Repository) CountByChannel(ctx context.Context, filters StatsFilters) ([]StatItem, error) {
	return r.countBy(ctx, filters, "canal", "No canal")
}

func (r *Repository) CountByFulfillmentType(ctx context.Context, filters StatsFilters) ([]StatItem, error) {
	return r.countBy(ctx, filters, "fulfillment_type", "No fulllfillment type")
}

func (r *Repository) CountByProductType(ctx context.Context, filters StatsFilters) ([]StatItem, error) {
	return r.countBy(ctx, filters, "product_type", "No product type")
}

func (r *Repository) baseQuery(ctx context.Context, filters StatsFilters) *gorm.DB {
	query := r.db.WithContext(ctx).Table(r.tableName)

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

	return query
}

func (r *Repository) countBy(
	ctx context.Context,
	filters StatsFilters,
	column string,
	emptyLabel string,
) ([]StatItem, error) {
	var result []StatItem

	query := r.baseQuery(ctx, filters).
		Select(
			"COALESCE(NULLIF(TRIM("+column+"), ''), ?) AS name, COUNT(*) AS total",
			emptyLabel,
		).
		Group("name").
		Order("total DESC")

	if err := query.Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
