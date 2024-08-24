import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

export const bankserviceApi = createApi({
  reducerPath: 'bankserviceApi',
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

  tagTypes: ['bank-services'],
  endpoints: (builder) => ({
    
    getbankService: builder.query({
      query: ({id}) => ({
        url: `/bank-services/${id}`,
        method: 'GET',
      }),
      providesTags: ['bank-services'],
    }),

    updatebankService: builder.mutation({
      query: ({id, name, type, icon, details, numberOfUsers, status}) => ({
        url: `/bank-services/${id}`,
        method: 'PUT',
        body: { name, type, icon, details, numberOfUsers, status },
      }),
      invalidatesTags: ['bank-services'],
    }),

    deletebankService: builder.mutation({
      query: ({id}) => ({
        url: `/bank-services/${id}`,
        method: 'DELETE',
      }),
      invalidatesTags: ['bank-services'],
    }),

    getbankServices: builder.query({
      query: ({page, size}) => ({
        url: `/bank-services?page=${page}&size=${size}`,
        method: 'GET',
      }),
      providesTags: ['bank-services'],
    }),
    
    postbankServices: builder.mutation({
      query: ({name, type, icon, details, numberOfUsers, status}) => ({
        url: '/bank-services',
        method: 'POST',
        body: { name, type, icon, details, numberOfUsers, status },
      }),
      invalidatesTags: ['bank-services'],
    }),

    searchbankService: builder.query({
      query: ({query}) => ({
        url: `/bank-services/search?query=${query}`,
        method: 'GET',
      }),
      providesTags: ['bank-services'],
    }),
  }),
});

export const {
  useGetbankServiceQuery,
  useUpdatebankServiceMutation,
  useDeletebankServiceMutation,
  useGetbankServicesQuery,
  usePostbankServicesMutation,
  useSearchbankServiceQuery } = bankserviceApi;
