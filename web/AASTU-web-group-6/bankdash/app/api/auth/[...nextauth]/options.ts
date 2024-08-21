import type { NextAuthOptions } from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";
import AuthService from "@/app/Services/api/authService";
import { Userx } from "./route";
import LoginValue from "@/types/LoginValue";

export const options: NextAuthOptions = {
  session: {
    strategy: "jwt",
  },
  jwt: {
    secret: process.env.JWT_SECRET,
    maxAge: 24 * 60 * 60,
  },
  providers: [
    CredentialsProvider({
      type: "credentials",
      credentials: {},
      async authorize(credentials:LoginValue | null): Promise<User | null> {
        // Cast credentials to the expected LoginValue type

        if (!credentials) {
          return null;
        }

        const response = await AuthService.login(credentials);
        if (response.success) {
          const data: any = response.data;

          // Define the user object with the required fields
          const userData: User = {
            refreshToken: data.refresh_token,
            accessToken: data.access_token,
            id: ""
          };

          return userData;
        } else {
          return null;
        }
      },
    }),
  ],
  callbacks: {
    async jwt({ token, user }) {
      // If a user was returned by the `authorize` function, merge tokens
      if (user) {
        token.accessToken = user.accessToken;
        token.refreshToken = user.refreshToken;
      }
      return token;
    },
    async session({ session, token }) {
      // Add tokens to session object
      session.accessToken = token.accessToken;
      session.refreshToken = token.refreshToken;
      return session;
    },
  },
  pages: {
    error: "/error",
  },
};
