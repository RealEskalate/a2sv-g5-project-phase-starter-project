import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { formType } from "@/types/formType";


export const authApi = createApi({
    reducerPath: "auth",
    baseQuery: fetchBaseQuery({ baseUrl: "https://bank-aait-web-group-1.onrender.com/"}),
    endpoints: (builder) => ({
        signUp: builder.mutation({
            query: (data:formType) => (console.log("data:", data),{
                url: `/auth/register`,
                method: "POST",
                headers: {"Content-Type": "application/json"},
                body: data
            })
        }),
        signIn: builder.mutation({
            query: (data) => ({
                url: `/auth/login`,
                method: "POST",
                headers: {"Content-Type": "application/json"},
                body: data
            })
        })
    })
 })




export const { useSignUpMutation, useSignInMutation } = authApi