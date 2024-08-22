import type { NextAuthOptions } from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";
import AuthService from "@/app/Services/api/authService";

// Define or import your User and Token types
interface User {
  refreshToken: string;
  accessToken: string;
}

interface Token {
  accessToken?: string;
  refreshToken?: string;
}

interface Session {
  accessToken?: string;
  refreshToken?: string;
}

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
      async authorize(credentials) {
        if (!credentials) {
          return null;
        }

        const response = await AuthService.login(credentials);
        if (response.success) {
          const data: any = response.data;
          console.log("Response",data)

          const userData: User = {
            refreshToken: data.refresh_token,
            accessToken: data.access_token,
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
      console.log("USER",user)
      if (user) {
        // Cast user to the User type
        token.accessToken = user.accessToken as string;
        token.refreshToken = user.refreshToken as string;
      }
      return token;
    },
    async session({ session, token }) {
      session.accessToken = token.accessToken;
      session.refreshToken = token.refreshToken;
      return session;
    },
  },
  pages: {
    error: "/error",
  },
};
