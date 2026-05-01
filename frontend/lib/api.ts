import axios from "axios";

import { clearSession, getAccessToken, getRefreshToken, saveSession } from "@/lib/auth";
import { config } from "@/lib/config";
import type {
  LoginApiResponse,
  LoginRequest,
  LoginResponse,
  RefreshApiResponse,
} from "@/types/auth";
import { mapLoginResponse, mapRefreshResponse } from "@/lib/mappers/auth";
import { mapStatsResponse } from "@/lib/mappers/stats";
import type { OrdersFilters, OrdersResponse, OrdersApiResponse, OrderFiltersOptionsApi, OrderFiltersOptions } from "@/types/orders";
import { mapOrdersResponse, mapOrderFiltersOptions } from "@/lib/mappers/orders";
import type { StatsApiResponse, StatsFilters, StatsResponse } from "@/types/stats";


export const api = axios.create({
  baseURL: config.apiUrl,
  headers: {
    "Content-Type": "application/json",
  },
});

api.interceptors.request.use((request) => {
  const token = getAccessToken();

  if (token) {
    request.headers.Authorization = `Bearer ${token}`;
  }

  return request;
});

api.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error?.config;

    if (
      error?.response?.status === 401 &&
      !originalRequest._retry &&
      !originalRequest.url?.includes("/auth/login") &&
      !originalRequest.url?.includes("/auth/refresh")
    ) {
      originalRequest._retry = true;

      const refreshToken = getRefreshToken();

      if (!refreshToken) {
        clearSession();
        window.location.href = "/login";
        return Promise.reject(error);
      }

      try {
        const { data } = await axios.post<RefreshApiResponse>(
          `${config.apiUrl}/auth/refresh`,
          { refreshToken },
          {
            headers: {
              "Content-Type": "application/json",
            },
          }
        );

        const mappedData = mapRefreshResponse(data);

        const currentUser = localStorage.getItem("user");

        saveSession({
          accessToken: mappedData.accessToken,
          refreshToken,
          user: currentUser ? JSON.parse(currentUser) : { id: "", email: "", name: "" },
        });

        originalRequest.headers.Authorization = `Bearer ${mappedData.accessToken}`;

        return api(originalRequest);
      } catch {
        clearSession();
        window.location.href = "/login";
      }
    }

    return Promise.reject(error);
  }
);

export async function login(payload: LoginRequest): Promise<LoginResponse> {
  const { data } = await api.post<LoginApiResponse>("/auth/login", payload);
  return mapLoginResponse(data);
}

export async function logout() {
  const refreshToken = getRefreshToken();

  if (!refreshToken) return;

  await api.post("/auth/logout", {
    refresh_token: refreshToken,
  });

  clearSession();
}

export async function getOrders(
  filters: OrdersFilters
): Promise<OrdersResponse> {
  const { data } = await api.get<OrdersApiResponse>("/orders", {
    params: filters,
  });

  return mapOrdersResponse(data);
}

export async function getStats(
  filters: StatsFilters
): Promise<StatsResponse> {
  const { data } = await api.get<StatsApiResponse>("/stats", {
    params: filters,
  });

  return mapStatsResponse(data);
}

export async function getOrderFilters(): Promise<OrderFiltersOptions> {
  const { data } = await api.get<OrderFiltersOptionsApi>("/orders/filters");

  return mapOrderFiltersOptions(data);
}