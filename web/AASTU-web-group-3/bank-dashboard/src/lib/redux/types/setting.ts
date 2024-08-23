export interface Preference{
    currency: string,
    sentOrReceiveDigitalCurrency: boolean,
    receiveMerchantOrder: boolean,
    accountRecommendations: boolean,
    timeZone: string,
    twoFactorAuthentication: boolean
  }

export interface settingPutUserRequest{
    name: string,
    email: string,
    dateOfBirth:string,
    permanentAddress: string,
    postalCode: string,
    username: string,
    presentAddress: string,
    city: string,
    country: string,
    profilePicture: string
  }


export interface settingPutUserResponse{
  data:{
    id: string,
    name: string,
    email: string,
    dateOfBirth:Date,
    permanentAddress: string,
    postalCode: string,
    username: string,
    presentAddress: string,
    city: string,
    country: string,
    profilePicture: string,
    accountBalance: number,
    role: string,
    preference:Preference,
  }
    }

    export interface RandomInvestmentData{
      data:{
        totalInvestment:number,
        rateOfReturn:number
        yearlyTotalInvestment:[
          {
            time:string
            value:number
          }
        ]
        monthlyRevenue:[
          {
            time:string
            value:number
          }
        ]
      }
    }