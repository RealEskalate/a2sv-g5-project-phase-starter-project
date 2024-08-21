import NextAuth from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";
import AuthService from "@/app/Services/api/authService";
import { options } from "./options";

export interface User {
  refreshToken: string;
  accessToken: string;
}

const handler = NextAuth(options);

export { handler as GET, handler as POST };
