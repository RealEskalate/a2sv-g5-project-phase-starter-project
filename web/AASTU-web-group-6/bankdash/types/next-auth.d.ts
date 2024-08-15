import NextAuth from "next-auth";

declare module "next-auth" {
  interface User {
    accessToken?: string;
    // other custom fields if needed
  }

  interface Session {
    accessToken?: string;
  }

  interface JWT {
    accessToken?: string;
  }
}
