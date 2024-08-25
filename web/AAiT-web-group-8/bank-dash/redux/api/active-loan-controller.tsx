import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

export const loanApi = createApi({
  reducerPath: 'loanApi',
  baseQuery: fetchBaseQuery({
    baseUrl: "https://bank-dashboard-latest.onrender.com",
    prepareHeaders: (headers) => {
      const token = localStorage.getItem('token');
      if (token) {
        headers.set('Authorization', `Bearer ${token}`);
      }
      return headers;
    },
  }),

  tagTypes: ['activeLoans'],
  endpoints: (builder) => ({

    activeLoans: builder.mutation({
      query: ({loanAmount, duration, interestRate, type}) => ({
        url: '/active-loans',
        method: 'POST',
        body: { loanAmount, duration, interestRate, type }
      }),
      invalidatesTags: ['activeLoans'],
    }),

    rejectActiveLoan: builder.mutation({
      query: ({id}) => ({
        url: `/active-loans/${id}/reject`,
        method: 'POST',
      }),
      invalidatesTags: ['activeLoans'],
    }),

    approveActiveLoan: builder.mutation({
      query: ({id}) => ({
        url: `/active-loans/${id}/approve`,
        method: 'POST',
      }),
      invalidatesTags: ['activeLoans'],
    }),

    singleActiveLoan: builder.query({
      query: ({id}) => ({
        url: `/active-loans/${id}`,
        method: 'GET',
      }),
      providesTags: ['activeLoans'],
    }),

    allmyActiveLoans: builder.query({
      query: ({page, size}) => ({
        url: `/active-loans/my-loans?page=${page}&size=${size}`,
        method: 'GET',
      }),
      providesTags: ['activeLoans'],
    }),

    activeLoansDetails: builder.query({
      query: () => ({
        url: '/active-loans/detail-data',
        method: 'GET',
      }),
      providesTags: ['activeLoans'],
    }),

    allActiveLoans: builder.query({
      query: ({page, size}) => ({
        url: `/active-loans/all?page=${page}&size=${size}`,
        method: 'GET',
      }),
      providesTags: ['activeLoans'],
    }),
  }),
});

export const {
  useActiveLoansMutation,
  useRejectActiveLoanMutation,
  useApproveActiveLoanMutation,
  useSingleActiveLoanQuery,
  useAllmyActiveLoansQuery,
  useActiveLoansDetailsQuery,
  useAllActiveLoansQuery } = loanApi;
  