import type { NextAuthOptions } from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";
import AuthService from "@/app/Services/api/authService";
import LoginValue from "@/types/LoginValue";

// Define or import your UserValue and Token types
interface UserValue {
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
  secret: process.env.NEXTAUTH_SECRET,
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
      async authorize(credentials): Promise<any> {
        if (!credentials) {
          return null;
        }
        const data =  credentials as LoginValue;
        const response = await AuthService.login(data);
        if (response.success) {
          const data: any = response.data;
          // console.log("Response",data)

          const userData: UserValue = {
            refreshToken: data.refresh_token,
            accessToken: data.access_token,
          };
          // console.log("UserValue AUTH",userData)

          return userData as UserValue;
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
        // Cast user to the UserValue type
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
