import { createApi } from "@reduxjs/toolkit/query/react";
import { baseQuery } from "./baseQuery";
import { UpdatedUser, UserResponseType } from "@/types/user.types";

export const profileAPI = createApi({
    reducerPath: "profileAPI",
    baseQuery: baseQuery(),
    endpoints: (builder) => ({
        getProfile: builder.query<UserResponseType, void>({
            query: () => `user/current`,
        }),
        
        getPreferences: builder.query<void, void>({
            query: () => `/user/update-preference`,
        }),

        updateUser: builder.mutation<void, UpdatedUser>({
            query: (body) => ({
                url: `/user/update`,
                method: "PUT",
                body,
            }),
        })
    }),
})

export const { useGetProfileQuery, useUpdateUserMutation } = profileAPI
