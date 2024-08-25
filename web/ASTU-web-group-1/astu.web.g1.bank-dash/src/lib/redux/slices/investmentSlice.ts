import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { baseQuery } from "../api/baseQuery";
import { InvestmentType } from "@/types/investmentType";

export const investmentApi = createApi({
  reducerPath: "investmentApi",
  baseQuery: baseQuery(),
  endpoints: (builder) => ({

    getInvestmentItems: builder.query<InvestmentType, { years: number; months: number }>({
      query: ({ years, months }) => `/user/random-investment-data?years=${years}&months=${months}`,
    }),
  }),
});

export const {
  useGetInvestmentItemsQuery,
} = investmentApi;
