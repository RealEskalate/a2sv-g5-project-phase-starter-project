import NextAuth, { NextAuthOptions } from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";

import { JWT } from "next-auth/jwt";
import { signInWithCredentials } from "@/lib/auth";

export const authOptions: NextAuthOptions = {
  providers: [
    CredentialsProvider({
      name: "credentials",
      credentials: {
        username: { label: "Username", type: "text" },
        password: { label: "Password", type: "password" },
      },
      async authorize(credentials) {
        try {
          if (!credentials) throw new Error("No credentials provided");

          const response = await signInWithCredentials(credentials);
          console.log("response", response);
          if (response.success) {
            return {
              id: response.data.id,
              name: response.data.name,
              email: response.data.email,
              accessToken: response.data.access_token,
              refreshToken: response.data.refresh_token,
            };
          }else{
            throw new Error(response.message || "Failed to sign in");
          }
          return null;
        } catch (error) {
          console.error("Authorize error:", error);
          return null;
        }
      },
    }),
  ],
  callbacks: {
    async jwt({ token, user }) {
      if (user) {
        token.accessToken = user.accessToken;
        token.refreshToken = user.refreshToken;
      }
      return token;
    },
    async session({ session, token }) {
        if (token) {
          session.user.accessToken = token.accessToken;
          session.user.refreshToken = token.refreshToken;
        }
      return session;
    },
  },
  pages: {
    signIn: "/auth/sign-in",
  },
  secret: process.env.NEXTAUTH_SECRET,
};


