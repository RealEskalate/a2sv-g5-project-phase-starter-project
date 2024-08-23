import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

export const bankserviceApi = createApi({
  reducerPath: 'bankserviceApi',
  baseQuery: fetchBaseQuery({
    baseUrl: "",
    // the header depends on where we put our access token. either cookies or local storage
    prepareHeaders: (headers) => {
      const token = localStorage.getItem('token');
      if (token) {
        headers.set('Authorization', `Bearer ${token}`);
      }
      return headers;
    },
  }),

  tagTypes: [],
  endpoints: (builder) => ({}),
});

export const {} = bankserviceApi;



