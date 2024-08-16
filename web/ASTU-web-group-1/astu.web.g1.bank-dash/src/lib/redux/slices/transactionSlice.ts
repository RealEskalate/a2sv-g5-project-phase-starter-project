import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query';
import { session } from '@/session';

export const transactionApi = createApi({
  reducerPath: 'transactionApi',
  baseQuery: fetchBaseQuery({
    baseUrl: 'https://bank-dashboard-6acc.onrender.com',
    prepareHeaders: (headers, { getState }) => {
      // You can add custom headers here
      headers.set('Authorization', `Bearer ${session}`);
      headers.set('Content-Type', 'application/json');
      return headers;
    },
  }),
  endpoints: (builder) => ({
    getAllTransactions: builder.query<void, void>({
      query: () => '/transactions',
    }),
  }),
});

export const { useGetAllTransactions } = transactionApi;
