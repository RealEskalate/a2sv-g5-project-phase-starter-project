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

 export type BankService = {
  id: string;
  name: string;
  details: string;
  numberOfUsers: number;
  status: string;
  type: string;
  icon: string;
};

 export type BankServiceResponse = {
  success: boolean;
  message: string;
  data: {
    content: BankService[];
  };
};

