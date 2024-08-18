import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import { getSession } from 'next-auth/react';
import { TransactionsResponse } from '../types/transactions'; // Import or define your types

export const transactionsApi = createApi({
  reducerPath: 'transactionsApi',
  baseQuery: fetchBaseQuery({
    baseUrl: 'https://bank-dashboard-6acc.onrender.com',
    prepareHeaders: async (headers) => {
      const session = await getSession();
      const token = session?.accessToken; // Get the access token from the session
      console.log("Token", token)



      if (token) {
        headers.set('Authorization', `Bearer ${token}`);
      }
      return headers;
    },
  }),
  endpoints: (builder) => ({
    getAllTransactions: builder.query<TransactionsResponse, { size: number; page: number }>({
      query: ({ size, page }) => `/transactions`,
      
    }),
  }),
});

export const { useGetAllTransactionsQuery } = transactionsApi;
