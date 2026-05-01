import type {
  Order,
  OrderApi,
  OrderFiltersOptions,
  OrderFiltersOptionsApi,
  OrdersApiResponse,
  OrdersResponse,
} from "@/types/orders";

export function mapOrder(order: OrderApi) {
  return {
    id: order.id,
    noPedido: order.no_pedido,
    canal: order.canal,
    sku: order.sku,
    fechaEstimada: order.fecha_estimada,
    fulfillmentType: order.fulfillment_type,
    productType: order.product_type,
    cantidad: order.cantidad,
    fechaCompra: order.fecha_compra,
    company: order.company,
    hasError: order.has_error,
  };
}

export function mapOrdersResponse(data: OrdersApiResponse): OrdersResponse {
  return {
    data: data.data.map(mapOrder),
    pagination: {
      page: data.pagination.page,
      pageSize: data.pagination.page_size,
      total: data.pagination.total,
      totalPages: data.pagination.total_pages,
    },
  };
}

export function mapOrderFiltersOptions(
  data: OrderFiltersOptionsApi
): OrderFiltersOptions {
  return {
    channels: data.channels ?? [],
    companies: data.companies ?? [],
    fulfillmentTypes: data.fulfillment_types ?? [],
    productTypes: data.product_types ?? [],
  };
}