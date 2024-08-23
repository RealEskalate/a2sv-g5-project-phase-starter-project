// setup credential provider for next-auth

import CredentialsProvider from "next-auth/providers/credentials";
import { JWT } from "next-auth/jwt";

export const options = {
  secret: process.env.NEXTAUTH_SECRET,
  // url: process.env.NEXTAUTH_URL,

  //   session: { strategy: "jwt" },
  // pages: {
  //   signIn: "/signup",
  // },
  providers: [
    CredentialsProvider({
      name: "Credentials",
      credentials: {
        username: {},
        password: {},
      },
      async authorize(credentials, req) {
        const username = credentials?.username;
        const password = credentials?.password;

        return null;
      },
    }),
  ],

  callbacks: {
    async session({ session, token }: { session: any; token: any }) {
      if (token) {
        session.user.accessToken = token.accessToken;
        session.user.role = token.role;
        session.user.id = token.id;
        session.user.refreshToken = token.refreshToken;
      }
      return session;
    },
    async jwt({ token, user }: { token: JWT; user: any }) {
      if (user) {
        token.id = user.id;
        token.role = user.role;
        token.accessToken = user.accessToken;
        token.refreshToken = user.refreshToken;
      }
      return token;
    },
  },
};
