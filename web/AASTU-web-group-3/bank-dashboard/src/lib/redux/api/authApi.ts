import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

interface Credentials {
  userName: string;
  password: string;
}

interface AuthResponse {
  data: {
    access_token: string;
    refresh_token: string;
    userId: string;
  };
  success: boolean;
}

export const authApi = createApi({
  reducerPath: 'authApi',
  baseQuery: fetchBaseQuery({
    baseUrl: process.env.NEXT_PUBLIC_BASE_URL,
  }),
  endpoints: (builder) => ({
    login: builder.mutation<AuthResponse, Credentials>({
      query: (credentials) => ({
        url: '/auth/login',
        method: 'POST',
        body: credentials,
        headers: { 'Content-Type': 'application/json' },
      }),
    }),
  }),
});

export const { useLoginMutation } = authApi;