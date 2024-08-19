import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { getSession } from "next-auth/react";
import {
  TransactionsResponse,
  TransactionRequest,
  TransactionResponse,
  TransactionDepositRequest
} from "../types/transactions";

export const transactionsApi = createApi({
  reducerPath: "transactionsApi",
  baseQuery: fetchBaseQuery({
    baseUrl: "https://bank-dashboard-6acc.onrender.com",
    prepareHeaders: async (headers) => {
      const session = await getSession();
      const token = session?.accessToken;

      if (token) {
        headers.set("Authorization", `Bearer ${token}`);
      }
      return headers;
    },
  }),
  endpoints: (builder) => ({
    getAllTransactions: builder.query<
      TransactionsResponse,
      { size: number; page: number }
    >({
      query: ({ size, page }) => `/transactions?size=${size}&page=${page}`,
    }),

    getTransactionById: builder.query<TransactionResponse, string>({
      query: (id) => `/transactions/${id}`,
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
    TransactionDepositRequest,
    TransactionResponse
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
  }),
});

export const {
  useGetAllTransactionsQuery,
  useGetTransactionByIdQuery, 
  useCreateTransactionMutation,
  useCreateTransactionDepositMutation,
} = transactionsApi;
