package stats

type StatsFilters struct {
	Canal           string
	Company         string
	FulfillmentType string
	ProductType     string
}

type StatItem struct {
	Name  string `json:"name"`
	Total int64  `json:"total"`
}

type StatsResponse struct {
	TotalOrders       int64      `json:"total_orders"`
	ErrorPercentage   float64    `json:"error_percentage"`
	ByChannel         []StatItem `json:"by_channel"`
	ByFulfillmentType []StatItem `json:"by_fulfillment_type"`
	ByProductType     []StatItem `json:"by_product_type"`
}
