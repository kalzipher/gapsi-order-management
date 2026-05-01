import type {
  LoginApiResponse,
  LoginResponse,
  RefreshApiResponse,
  RefreshResponse,
} from "@/types/auth";

export function mapLoginResponse(data: LoginApiResponse): LoginResponse {
  return {
    accessToken: data.access_token,
    refreshToken: data.refresh_token,
    user: data.user,
  };
}

export function mapRefreshResponse(data: RefreshApiResponse): RefreshResponse {
  return {
    accessToken: data.access_token,
  };
}