import { createApi } from "@reduxjs/toolkit/query/react";
import { baseQuery } from "./baseQuery";
import { ChangePasswordResponse, UpdatedUser, UserPreferenceType, UserResponseType } from "@/types/user.types";

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

    changePassoword: builder.mutation<ChangePasswordResponse, { password: string; newPassword: string }>({
      query: ({ password, newPassword }) => ({
        url: `/auth/change_password`,
        method: "POST",
        body: { password, newPassword },
      }),
    }),
  }),
});

export const { useGetProfileQuery, useUpdateUserMutation, useUpdatePreferenceMutation, useChangePassowordMutation } = profileAPI;
