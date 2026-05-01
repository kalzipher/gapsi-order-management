export type OrderApi = {
  id: string;
  no_pedido: string;
  canal: string;
  sku: string;
  fecha_estimada: string;
  fulfillment_type: string;
  product_type: string;
  cantidad: number | null;
  fecha_compra: string | null;
  company: string;
  has_error: boolean;
};

export type PaginationApi = {
  page: number;
  page_size: number;
  total: number;
  total_pages: number;
};

export type OrdersApiResponse = {
  data: OrderApi[];
  pagination: PaginationApi;
};

export type Order = {
  id: string;
  noPedido: string;
  canal: string;
  sku: string;
  fechaEstimada: string;
  fulfillmentType: string;
  productType: string;
  cantidad: number | null;
  fechaCompra: string | null;
  company: string;
  hasError: boolean;
};

export type Pagination = {
  page: number;
  pageSize: number;
  total: number;
  totalPages: number;
};

export type OrdersResponse = {
  data: Order[];
  pagination: Pagination;
};

export type OrdersFilters = {
  page: number;
  pageSize: number;
  canal?: string;
  company?: string;
  fulfillmentType?: string;
  productType?: string;
};

export type OrderFiltersOptionsApi = {
  channels: string[];
  companies: string[];
  fulfillment_types: string[];
  product_types: string[];
};

export type OrderFiltersOptions = {
  channels: string[];
  companies: string[];
  fulfillmentTypes: string[];
  productTypes: string[];
};