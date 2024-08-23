import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const transactionApi = createApi({
  reducerPath: "transactionDash",
  baseQuery: fetchBaseQuery({
    baseUrl: "https://astu-bank-dashboard.onrender.com",
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
    getExpenses: builder.query({
      query: (accessToken: string) => ({
        url: "/transactions/expenses",
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
        params: {
          page: 1,
          size: 7,
        },
      }),
    }),
    getBalanceHistory: builder.query({
      query: (accessToken: string) => ({
        url: "/transactions/balance-history",
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      }),
    }),
    getInvestmentHistory: builder.query({
      query: (accessToken: string) => ({
        url: "/user/random-investment-data",
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
        params: {
          years: 10,
          months: 0,
        },
      }),
    }),
  }),
});

export const {
  useGetAllTransactionQuery,
  useSignUpMutation,
  useGetExpensesQuery,
  useGetBalanceHistoryQuery,
  useGetInvestmentHistoryQuery,
} = transactionApi;
