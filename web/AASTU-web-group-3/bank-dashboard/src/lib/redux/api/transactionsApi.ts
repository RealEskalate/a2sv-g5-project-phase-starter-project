import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { getSession } from "next-auth/react";
import {
  TransactionsResponse,
  TransactionRequest,
  TransactionResponse,
  TransactionDepositRequest,
  LatestTransferResponse,
  BalanceHistoryResponse,
  MyExpenseResponse,
  IncomeResponse,
  getQuickTransfersResponse,
} from "../types/transactions";

export const transactionsApi = createApi({
  reducerPath: "transactionsApi",
  baseQuery: fetchBaseQuery({
    baseUrl: process.env.NEXT_PUBLIC_BASE_URL,
    prepareHeaders: async (headers) => {
      const session = await getSession();
      const token = session?.accessToken;
      console.log("token form rtk query", token);

      if (token) {
        headers.set("Authorization", `Bearer ${token}`);
      }
      return headers;
    },
  }),
  endpoints: (builder) => ({
    getAllTransactions: builder.query<
      MyExpenseResponse,
      { size: number; page: number }
    >({
      query: ({ size, page }) => `/transactions?page=${page}&size=${size}`,
    }),

    getTransactionById: builder.query<TransactionResponse, string>({
      query: (id) => `/transactions/${id}`,
    }),

    getRandomBalanceHistory: builder.query<
      any,
      { monthsBeforeFirstTransaction: number }
    >({
      query: ({ monthsBeforeFirstTransaction }) =>
        `/transactions/random-balance-history?monthsBeforeFirstTransaction=${monthsBeforeFirstTransaction}`,
    }),

    getIncomeTransactions: builder.query<
      IncomeResponse,
      { size: number; page: number }
    >({
      query: ({ size, page }) =>
        `/transactions/incomes?page=${page}&size=${size}`,
    }),

    getExpenseTransactions: builder.query<
      MyExpenseResponse,
      { size: number; page: number }
    >({
      query: ({ size, page }) =>
        `/transactions/expenses?page=${page}&size=${size}`,
    }),

    createTransaction: builder.mutation<
      TransactionResponse,
      TransactionRequest
    >({
      query: (transaction) => ({
        url: "/transactions",
        method: "POST",
        body: transaction,
        headers: {
          "Content-Type": "application/json",
        },
      }),
    }),

    createTransactionDeposit: builder.mutation<
      TransactionResponse,
      TransactionDepositRequest
    >({
      query: (deposit) => ({
        url: "/transactions/deposit",
        method: "POST",
        body: deposit,
        headers: {
          "Content-Type": "application/json",
        },
      }),
    }),
    getLatestTransfer: builder.query<LatestTransferResponse, { num: number }>({
      query: ({ num }) => `/transactions/latest-transfers${num}`,
    }),
    getBalanceHistory: builder.query<BalanceHistoryResponse, {}>({
      query: () => "/transactions/balance-history",
    }),
    getQuickTransfers: builder.query<
      getQuickTransfersResponse,
      { num: number }
    >({
      query: ({ num }) => `/transactions/quick-transfers?number=${num}`,
    }),
  }),
});

export const {
  useGetAllTransactionsQuery,
  useGetTransactionByIdQuery,
  useGetRandomBalanceHistoryQuery,
  useGetIncomeTransactionsQuery,
  useGetExpenseTransactionsQuery,
  useGetLatestTransferQuery,
  useCreateTransactionMutation,
  useCreateTransactionDepositMutation,
  useGetBalanceHistoryQuery,
  useGetQuickTransfersQuery,
} = transactionsApi;
