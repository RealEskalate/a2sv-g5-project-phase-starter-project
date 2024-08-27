import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';


export const cardApi = createApi({
  reducerPath: 'cardApi',
  baseQuery: fetchBaseQuery({
    baseUrl: "https://bank-dashboard-aait-latest-sy48.onrender.com",
    prepareHeaders: (headers) => {
      const token = localStorage.getItem('token');
      if (token) {
        headers.set('Authorization', `Bearer ${token}`);
      }
      return headers;
    },
  }),

  tagTypes: [],
  endpoints: (builder) => ({

    getCards: builder.query({
      query: ({page, size}) => ({
        url: `/cards?page=${page}&size=${size}`,
        method: 'GET',
      }),
    }),

    postCards: builder.mutation({
      query: ({balance, cardHolder, expiryDate, passcode, cardType}) => ({
        url: '/cards',
        method: 'POST',
        body: { cardHolder, passcode, cardType, expiryDate, balance },
      }),
    }),

    getCard: builder.query({
      query: ({id}) => ({
        url: `/cards/${id}`,
        method: 'GET',
      }),
    }),

    deleteCard: builder.mutation({
      query: ({id}) => ({
        url: `/cards/${id}`,
        method: 'DELETE',
      }),
    }),

  }),
});

export const {
  useGetCardsQuery,
  usePostCardsMutation,
  useGetCardQuery,
  useDeleteCardMutation,
} = cardApi;



