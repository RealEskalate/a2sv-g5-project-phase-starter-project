import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const transactionApi = createApi({
  reducerPath: "bankdashboard",
  baseQuery: fetchBaseQuery({
    baseUrl: "https://bank-dashboard-6acc.onrender.com",
  }),
  endpoints: (builder) => ({
    signUp: builder.mutation({
      query: (data) => ({
        url: "/auth/register",
        method: "POST",
        body: data,
      }),
    }),
    getAllTransaction: builder.query({
      query: (accessToken: string) => ({
        url: "/transactions",
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      }),
    }),
  }),
});

export const { useGetAllTransactionQuery, useSignUpMutation } = transactionApi;
