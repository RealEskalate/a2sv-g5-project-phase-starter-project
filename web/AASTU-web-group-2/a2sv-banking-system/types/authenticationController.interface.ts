// authenticationController.interface.ts

export interface Preference {
    currency: string;
    sentOrReceiveDigitalCurrency: boolean;
    receiveMerchantOrder: boolean;
    accountRecommendations: boolean;
    timeZone: string;
    twoFactorAuthentication: boolean;
  }
  
  export interface RegisterRequest {
    name: string;
    email: string;
    dateOfBirth: string;
    permanentAddress: string;
    postalCode: string;
    username: string;
    password: string;
    presentAddress: string;
    city: string;
    country: string;
    profilePicture: string;
    preference: Preference;
  }
  
  export interface RegisterResponseData {
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
  
  export interface RegisterResponse {
    success: boolean;
    message: string;
    data: {
      access_token: string;
      refresh_token: string;
      data: RegisterResponseData;
    };
  }
  
  export interface RefreshTokenResponse {
    success: boolean;
    message: string;
    data: string; // New access token
  }
  
  export interface LoginRequest {
    userName: string;
    password: string;
  }
  
  export interface LoginResponse {
    success: boolean;
    message: string;
    data: {
      access_token: string;
      refresh_token: string;
      data: any; // Assuming `data` contains user information or other details
    };
  }
  export interface ChangePasswordRequest {
    password: string;
    newPassword: string;
  }
  
  export interface ChangePasswordResponse {
    success: boolean;
    message: string;
    data: {}; // Empty object if no additional data is returned
  }  