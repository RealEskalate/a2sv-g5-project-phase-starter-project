import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import { getSession } from 'next-auth/react';
import { baseQuery } from '../api/baseQuery';
import { pageSize } from '@/types/page-size.type';

async function get_session() {
  const session = await getSession();
  return session;
}

export const cardApi = createApi({
  reducerPath: 'cardApi',
  baseQuery: baseQuery(),
  endpoints: (builder) => ({
    getAllCards: builder.query<void, pageSize>({
      query: (page) => `/cards?page=${page.page}&size=${page.size}`,
    }),
  }),
});

export const { useGetAllCardsQuery } = cardApi;
