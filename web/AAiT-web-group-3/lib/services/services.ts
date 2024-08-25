import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
const SERVER_URL = process.env.NEXT_PUBLIC_SERVER_URL;

if (!SERVER_URL) throw Error("SERVER_URL is undefined");

import { BankServiceResponse } from "@/types";

export const serviceAPI = createApi({
  reducerPath: "services",
  baseQuery: fetchBaseQuery({ baseUrl: SERVER_URL }),
  endpoints: (builder) => ({
    getAllServices: builder.query<BankServiceResponse, { accessToken: string, page: number, size: number }>({
      query: ({ accessToken, page, size }) => ({
        url: `/bank-services`,
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${accessToken}`, 
        },
        params: {
          page,
          size,
        },
      }) 
    }),
  }),
});

export const { useGetAllServicesQuery } = serviceAPI;
