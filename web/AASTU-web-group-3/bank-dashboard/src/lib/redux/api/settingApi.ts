import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { getSession } from "next-auth/react";
import { settingPutUserResponse, settingPutUserRequest, Preference,RandomInvestmentData } from "../types/setting";
import { url } from "inspector";
import { METHODS } from "http";

export const settingApi = createApi({
  reducerPath: "settingApi",
  baseQuery: fetchBaseQuery({
    baseUrl: process.env.NEXT_PUBLIC_BASE_URL,
    prepareHeaders: async (headers) => {
      const session = await getSession();
      const token = session?.accessToken;

      if (token) {
        console.log(token);
        headers.set("Authorization", `Bearer ${token}`);
      }
      return headers;
    },
  }),
  endpoints: (builder) => ({
    putSetting: builder.mutation<settingPutUserResponse, settingPutUserRequest>(
      {
        query: (formData) => ({
          url: `/user/update`,
          method: "PUT",
          body: formData,
          headers: {
            "Content-Type": "application/json",
          },
        }),
      }
    ),
    putPreference: builder.mutation<settingPutUserResponse, Preference>({
      query: (preferenceData) => ({
        url: `user/update-preference`,
        method: "PUT",
        body: preferenceData,
        headers: {
          "Content-Type": "application/json",
        },
      }),
    }),

    getCurrentUser: builder.query<settingPutUserResponse, void>({
      query: () => ({
        url: `/user/current`,
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      }),
  }),

  getRandomInvestmentData: builder.query<
      RandomInvestmentData,
      { years: number; months: number }
    >({
      query: ({ years, months }) => `/user/random-investment-data?years=${years}&months=${months}`,
    }),
}),

});

export const {
  usePutSettingMutation,
  usePutPreferenceMutation,
  useGetCurrentUserQuery,
  useGetRandomInvestmentDataQuery
} = settingApi;
