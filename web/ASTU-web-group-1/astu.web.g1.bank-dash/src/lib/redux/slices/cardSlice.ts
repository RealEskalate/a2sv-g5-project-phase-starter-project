import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import { getSession } from 'next-auth/react';

async function get_session() {
  const session = await getSession();
  return session;
}

export const cardApi = createApi({
  reducerPath: 'cardApi',
  baseQuery: fetchBaseQuery({
    baseUrl: 'https://bank-dashboard-6acc.onrender.com',
    prepareHeaders: async (headers, { getState }) => {
      const session = await get_session();
      if (session && session.accessToken) {
        headers.set('Authorization', `Bearer ${session.accessToken}`);
      }
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
