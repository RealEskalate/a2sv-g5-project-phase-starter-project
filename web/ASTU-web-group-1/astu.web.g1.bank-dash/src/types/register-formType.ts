export interface FormType {
  name: string;
  email: string;
  dateOfBirth: Date;
  permanentAddress: string;
  postalCode: string;
  username: string;
  password: string;
  presentAddress: string;
  city: string;
  country: string;
  profilePicture: string;
  preference: {
    currency: string;
    sentOrReceiveDigitalCurrency: boolean;
    receiveMerchantOrder: boolean;
    accountRecommendations: boolean;
    timeZone: string;
    twoFactorAuthentication: boolean;
  };
}
