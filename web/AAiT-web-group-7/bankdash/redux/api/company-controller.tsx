import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

export const companyApi = createApi({
  reducerPath: 'companyApi',
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

  tagTypes: ['companies'],
  endpoints: (builder) => ({

    getCompany: builder.query({
      query: ({id}) => ({
        url: `/companies/${id}`,
        method: 'GET',
      }),
      providesTags: ['companies'],
    }),

    updateCompany: builder.mutation({
      query: ({id, companyName, type, icon}) => ({
        url: `/companies/${id}`,
        method: 'PUT',
        body: { companyName, type, icon },
      }),
      invalidatesTags: ['companies'],
    }),

    deleteCompany: builder.mutation({
      query: ({id}) => ({
        url: `/companies/${id}`,
        method: 'DELETE',
      }),
      invalidatesTags: ['companies'],
    }),

    getCompanies: builder.query({
      query: ({page, size}) => ({
        url: `/companies?page=${page}&size=${size}`,
        method: 'GET',
      }),
      providesTags: ['companies'],
    }),

    postCompany: builder.mutation({
      query: ({companyName, type, icon}) => ({
        url: '/companies',
        method: 'POST',
        body: { companyName, type, icon },
      }),
      invalidatesTags: ['companies'],
    }),

    getTrendingCompanies: builder.query({
      query: () => ({
        url: '/companies/trending-companies',
        method: 'GET',
      }),
      providesTags: ['companies'],
    }),
  }),
});


export const {
  useGetCompanyQuery,
  useUpdateCompanyMutation,
  useDeleteCompanyMutation,
  useGetCompaniesQuery,
  usePostCompanyMutation,
  useGetTrendingCompaniesQuery,
} = companyApi;



