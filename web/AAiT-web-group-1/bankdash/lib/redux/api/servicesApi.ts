import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { getSession } from "next-auth/react";

export const serviceApi = createApi({
    reducerPath: "services",
    baseQuery: fetchBaseQuery({
        baseUrl: "https://bank-aait-web-group-1.onrender.com/",
        prepareHeaders: async (headers) => {
            const session = await getSession();
            if (session && session.user?.accessToken) {
                headers.set('Authorization', `Bearer ${session.user.accessToken}`);
            }
            return headers;
        }
    }),
    endpoints: (builder) => ({
        getServices: builder.query({
            query: ({ page, size }: { page: number; size: number }) => 
                `/bank-services?page=${page}&size=${size}`,
            keepUnusedDataFor: 300,
        }),
    }),
});

export const { useGetServicesQuery } = serviceApi;
