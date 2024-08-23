import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const CreditCardInfoApi = createApi({
  reducerPath: "creditCard",
  baseQuery: fetchBaseQuery({
    baseUrl: "https://astu-bank-dashboard.onrender.com",
  }),
  endpoints: (builder) => ({
    getAllCardInfo: builder.query({
      query: (accessToken : string) => ({
        url: `/cards?page=0&size=${4}`,
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      }),
    }),
    AddCreditCard: builder.mutation({
      query: (data: { accessToken: string; passcode: string }) => ({
        url: "/cards",
        method: "POST",
        headers: {
          Authorization: `Bearer ${data.accessToken}`,
        },
        body: data,
      }),
    }),
    retiriveCardInfo: builder.query({
      query: (data: { token: string; id: string }) => ({
        url: `/cards/${data.id}`,
        method: "GET",
        headers: {
          Authorization: `Bearer ${data.token}`,
        },
      }),
    }),
  }),
});

export const {
  useGetAllCardInfoQuery,
  useAddCreditCardMutation,
  useRetiriveCardInfoQuery,
  useLazyRetiriveCardInfoQuery
} = CreditCardInfoApi;
