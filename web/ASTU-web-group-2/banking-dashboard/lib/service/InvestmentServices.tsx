import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const investmentApi = createApi({
  reducerPath: "bankdashboard",
  baseQuery: fetchBaseQuery({
    baseUrl: "https://bank-dashboard-6acc.onrender.com/user/random-investment-data", 
  }),
  endpoints: (builder) => ({
    signUp: builder.mutation({
      query: (data) => ({
        url: "/auth/register",
        method: "POST",
        body: data,
      }),
    }),
    getInvestment: builder.query({
      query: (data : {accessToken: string, years: number, months: number}) => ({
        // Construct the URL with parameters
         url : `/year/${data.years}/months/${data.months}`,
          method: "GET",
          headers: {
            Authorization: `Bearer ${data.accessToken}`, 
        },
      }),
    }),
  }),
});

export const { useGetInvestmentQuery } = investmentApi;