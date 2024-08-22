import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const CreditCardInfoApi = createApi({
  reducerPath: "creditCard",
  baseQuery: fetchBaseQuery({
<<<<<<< HEAD:web/ASTU-web-group-2/banking-dashboard/lib/service/CardService.tsx
   baseUrl: "https://astu-bank-dashboard.onrender.com",
=======
    baseUrl: "https://astu-bank-dashboard.onrender.com",
>>>>>>> df7412d5a7cf02a4ba73c728ff86bcb98555c1ab:web/ASTU-web-group-2/banking-dashboard/lib/service/CardService.ts
  }),
  endpoints: (builder) => ({
    getAllCardInfo: builder.query({
      query: (data: { token: string; size: number }) => ({
        url: `/cards?page=0&size=${data.size}`,
        method: "GET",
        headers: {
          Authorization: `Bearer ${data.token}`,
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
} = CreditCardInfoApi;
