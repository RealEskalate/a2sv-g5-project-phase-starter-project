import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'
import {User} from '../../../types/User'
import { UserResponse } from '../types/userResponse'
import { UserSignup } from '@/types/UserSignUp'

export const bankApi = createApi({
  reducerPath: 'api',
  baseQuery: fetchBaseQuery({ baseUrl: 'https://bank-dashboard-latest.onrender.com/' }),
  endpoints: builder => ({
    getUserByUsername: builder.query<{data: User}, string>({
      query: (username) => `/user/${username}`
    }),
    signup: builder.mutation<UserResponse, UserSignup>({
      query: (user) => ({
        url: '/auth/register',
        method: 'Post',
        body: user,
      }),
    }),

    signin: builder.mutation<LoginResponse, SigninCredential>({
      query: (credential) => ({
        url: '/auth/login',
        method: 'Post',
        body: credential,
      }),
    }),

    updateUserProfile: builder.mutation<UserUpdate, User>({
        query: (userUpdate) => ({
          url: '/user/update',
          method: 'PUT',
          body: userUpdate,
        }),
      }),
     
  })
})

export const {
  useGetUserByUsernameQuery,
  useSignupMutation,
  useSigninMutation,
  useUpdateUserProfileMutation
} = bankApi