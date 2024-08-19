import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const transactionApi = createApi({
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
    getAllTransaction: builder.query({
      query: (accessToken: string, years: number, months: number) => {
        // Construct the URL with parameters
        const url = `/year/${years}/months/${months}`; 
        return {
          url, 
          method: "GET",
          headers: {
            Authorization: `Bearer ${accessToken}`, 
          },
        };
      },
    }),
  }),
});

export const { useGetAllTransactionQuery, useSignUpMutation } = transactionApi;