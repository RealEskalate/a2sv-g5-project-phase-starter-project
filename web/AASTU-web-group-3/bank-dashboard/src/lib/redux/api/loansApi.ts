import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import { RootState } from '../store';
import { getSession } from 'next-auth/react';
import { Loan } from '../types/loans';

interface LoanRequest {
  loanAmount: number;
  duration: number;
  interestRate: number;
  type: string;
}


export interface LoanResponse {
  success: boolean;
  message: string;
  data: {
    serialNumber: string;
    loanAmount: number;
    amountLeftToRepay: number;
    duration: number;
    interestRate: number;
    installment: number;
    type: string;
    activeLoneStatus: string;
    userId: string;
  };
}

interface LoansResponse {
  success: boolean;
  message: string;
  data: LoanResponse["data"][];
}

interface LoanDetailDataResponse {
  success: boolean;
  message: string;
  data: {
    personalLoan: number;
    businessLoan: number;
    corporateLoan: number;
  };
}

export const loansApi = createApi({
  reducerPath: "loansApi",
  baseQuery: fetchBaseQuery({
    baseUrl: process.env.NEXT_PUBLIC_BASE_URL,
    prepareHeaders: async (headers, { getState }) => {
      const session = await getSession();
      const token = session?.accessToken;
      if (token) {
        headers.set("Authorization", `Bearer ${token}`);
      }
      return headers;
    },
  }),
  endpoints: (builder) => ({
    createLoan: builder.mutation<LoanResponse, LoanRequest>({
      query: (loanData) => ({
        url: "/active-loans",
        method: "POST",
        body: loanData,
      }),
    }),
    rejectLoan: builder.mutation<{ success: boolean; message: string }, string>(
      {
        query: (id) => ({
          url: `/active-loans/${id}/reject`,
          method: "POST",
        }),
      }
    ),
    approveLoan: builder.mutation<LoanResponse, string>({
      query: (id) => ({
        url: `/active-loans/${id}/approve`,
        method: "POST",
      }),
    }),
    getLoanById: builder.query<LoanResponse, string>({
      query: (id) => `/active-loans/${id}`,
    }),
    getMyLoans: builder.query<Loan, {page:number, size:number}>({
      query: ({ page, size }) => `/active-loans/my-loans?page=${page}&size=${size}`,
    }),
    getLoanDetailData: builder.query<LoanDetailDataResponse, void>({
      query: () => "/active-loans/detail-data",
    }),
    getAllLoans: builder.query<LoansResponse, void>({
      query: () => "/active-loans/all",
    }),
  }),
});

export const {
  useCreateLoanMutation,
  useRejectLoanMutation,
  useApproveLoanMutation,
  useGetLoanByIdQuery,
  useGetMyLoansQuery,
  useGetLoanDetailDataQuery,
  useGetAllLoansQuery,
} = loansApi;
