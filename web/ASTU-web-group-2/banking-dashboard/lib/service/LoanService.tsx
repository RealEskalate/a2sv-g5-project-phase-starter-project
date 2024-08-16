import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const loanApi = createApi({
  reducerPath: "loanDash",
  baseQuery: fetchBaseQuery({
    baseUrl: "https://bank-dashboard-6acc.onrender.com",
  }),

  endpoints: (builder) => ({
    getAllLoanService: builder.query({
      query: (accessToken: string) => ({
        url: "/active-loans/all",
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      }),
    }),

    // Add your requests here the same way
  }),
});

export const { useGetAllLoanServiceQuery } = loanApi;
