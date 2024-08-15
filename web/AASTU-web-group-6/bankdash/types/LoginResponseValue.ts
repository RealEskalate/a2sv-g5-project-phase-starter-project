interface InnerData{
    [key: string]: any;
}

interface LoginResponseDataValue{
    access_token: string;
    refresh_token : string;
    data : InnerData;
  }

interface LoginResponseValue{

    success : boolean;
    message : string;
    data : LoginResponseDataValue;
}

export default LoginResponseValue