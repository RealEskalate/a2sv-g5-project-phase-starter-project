import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
const SERVER_URL = process.env.NEXT_PUBLIC_SERVER_URL;

if (!SERVER_URL) throw Error("SERVER_URL is undefined");

export const transactionAPI = createApi({
  reducerPath: "transactions",
  baseQuery: fetchBaseQuery({ baseUrl: SERVER_URL }),
  endpoints: (builder) => ({
    // TODO : This is sample endpoint REMOVE IT 🫵🫵🫵!
    getAllTransactions: builder.query({
      query: (accessToken: string) => ({
        url: `/transactions`,
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${accessToken}`,
        },
      }),
    }),
  }),
});

export const { useGetAllTransactionsQuery } = transactionAPI;
