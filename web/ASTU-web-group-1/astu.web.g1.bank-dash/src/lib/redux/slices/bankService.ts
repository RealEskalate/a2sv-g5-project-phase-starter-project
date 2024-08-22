import { createApi } from '@reduxjs/toolkit/query/react';
import { baseQuery } from '../api/baseQuery';
import { pageSize } from '@/types/page-size.type';
import { BankServiceResponseType, singleBankServiceResponseType } from '@/types/bank-service.types';

export const bankServiceApi = createApi({
  reducerPath: 'bankServiceApi',
  baseQuery: baseQuery(),
  endpoints: (builder) => ({
    getBankService: builder.query<BankServiceResponseType, pageSize>({
      query: (page) => `/bank-services?page=${page.page}&size=${page.size}`,
    }),
    getBankServiceById: builder.query<singleBankServiceResponseType, string>({
      query: (id) => `/bank-services/${id}`,
    }),
  }),
});

export const { useGetBankServiceQuery, useGetBankServiceByIdQuery } = bankServiceApi;
