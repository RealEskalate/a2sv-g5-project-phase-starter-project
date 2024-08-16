import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const companyApi = createApi({
  reducerPath: "companyDash",
  baseQuery: fetchBaseQuery({
    baseUrl: "https://bank-dashboard-6acc.onrender.com",
  }),
  endpoints: (builder) => ({
    getCompanies: builder.query({
      query: (accessToken: string) => ({
        url: "/companies",
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      }),
    }),

    // Add your requests here the same way
  }),
});

export const { useGetCompaniesQuery } = companyApi;
