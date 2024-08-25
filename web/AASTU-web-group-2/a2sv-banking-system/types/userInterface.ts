export default interface User {
  user_info: UserInfo;
  preference: Preference;
}

export type UserInfo = {
  id?: string;
  name: string;
  email: string;
  dateOfBirth: string | null;
  permanentAddress: string;
  postalCode: string;
  username: string;
  password?: string;
  presentAddress: string;
  city: string;
  country: string;
  profilePicture: string;
  role?: string;
  accountBalance?: number;
};
export type Preference = {
  currency: string;
  sentOrReceiveDigitalCurrency: boolean;
  receiveMerchantOrder: boolean;
  accountRecommendations: boolean;
  timeZone: string;
  twoFactorAuthentication: boolean;
};
