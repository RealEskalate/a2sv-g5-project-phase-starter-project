// setup credential provider for next-auth

import CredentialsProvider from "next-auth/providers/credentials";
import { JWT } from "next-auth/jwt";

export const options = {
  secret: process.env.NEXTAUTH_SECRET,
  providers: [
    CredentialsProvider({
      name: "Credentials",
      credentials: {
        userName: {},
        password: {},
      },
      async authorize(credentials, req) {
        const userName = credentials?.userName;
        const password = credentials?.password;
        try {
          const fetchResponse = await fetch(
            "https://bank-dashboard-aait-latest-sy48.onrender.com/auth/login",
            {
              method: "POST",
              headers: {
                "Content-Type": "application/json",
              },
              body: JSON.stringify({ userName, password }),
            }
          );
          const responseData = await fetchResponse.json();
          console.log("data returned from the server on fetching", responseData);
          if (responseData.success) {
            return responseData;
          } 
        } catch (err) {
          console.log("error", err);
        }
      },
    }),
  ],

  callbacks: {
    async session({ session, token }: { session: any; token: any }) {
      if (token) {
        session.user.access_token = token.access_token;

        session.user.refresh_token = token.refresh_token;
      }
      return session;
    },
    async jwt({ token, user }: { token: JWT; user: any }) {
      if (user) {
        token.access_token = user.access_token;
        token.refresh_token = user.refresh_token;
      }
      return token;
    },
  },
};
