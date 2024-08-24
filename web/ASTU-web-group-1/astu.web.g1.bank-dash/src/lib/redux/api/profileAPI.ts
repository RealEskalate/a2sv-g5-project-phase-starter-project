import { createApi } from "@reduxjs/toolkit/query/react";
import { baseQuery } from "./baseQuery";
import { UpdatedUser, UserPreferenceType, UserResponseType } from "@/types/user.types";

export const profileAPI = createApi({
  reducerPath: "profileAPI",
  baseQuery: baseQuery(),
  endpoints: (builder) => ({
    getProfile: builder.query<UserResponseType, void>({
      query: () => `user/current`,
    }),

    updateUser: builder.mutation<void, UpdatedUser>({
      query: (body) => ({
        url: `/user/update`,
        method: "PUT",
        body,
      }),
    }),

    updatePreference: builder.mutation<void, UserPreferenceType>({
      query: (body) => ({
        url: `/user/update-preference`,
        method: "PUT",
        body,
      }),
    }),
  }),
});

export const { useGetProfileQuery, useUpdateUserMutation, useUpdatePreferenceMutation } = profileAPI;
