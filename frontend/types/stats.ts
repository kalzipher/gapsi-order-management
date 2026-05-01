export type StatItemApi = {
  name: string;
  total: number;
};

export type StatsApiResponse = {
  total_orders: number;
  error_percentage: number;
  by_channel: StatItemApi[];
  by_fulfillment_type: StatItemApi[];
  by_product_type: StatItemApi[];
};

export type StatItem = {
  name: string;
  total: number;
};

export type StatsResponse = {
  totalOrders: number;
  errorPercentage: number;
  byChannel: StatItem[];
  byFulfillmentType: StatItem[];
  byProductType: StatItem[];
};

export type StatsFilters = {
  canal?: string;
  company?: string;
  fulfillmentType?: string;
  productType?: string;
};