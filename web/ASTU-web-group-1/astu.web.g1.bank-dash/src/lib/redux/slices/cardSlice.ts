import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import { session } from '@/session';

export const cardApi = createApi({
  reducerPath: 'cardApi',
  baseQuery: fetchBaseQuery({
    baseUrl: 'https://bank-dashboard-6acc.onrender.com',
    prepareHeaders: (headers, { getState }) => {
      // You can add custom headers here
      headers.set('Authorization', `Bearer ${session}`);
      headers.set('Content-Type', 'application/json');
      return headers;
    },
  }),
  endpoints: (builder) => ({
    getAllCards: builder.query<void, void>({
      query: () => '/cards',
    }),
  }),
});

export const { useGetAllCardsQuery } = cardApi;
