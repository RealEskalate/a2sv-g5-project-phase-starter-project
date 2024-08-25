import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'
import {Preference, User} from '../../../types/User'
import { UserResponse } from '../types/userResponse'
import { UserSignup } from '@/types/UserSignUp'
import { RootState } from '../store'
import { useSession } from 'next-auth/react'

export const bankApi = createApi({
  reducerPath: 'api',
  baseQuery: fetchBaseQuery({ 
    baseUrl: 'https://bank-dashboard-aait-team-2.onrender.com',
  }),
  
  endpoints: builder => ({
    getUserByUsername: builder.query<{data: User}, string>({
      query: (username) => `/user/${username}`
    }),
    getCurrentUser: builder.query<{data: User}, string>({
      query: (token) => ({
          url: `/user/current`,
          method: 'Get',
          headers: {
            contentType: 'application/json',
            Authorization: `Bearer ${token}`
          }
        })
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

    updateUserProfile: builder.mutation<User, {userUpdate: UserUpdate, token: string} >({
        query: ({userUpdate, token}) => ({
          url: '/user/update',
          method: 'PUT',
          headers: {
            contentType: 'application/json',
            Authorization: `Bearer ${token}`
          },
          body: userUpdate,
        }),
      }),
    updateUserPreference: builder.mutation<User, {userUpdate: Preference, token: string} >({
      query: ({userUpdate, token}) => ({
        url: '/user/update-preference',
        method: 'PUT',
        headers: {
          contentType: 'application/json',
          Authorization: `Bearer ${token}`
        },
        body: userUpdate,
      }),
    }),
    changePassword:builder.mutation<User, {credentials: PasswordChange, token: string} >({
      query: ({credentials, token}) => ({
        url: '/auth/change_password',
        method: 'Post',
        headers: {
          contentType: 'application/json',
          Authorization: `Bearer ${token}`
        },
        body: credentials,
      }),
    }),
     
  })
})

export const {
  useGetUserByUsernameQuery,
  useGetCurrentUserQuery,
  useSignupMutation,
  useSigninMutation,
  useUpdateUserProfileMutation,
  useUpdateUserPreferenceMutation,
  useChangePasswordMutation

} = bankApi