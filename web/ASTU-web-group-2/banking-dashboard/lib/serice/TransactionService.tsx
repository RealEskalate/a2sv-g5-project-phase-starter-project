import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const transactionApi = createApi({
  reducerPath: "bankdashboard",
  baseQuery: fetchBaseQuery({
    baseUrl: "https://bank-dashboard.6acc.onrender.com",
  }),
  endpoints: (builder) => ({
    getAllTransaction: builder.query({
      query: (accessToken: string = "") => ({
        url: "/transactions",
        method: "GET",
        headers: {
          Authorizations: `Bearer ${accessToken}`,
        },
      }),
    }),
  }),
});

export const { useGetAllTransactionQuery } = transactionApi;
