import { User } from '@/types/User';
import NextAuth, { DefaultSession } from "next-auth";

interface AuthUser {
    id: string;
    name: string;
    email: string;
    dateOfBirth: string; 
    permanentAddress: string;
    postalCode: string;
    username: string;
    presentAddress: string;
    city: string;
    country: string;
    profilePicture: string;
    accountBalance: number;
    role: string;
    preference: Preference;
}
declare module "next-auth" {
  interface Session {
    user?: AuthUser;
    access_token?: string;
    refresh_token?: string;
  }

  interface User {
    access_token?: string;
    refresh_token?: string;
    id?: string;
    data?: User;

  }
  
}