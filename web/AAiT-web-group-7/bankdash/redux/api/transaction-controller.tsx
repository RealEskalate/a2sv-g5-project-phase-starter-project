import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

export const transactionApi = createApi({
  reducerPath: 'transactionApi',
  baseQuery: fetchBaseQuery({
    baseUrl: "https://bank-dashboard-latest.onrender.com",
    prepareHeaders: (headers) => {
      const token = localStorage.getItem('token');
      if (token) {
        headers.set('Authorization', `Bearer ${token}`);
      }
      return headers;
    },
  }),

  tagTypes: ['transactions'],
  endpoints: (builder) => ({

    getTransactions: builder.query({
      query: ({page, size}) => ({
        url: `/transactions?page=${page}&size=${size}`,
        method: 'GET',
      }),
      providesTags: ['transactions'],
    }),

    postTransaction: builder.mutation({
      query: ({amount, type, description, receiverUserName}) => ({
        url: '/transactions',
        method: 'POST',
        body: { amount, type, description, receiverUserName },
      }),
      invalidatesTags: ['transactions'],
    }),

    postTransactionDeposit: builder.mutation({
      query: ({amount, description}) => ({
        url: '/transactions/deposit',
        method: 'POST',
        body: { amount, description },
      }),
      invalidatesTags: ['transactions'],
    }),

    getTransaction: builder.query({
      query: ({id}) => ({
        url: `/transactions/${id}`,
        method: 'GET',
      }),
      providesTags: ['transactions'],
    }),

    getTransactionBalanceHistory: builder.query({
      query: ({monthsBeforeFirstTransaction}) => ({
        url: `/transactions/random-balance-history?monthsBeforeFirstTransaction=${monthsBeforeFirstTransaction}`,
        method: 'GET',
      }),
      providesTags: ['transactions'],
    }),

    transactionQuickTransfer: builder.query({
      query: ({query}) => ({
        url: `/transactions/quick-transfers?number=${query}`,
        method: 'GET',
      }),
      providesTags: ['transactions'],
    }),

    getTransactionsIncomes: builder.query({
      query: ({page, size}) => ({
        url: `/transactions/incomes?page=${page}&size=${size}`,
        method: 'GET',
      }),
      providesTags: ['transactions'],
    }),

    getTransactionsExpenses: builder.query({
      query: ({page, size}) => ({
        url: `/transactions/expenses?page=${page}&size=${size}`,
        method: 'GET',
      }),
      providesTags: ['transactions'],
    }),

    getTransactionsExpenseSummary: builder.query({
      query: ({startDate, type}) => ({
        url: `/transactions/expenses-summary?startDate=${startDate}&type=${type}`,
        method: 'GET',
      }),
      providesTags: ['transactions'],
    }),

    getTransactionsBalanceHistory: builder.query({
      query: () => ({
        url: '/transactions/balance-history',
        method: 'GET',
      }),
      providesTags: ['transactions'],
    })
    ,
  }),
});

export const {
  useGetTransactionsQuery,
  usePostTransactionMutation,
  usePostTransactionDepositMutation,
  useGetTransactionQuery,
  useGetTransactionBalanceHistoryQuery,
  useTransactionQuickTransferQuery,
  useGetTransactionsIncomesQuery,
  useGetTransactionsExpensesQuery,
  useGetTransactionsExpenseSummaryQuery,
  useGetTransactionsBalanceHistoryQuery,
} = transactionApi;



