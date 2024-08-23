import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { baseQuery } from "../api/baseQuery";
import { ActiveLoanResponseType, ActiveLoansDetailResponseType } from "@/types/active-loan.types";

export const activeLoanApi = createApi({
  reducerPath: "activeLoanApi",
  baseQuery: baseQuery(),
  endpoints: (builder) => ({
      getAllActiveLoans: builder.query<ActiveLoanResponseType, void>({
        query: () => `/active-loans/all`,
    }),
    getActiveLoanById: builder.query<void, string>({
      query: (id) => `/active-loans/${id}`,
    }),

    getMyActiveLoans: builder.query<void, void>({
      query: () => `/active-loans/my-loans`,
    }),

    getDetailActiveLoans: builder.query<ActiveLoansDetailResponseType, void>({
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
