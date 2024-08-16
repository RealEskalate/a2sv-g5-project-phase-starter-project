import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import { session } from '@/session';
import { get } from 'http';

const size = 5;

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
    getAllTransactions: builder.query<void, string>({
      query: (page) => `/transactions?page=${page}&size=${size}`,
    }),
    getTransactionById: builder.query<void, string>({
      query: (id) => `/transactions/${id}`,
    }),
    getTransactionIncome: builder.query<void, string>({
      query: (page) => `/transactions/incomes?page=${page}&size=${size}`,
    }),
    getTransactionExpense: builder.query<void, string>({
      query: (page) => `/transactions/expense?page=${page}&size=${size}`,
    }),
    postDeposit: builder.mutation<
      void,
      { amount: number; description: string; type: string; receiverUserName: string }
    >({
      query: ({ amount, description, type, receiverUserName }) => ({
        url: '/transactions/deposit',
        method: 'POST',
        body: { amount, description, type, receiverUserName },
      }),
    }),
  }),
});

export const {
  useGetAllTransactionsQuery,
  useGetTransactionByIdQuery,
  useGetTransactionIncomeQuery,
  useGetTransactionExpenseQuery,
  usePostDepositMutation,
} = transactionApi;
