import NextAuth, { DefaultSession } from "next-auth";

declare module "next-auth" {
  interface Session {
    user?: {
      access_token?: string;
      refresh_token?: string;
    } & DefaultSession["user"];
  }

  interface JWT {
    access_token: string;
    refresh_token?: string;
  }

  interface User {
    access_token?: string;
    refresh_token?: string;
  }
}
