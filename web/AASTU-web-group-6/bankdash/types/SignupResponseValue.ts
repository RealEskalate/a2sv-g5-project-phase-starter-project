import UserValue from "./UserValue";

  interface LoginResponseData {
    access_token: string;
    refresh_token: string;
    data: UserValue;
  }
  
  interface SignupResponseValue {
    success: boolean;
    message: string;
    data: LoginResponseData;
  }
  
  export default SignupResponseValue;
  