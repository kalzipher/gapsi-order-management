package orders

import (
	"time"

	"github.com/github/gapsi-order-management-dashboard/backend/internal/database"
)

type OrderEntity struct {
	ID              string     `gorm:"column:id;primaryKey"`
	Canal           string     `gorm:"column:canal"`
	Cantidad        *int       `gorm:"column:cantidad"`
	Company         string     `gorm:"column:company"`
	CP              string     `gorm:"column:cp"`
	CreatedAt       *time.Time `gorm:"column:created_at"`
	DaysToDelivery  *int       `gorm:"column:days_to_delivery"`
	ErrorCode       string     `gorm:"column:error_code"`
	ErrorMessage    string     `gorm:"column:error_message"`
	FechaCompra     *time.Time `gorm:"column:fecha_compra"`
	FechaEstimada   string     `gorm:"column:fecha_estimada"`
	FulfillmentType string     `gorm:"column:fulfillment_type"`
	IsFlash         *bool      `gorm:"column:is_flash"`
	IsMarketplace   *bool      `gorm:"column:is_marketplace"`
	NoPedido        string     `gorm:"column:no_pedido"`
	Plan            string     `gorm:"column:plan"`
	ProductType     string     `gorm:"column:product_type"`
	SKU             string     `gorm:"column:sku"`
	StoreSelected   string     `gorm:"column:store_selected"`
	TipoPago        string     `gorm:"column:tipo_pago"`
	EDD1            string     `gorm:"column:edd1"`
	EDD2            string     `gorm:"column:edd2"`
}

func (OrderEntity) TableName() string {
	return database.TableNameOrders
}
