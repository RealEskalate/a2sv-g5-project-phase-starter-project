import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

type UpdatedUser = {
  name: string;
  email: string;
  dateOfBirth: string;
  permanentAddress: string;
  postalCode: string;
  username: string;
  presentAddress: string;
  city: string;
  country: string;
  profilePicture?: string | null;
};
type UpdatedPreference = {
  currency: string;
  sentOrReceiveDigitalCurrency: boolean;
  receiveMerchantOrder: boolean;
  accountRecommendations: boolean;
  timeZone: string;
  twoFactorAuthentication: boolean;
};

export const userApi = createApi({
  reducerPath: "userDashboard",
  baseQuery: fetchBaseQuery({
    baseUrl: "https://astu-bank-dashboard.onrender.com",
  }),
  endpoints: (builder) => ({
    getCurrentUser: builder.query({
      query: (accessToken: string) => ({
        url: "/user/current",
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      }),
    }),
    updateUser: builder.mutation({
      query: ({
        accessToken,
        updatedUser,
      }: {
        accessToken: string;
        updatedUser: UpdatedUser;
      }) => ({
        url: "/user/update",
        method: "PUT",
        body: updatedUser,
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      }),
    }),
    updatePreference: builder.mutation({
      query: ({
        accessToken,
        updatedPreference,
      }: {
        accessToken: string;
        updatedPreference: UpdatedPreference;
      }) => ({
        url: "/user/update-preference",
        method: "PUT",
        body: updatedPreference,
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      }),
    }),
  }),
});

export const {
  useGetCurrentUserQuery,
  useUpdateUserMutation,
  useUpdatePreferenceMutation,
} = userApi;
