package orders

import "time"

type Order struct {
	ID              string
	Canal           string
	Cantidad        *int
	Company         string
	CP              string
	CreatedAt       *time.Time
	DaysToDelivery  *int
	ErrorCode       string
	ErrorMessage    string
	FechaCompra     *time.Time
	FechaEstimada   string
	FulfillmentType string
	IsFlash         *bool
	IsMarketplace   *bool
	NoPedido        string
	Plan            string
	ProductType     string
	SKU             string
	StoreSelected   string
	TipoPago        string
	EDD1            string
	EDD2            string
}

type OrderDTO struct {
	ID              string     `json:"id"`
	NoPedido        string     `json:"no_requested"`
	Canal           string     `json:"channel"`
	SKU             string     `json:"sku"`
	FechaEstimada   string     `json:"date_of_estimated_delivery"`
	FulfillmentType string     `json:"fulfillment_type"`
	ProductType     string     `json:"product_type"`
	Cantidad        *int       `json:"quantity"`
	FechaCompra     *time.Time `json:"date_of_purchase"`
	Company         string     `json:"company"`
	HasError        bool       `json:"has_error"`
}

type ListOrdersFilters struct {
	Page            int
	PageSize        int
	Canal           string
	Company         string
	FulfillmentType string
	ProductType     string
}

type Pagination struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

type ListOrdersResponse struct {
	Data       []OrderDTO `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type OrderFiltersOptions struct {
	Channels         []string `json:"channels"`
	Companies        []string `json:"companies"`
	FulfillmentTypes []string `json:"fulfillment_types"`
	ProductTypes     []string `json:"product_types"`
}
