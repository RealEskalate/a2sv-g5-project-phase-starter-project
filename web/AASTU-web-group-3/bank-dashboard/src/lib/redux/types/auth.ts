export interface RegisterRequest {
    name: string;
    email: string;
    dateOfBirth: string;
    permanentAddress: string;
    postalCode: string;
    presentAddress: string;
    city: string;
    country: string;
    profilePicture: string;
    username: string;
    password: string;
    preference: {
      currency: string;
      sentOrReceiveDigitalCurrency: boolean;
      receiveMerchantOrder: boolean;
      accountRecommendations: boolean;
      timeZone: string;
      twoFactorAuthentication: boolean;
    };
  };

  export interface RegisterResponse{
        success: boolean,
        message: string,
        data: {
          access_token: string,
          refresh_token: string,
          data: {
            id: string,
            name: string,
            email: string,
            dateOfBirth:string,
            permanentAddress: string,
            postalCode: string,
            username: string,
            presentAddress: string,
            city: string,
            country: string,
            profilePicture: string,
            accountBalance: number,
            role: string,
            preference: {
              currency: string,
              sentOrReceiveDigitalCurrency: boolean,
              receiveMerchantOrder: boolean,
              accountRecommendations: boolean,
              timeZone: string,
              twoFactorAuthentication: boolean
            }
          }
        }
      }