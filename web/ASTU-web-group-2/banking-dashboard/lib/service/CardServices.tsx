import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const CreditCardInfoApi = createApi({
  reducerPath: "creditCard",
  baseQuery: fetchBaseQuery({
    baseUrl: "https://bank-dashboard-6acc.onrender.com",
  }),
  endpoints: (builder) => ({
    getAllCardInfo: builder.query({
      query: (accessToken: string = "") => ({
        url: "/cards",
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
    getAllCards: builder.query({
      query: (data: { token: string }) => ({
        url: `/cards`,
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
  useGetAllCardsQuery,
  useLazyRetiriveCardInfoQuery
} = CreditCardInfoApi;
