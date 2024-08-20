import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { session } from "@/session";

export const activeLoanApi = createApi({
  reducerPath: "activeLoanApi",
  baseQuery: fetchBaseQuery({
    baseUrl: "https://bank-dashboard-6acc.onrender.com",
    prepareHeaders: (headers) => {
      headers.set("Authorization", `Bearer ${session}`);
      return headers;
    },
  }),
  endpoints: (builder) => ({
    getAllActiveLoans: builder.query<void, string>({
      query: () => `/active-loans/all`,
    }),
    getActiveLoanById: builder.query<void, string>({
      query: (id) => `/active-loans/${id}`,
    }),

    getMyActiveLoans: builder.query<void, string>({
      query: () => `/active-loans/my-loans`,
    }),

    getDetailActiveLoans: builder.query<void, string>({
      query: () => `/active-loans/detail-data`,
    }),

    postActiveLoans: builder.mutation<
      void,
      {
        loanAmount: number;
        duration: number;
        interestRate: number;
        type: string;
      }
    >({
      query: ({ loanAmount, duration, interestRate, type }) => ({
        url: "/active-loans",
        method: "POST",
        body: { loanAmount, duration, interestRate, type },
        headers: {
          "Content-Type": "application/json",
        },
      }),
    }),

    rejectActiveLoans: builder.mutation<void, {}>({
      query: (id) => ({
        url: `/active-loans/${id}/reject`,
        method: "POST",
      }),
    }),

    acceptActiveLoans: builder.mutation<void, {}>({
      query: (id) => ({
        url: `/active-loans/${id}/accept`,
        method: "POST",
      }),
    }),
  }),
});

export const {
  useGetAllActiveLoansQuery,
  useGetActiveLoanByIdQuery,
  useGetMyActiveLoansQuery,
  useGetDetailActiveLoansQuery,
  usePostActiveLoansMutation,
  useRejectActiveLoansMutation,
  useAcceptActiveLoansMutation,
} = activeLoanApi;
