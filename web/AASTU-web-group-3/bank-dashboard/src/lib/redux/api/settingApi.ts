import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { getSession } from "next-auth/react";
import { settingPutUserResponse, settingPutUserRequest, Preference } from "../types/setting";

export const settingApi = createApi({
  reducerPath: "settingApi",
  baseQuery: fetchBaseQuery({
    baseUrl: "https://bank-dashboard-1tst.onrender.com",
    prepareHeaders: async (headers) => {
      const session = await getSession();
      const token = session?.accessToken;

      if (token) {
        console.log(token)
        headers.set("Authorization", `Bearer ${token}`);
      }
      return headers;
    },
  }),
  endpoints: (builder) => ({
    putSetting: builder.mutation<settingPutUserResponse, settingPutUserRequest>({
      query: (formData) => ({
        url: `/user/update`,
        method: 'PUT',
        body: formData, 
        headers: {
            'Content-Type': 'application/json',
          },
      }),
    }),
    putPreference: builder.mutation<settingPutUserResponse, Preference>({
      query: (preferenceData) => ({
        url: `user/update-preference`,
        method: 'PUT',
        body: preferenceData,
        headers: {
          'Content-Type': 'application/json',
        },
      }),
    }),
  }),
});

export const {
  usePutSettingMutation,
  usePutPreferenceMutation
} = settingApi;
