import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const loanApi = createApi({
  reducerPath: "loanDash",
  baseQuery: fetchBaseQuery({
    baseUrl: "https://astu-bank-dashboard.onrender.com",
  }),

  endpoints: (builder) => ({
    getMyLoanService: builder.query({
      query: (accessToken: string) => ({
        url: "/active-loans/my-loans",
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      }),
    }),
    getMyLoansDetail: builder.query({
      query: (accessToken: string) => ({
        url: "/active-loans/detail-data",
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      }),
    }),

    // Add your requests here the same way
  }),
});

export const { useGetMyLoanServiceQuery, useGetMyLoansDetailQuery } = loanApi;
