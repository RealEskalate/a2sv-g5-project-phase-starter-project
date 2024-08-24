import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const authApi = createApi({
  reducerPath: "authApi",
  baseQuery: fetchBaseQuery({
    baseUrl: "https://bank-dashboard-latest.onrender.com",
    prepareHeaders: (headers) => {
      const token = localStorage.getItem("token");
      if (token) {
        headers.set("Authorization", `Bearer ${token}`);
      }
      return headers;
    },
  }),

  endpoints: (builder) => ({
    userRegistration: builder.mutation({
      query: ({
        name,
        email,
        dateOfBirth,
        password,
        username,
        permanentAddress,
        postalCode,
        presentAddress,
        city,
        country,
        profilePicture,
        preference,
      }) => ({
        url: "/auth/register",
        method: "POST",
        body: {
          name,
          email,
          dateOfBirth,
          password,
          username,
          permanentAddress,
          postalCode,
          presentAddress,
          city,
          country,
          profilePicture,
          preference,
        },
      }),
    }),

    userLogin: builder.mutation({
      query: ({ userName, password }) => ({
        url: "/auth/login",
        method: "POST",
        body: { userName, password },
      }),
    }),

    refreshToken: builder.mutation({
      query: () => ({
        url: "/auth/refresh_token",
        method: "POST",
      }),
    }),

    changePassword: builder.mutation({
      query: ({ password, newPassword }) => ({
        url: "/auth/change_password",
        method: "POST",
        body: { password, newPassword },
      }),
    }),
  }),
});

export const {
  useUserRegistrationMutation,
  useUserLoginMutation,
  useRefreshTokenMutation,
  useChangePasswordMutation } = authApi;
  