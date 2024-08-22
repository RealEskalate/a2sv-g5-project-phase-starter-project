import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

export const endpointsApi = createApi({
  reducerPath: 'endpointsApi',
  baseQuery: fetchBaseQuery({
    baseUrl: "http://localhost:5000",
  }),
  endpoints: (builder) => ({

  }),
})
