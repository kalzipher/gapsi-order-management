import type { StatsApiResponse, StatsResponse } from "@/types/stats";

export function mapStatsResponse(data: StatsApiResponse): StatsResponse {
  return {
    totalOrders: data.total_orders ?? 0,
    errorPercentage: data.error_percentage ?? 0,
    byChannel: data.by_channel ?? [],
    byFulfillmentType: data.by_fulfillment_type ?? [],
    byProductType: data.by_product_type ?? [],
  };
}