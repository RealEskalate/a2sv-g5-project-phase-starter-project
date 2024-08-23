import { createApi } from "@reduxjs/toolkit/query/react";
import { baseQuery } from "./baseQuery";
import { UserDataType, UserResponseType } from "@/types/user.types";

export const profileAPI = createApi({
    reducerPath: "profileAPI",
    baseQuery: baseQuery(),
    endpoints: (builder) => ({
        getProfile: builder.query<UserResponseType, void>({
            query: () => `user/current`,
        }),

        updateProfile: builder.mutation<void, UserDataType>({
            query: (body) => ({
                url: `user/update`,
                method: "POST",
                body,
            }),
        })
    }),
})

export const { useGetProfileQuery, useUpdateProfileMutation } = profileAPI
