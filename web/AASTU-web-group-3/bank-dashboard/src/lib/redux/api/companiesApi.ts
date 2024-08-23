import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { getSession } from "next-auth/react";
import {
  CompanyResponse,
  CompaniesResponse,
  TrendingCompaniesResponse,
  Company,
} from "../types/companies";

export const companiesApi = createApi({
  reducerPath: "companiesApi",
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
    getCompanyById: builder.query<CompanyResponse, string>({
      query: (id) => `/companies/${id}`,
    }),

    updateCompany: builder.mutation<
      CompanyResponse,
      { id: string; company: Omit<Company, "id"> }
    >({
      query: ({ id, company }) => ({
        url: `/companies/${id}`,
        method: "PUT",
        body: company,
        headers: {
          "Content-Type": "application/json",
        },
      }),
    }),

    deleteCompany: builder.mutation<
      { success: boolean; message: string },
      string
    >({
      query: (id) => ({
        url: `/companies/${id}`,
        method: "DELETE",
      }),
    }),

    getAllCompanies: builder.query<
      CompaniesResponse,
      { size: number; page: number }
    >({
      query: ({ size, page }) => `/companies?size=${size}&page=${page}`,
    }),

    createCompany: builder.mutation<CompanyResponse, Omit<Company, "id">>({
      query: (company) => ({
        url: "/companies",
        method: "POST",
        body: company,
        headers: {
          "Content-Type": "application/json",
        },
      }),
    }),

    getTrendingCompanies: builder.query<TrendingCompaniesResponse, void>({
      query: () => `/companies/trending-companies`,
    }),
  }),
});

export const {
  useGetCompanyByIdQuery,
  useUpdateCompanyMutation,
  useDeleteCompanyMutation,
  useGetAllCompaniesQuery,
  useCreateCompanyMutation,
  useGetTrendingCompaniesQuery,
} = companiesApi;
