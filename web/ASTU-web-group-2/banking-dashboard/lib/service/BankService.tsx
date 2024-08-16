import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const bankApi = createApi({
  reducerPath: "bankDash",
  baseQuery: fetchBaseQuery({
    baseUrl: "https://bank-dashboard-6acc.onrender.com",
  }),
  endpoints: (builder) => ({
    getBankService: builder.query({
      query: (accessToken: string) => ({
        url: "/bank-services",
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      }),
    }),

    // Add your requests here the same way
  }),
});

export const { useGetBankServiceQuery } = bankApi;
