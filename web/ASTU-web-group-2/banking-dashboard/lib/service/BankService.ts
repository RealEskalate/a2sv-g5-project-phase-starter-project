import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import process from "process";

export const bankApi = createApi({
  reducerPath: "bankDash",
  baseQuery: fetchBaseQuery({
    baseUrl: "https://astu-bank-dashboard.onrender.com",
<<<<<<< HEAD:web/ASTU-web-group-2/banking-dashboard/lib/service/BankService.tsx
    
=======
>>>>>>> df7412d5a7cf02a4ba73c728ff86bcb98555c1ab:web/ASTU-web-group-2/banking-dashboard/lib/service/BankService.ts
  }),
  endpoints: (builder) => ({
    getBankService: builder.query({
      query: (data: { accessToken: string; size: number; page: number }) => ({
        url: `/bank-services?size=${data.size}&page=${data.page}`,
        method: "GET",
        headers: {
          Authorization: `Bearer ${data.accessToken}`,
        },
      }),
    }),

    // Add your requests here the same way 
    postBankService: builder.mutation({
      query: (data: {
        accessToken: string;
        name: string;
        details: string;
        numberOfUsers: number;
        status: string;
        type: string;
        icon: string;
      }) => ({
        url: `/bank-services`,
        method: "GET",
        headers: {
          Authorization: `Bearer ${data.accessToken}`,
        },
        body: {
          name: data.name,
          details: data.details,
          numberOfUsers: data.numberOfUsers,
          status: data.status,
          type: data.type,
          icon: data.icon,
        },
      }),
    }),
  }),
});

export const { useGetBankServiceQuery, usePostBankServiceMutation } = bankApi;
