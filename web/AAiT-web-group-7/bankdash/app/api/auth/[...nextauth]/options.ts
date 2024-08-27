// setup credential provider for next-auth

import CredentialsProvider from "next-auth/providers/credentials";
import { JWT } from "next-auth/jwt";

export const options = {
  secret: process.env.NEXTAUTH_SECRET,
  providers: [
    CredentialsProvider({
      name: "Credentials",
      credentials: {
        username: {},
        password: {},
      },
      async authorize(credentials, req) {
        const userName = credentials?.username;
        const password = credentials?.password;
        console.log(userName, password,"miew");
        try {
          const response = await fetch(
            "https://bank-dashboard-aait-latest-sy48.onrender.com/auth/login",
            {
              method: "POST",
              headers: {
                "Content-Type": "application/json",
              },
              body: JSON.stringify({ userName, password }),
            }
          );
          console.log("first",response)
          const data = await response.json();
          console.log("data returned from the server on fetching", data);
          if (data.success) {
            return data;
          } else {
            throw new Error("Invalid email or password");
          }
        } catch (error) {
          console.log("error", error);
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
