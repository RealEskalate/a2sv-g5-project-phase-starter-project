import UserPreferenceValue from "./UserPreferenceValue";

  interface UserValue {
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
    preference: UserPreferenceValue;
  }
  
  export default UserValue;
  