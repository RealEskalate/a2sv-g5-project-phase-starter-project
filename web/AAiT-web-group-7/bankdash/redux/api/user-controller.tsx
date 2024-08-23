import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

export const userApi = createApi({
  reducerPath: 'userApi',
  baseQuery: fetchBaseQuery({
    baseUrl: "https://bank-dashboard-latest.onrender.com",
    prepareHeaders: (headers) => {
      const token = localStorage.getItem('token');
      if (token) {
        headers.set('Authorization', `Bearer ${token}`);
      }
      return headers;
    },
  }),

  tagTypes: ['users'],
  endpoints: (builder) => ({

    updateUser: builder.mutation({
      query: ({ name, email, dateOfBirth, username, permanentAddress, postalCode, presentAddress, city, country, profilePicture }) => ({
        url: '/user/update',
        method: 'PUT',
        body: { name, email, dateOfBirth, username, permanentAddress, postalCode, presentAddress, city, country, profilePicture },
      }),
      invalidatesTags: ['users'],
    }),

    userUpdatePreference: builder.mutation({
      query: ({ currency, sentOrReceiveDigitalCurrency, receiveMerchantOrder, accountRecommendations, timeZone, twoFactorAuthentication }) => ({
        url: '/user/update-preference',
        method: 'PUT',
        body: { currency, sentOrReceiveDigitalCurrency, receiveMerchantOrder, accountRecommendations, timeZone, twoFactorAuthentication },
      }),

      invalidatesTags: ['users'],
    }),

    getUser: builder.query({
      query: ({username}) => ({
        url: `/user/${username}`,
        method: 'GET',
      }),
      providesTags: ['users'],
    }),

    userRandomInvestmentData: builder.query({
      query: ({years, months}) => ({
        url: `/user/random-investment-data?years=${years}&months=${months}`,
        method: 'GET',
      }),
      providesTags: ['users'],
    }),

    currentUser: builder.query({
      query: () => ({
        url: '/user/current',
        method: 'GET',
      }),
      providesTags: ['users'],
    }),

  }),
});

export const {
  useUpdateUserMutation,
  useUserUpdatePreferenceMutation,
  useGetUserQuery,
  useUserRandomInvestmentDataQuery,
  useCurrentUserQuery,
} = userApi;



