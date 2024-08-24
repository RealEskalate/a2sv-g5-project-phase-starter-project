
export interface UpdatedUser  {
  name: string;
  email: string;
  dateOfBirth: string;
  permanentAddress: string;
  postalCode: string;
  username: string;
  presentAddress: string;
  city: string;
  country: string;
<<<<<<< HEAD
  profilePicture?: string;
=======
  profilePicture?: string | null;
>>>>>>> 8c3532ff24279cab0d57070c8126ce87ede8b7a4
};

export interface UserPreferenceType {
  currency: string;
  sentOrReceiveDigitalCurrency: boolean;
  receiveMerchantOrder: boolean;
  accountRecommendations: boolean;
  timeZone: string;
  twoFactorAuthentication: boolean;
}

export interface UserDataType {
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
  preference: UserPreferenceType;
}

export interface UserResponseType {
  success: boolean;
  message: string;
  data: UserDataType;
}

export interface InvestmentResponseType {
  success: boolean;
  message: string;
  data: {
    totalInvestment: number;
    rateOfReturn: number;
    yearlyTotalInvestment: {
      time: string;
      value: number;
    }[];
    monthlyRevenue: {
      time: string;
      value: number;
    }[];
  };
}
