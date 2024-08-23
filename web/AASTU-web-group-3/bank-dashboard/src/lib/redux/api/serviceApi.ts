import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { getSession } from "next-auth/react";
import { Service, ServicePostRequest, ServiceResponce } from "../types/service";

export const serviceApi = createApi({
  reducerPath: "serviceApi",
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
    getAllService: builder.query<
      ServiceResponce,
      { size: number; page: number }
    >({
      query: ({ size, page }) => `/bank-services?size=${size}&page=${page}`,
    }),

    getServiceById: builder.query<ServiceResponce, string>({
      query: (id) => `/bank-services/${id}`,
    }),

    createService: builder.mutation<ServiceResponce, ServicePostRequest>({
      query: (Service) => ({
        url: "/bank-services",
        method: "POST",
        body: Service,
        headers: {
          "Content-Type": "application/json",
        },
      }),
    }),

    deleteService: builder.mutation<ServiceResponce, string>({
      query: (id) => ({
        url: `/bank-services/${id}`,
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
        },
      }),
    }),

    putService: builder.mutation<ServiceResponce, Service>({
      query: (service) => ({
        url: `/bank-services/${service.id}`,
        method: "PUT",
        body: {
          name: service.name,
          details: service.details,
          numberOfUsers: service.numberOfUsers,
          status: service.status,
          type: service.type,
          icon: service.icon,
        },
        headers: {
          "Content-Type": "application/json",
        },
      }),
    }),

    getServiceBySearch: builder.query<ServiceResponce, string>({
      query: (query) => `/bank-services/search?query=${query}`,
    }),
  }),
});

export const {
  useGetAllServiceQuery,
  useGetServiceByIdQuery,
  useCreateServiceMutation,
  useDeleteServiceMutation,
  usePutServiceMutation,
  useGetServiceBySearchQuery,
} = serviceApi;
