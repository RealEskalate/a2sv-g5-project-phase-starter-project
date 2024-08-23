import NextAuth, { DefaultSession } from "next-auth";

declare module "next-auth" {
  interface Session {
    user?: {
      accessToken?: string;
      access_token?: string;
      refreshToken?: string;
      refresh_token?: string;
      id_token?: string;
    } & DefaultSession["user"];
  }

  interface JWT {
    accessToken: string;
    access_token: string;
    refreshToken?: string;
    refresh_token?: string;
    id_token?: string;
  }

  interface User {
    accessToken?: string;
    access_token?: string;
    refreshToken?: string;
    refresh_token?: string;
    id_token?: string;
  }
}
