
// types.d.ts
import { Session } from "next-auth";
import { JWT } from "next-auth/jwt";
import { IconType } from "react-icons";

export interface Props {
  children: React.ReactNode;
}

export type ElementType = {
  text: string;
  destination: string;
  icon: IconType;
};

export interface navigationValue {
  activePage: string;
  toggle: boolean;
}

export interface EditProfileFormData {
  name: string;
  userName: string;
  email: string;
  dateOfBirth: string;
  presentAddress: string;
  city: string;
  permanentAddress: string;
  postalCode: string;
  country: string;
}

export interface PreferenceFormData {
  currency: string;
  timeZone: string;
  digitalCurrency: boolean;
  merchantOrder: boolean;
  recommendations: boolean;
}

export interface SecurityFormData {
  twoFactorAuthentication: boolean;
  currentPassword: string;
  newPassword: string;
}

declare module "next-auth" {
  interface Session {
    user: {
      name?: string | null;
      email?: string | null;
      image?: string | null;
      role?: string;
      accessToken?: string;
      refreshToken?: string;
    };
  }
}

declare module "next-auth/jwt" {
  interface JWT {
    role?: string;
    accessToken?: string;
    refreshToken?: string;
  }
}

export interface authValue {
  // name: string | null;
  // email: string | null;
  // password: string | null;
  // role: string | null;
  accessToken: string;
  errorMessage: string | null;
  isLoading: boolean;
}