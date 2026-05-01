export type User = {
  id: string;
  email: string;
  name: string;
};

export type LoginRequest = {
  email: string;
  password: string;
};

export type LoginApiResponse = {
  access_token: string;
  refresh_token: string;
  user: User;
};

export type LoginResponse = {
  accessToken: string;
  refreshToken: string;
  user: User;
};

export type RefreshApiResponse = {
  access_token: string;
};

export type RefreshResponse = {
  accessToken: string;
};