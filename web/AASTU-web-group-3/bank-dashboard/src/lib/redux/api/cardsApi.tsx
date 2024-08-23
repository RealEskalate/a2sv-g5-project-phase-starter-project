import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { getSession } from "next-auth/react";
import {
  PaginatedCardsResponse,
  CreateCardRequest,
  CreateCardResponse,
  CardDetail,
  Card,
} from "../types/cards";

export const cardsApi = createApi({
  reducerPath: "cardsApi",
  baseQuery: fetchBaseQuery({
    baseUrl: process.env.NEXT_PUBLIC_BASE_URL,
    prepareHeaders: async (headers) => {
      const session = await getSession();
      const token = session?.accessToken;

      if (token) {
        headers.set("Authorization", `Bearer ${token}`);
      }
      return headers;
    },
  }),
  endpoints: (builder) => ({
    getCards: builder.query<
      PaginatedCardsResponse,
      { page: number; size: number }
    >({
      query: ({ page, size }) => `/cards?page=${page}&size=${size}`,
    }),

    getCardById: builder.query<CardDetail, string>({
      query: (id) => `/cards/${id}`,
    }),

    createCard: builder.mutation<CreateCardResponse, CreateCardRequest>({
      query: (card) => ({
        url: "/cards",
        method: "POST",
        body: card,
        headers: {
          "Content-Type": "application/json",
        },
      }),
    }),

    deleteCard: builder.mutation<void, string>({
      query: (id) => ({
        url: `/cards/${id}`,
        method: "DELETE",
      }),
    }),
  }),
});

export const {
  useGetCardsQuery,
  useGetCardByIdQuery,
  useCreateCardMutation,
  useDeleteCardMutation,
} = cardsApi;
