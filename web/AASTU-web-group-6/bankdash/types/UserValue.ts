interface UserPreference {
    currency: string;
    sentOrReceiveDigitalCurrency: boolean;
    receiveMerchantOrder: boolean;
    accountRecommendations: boolean;
    timeZone: string;
    twoFactorAuthentication: boolean;
  }
  
  interface UserValue {
    name: string;
    email: string;
    dateOfBirth: string;
    permanentAddress: string;
    postalCode: string;
    username: string;
    password: string;
    confirmPassword: string;
    presentAddress: string;
    city: string;
    country: string;
    profilePicture: string;
    preference: UserPreference;
  }
  
  export default UserValue;
  