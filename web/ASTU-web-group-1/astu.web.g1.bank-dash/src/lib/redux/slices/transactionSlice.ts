import { createApi } from '@reduxjs/toolkit/query/react';
import { baseQuery } from '../api/baseQuery';
import { TransactionResponseType } from '@/types/transaction.types';
import { pageSize } from '@/types/page-size.type';

export const transactionApi = createApi({
  reducerPath: 'transactionApi',
  baseQuery: baseQuery(),
  endpoints: (builder) => ({
    getAllTransactions: builder.query<TransactionResponseType, pageSize>({
      query: (page) => `/transactions?page=${page.page}&size=${page.size}`,
    }),
    getTransactionById: builder.query<void, string>({
      query: (id) => `/transactions/${id}`,
    }),
    getTransactionIncome: builder.query<TransactionResponseType, pageSize>({
      query: (page) => `/transactions/incomes?page=${page.page}&size=${page.size}`,
    }),
    getTransactionExpense: builder.query<TransactionResponseType, pageSize>({
      query: (page) => `/transactions/expenses?page=${page.page}&size=${page.size}`,
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
