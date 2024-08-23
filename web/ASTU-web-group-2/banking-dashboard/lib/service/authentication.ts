import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const authenticationApi = createApi({
  reducerPath: "authentication",
  baseQuery: fetchBaseQuery({
    baseUrl: "https://astu-bank-dashboard.onrender.com",
  }),

  endpoints: (builder) => ({
    refreshAccessToken: builder.mutation({
      query: (refreshToken: string) => ({
        url: "/auth/refresh_token",
        method: "POST",
        headers: {
          Authorization: `Bearer ${refreshToken}`,
        },
      }),
    }),
  }),
});

export const { useRefreshAccessTokenMutation } = authenticationApi;
