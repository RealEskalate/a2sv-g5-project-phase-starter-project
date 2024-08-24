import { UserDataType } from "@/types/user.types";
import { createSlice } from "@reduxjs/toolkit";

const initialState: UserDataType = {
  id: "",
  name: "",
  email: "",
  dateOfBirth: "",
  permanentAddress: "",
  postalCode: "",
  username: "",
  presentAddress: "",
  city: "",
  country: "",
  profilePicture: "",
  accountBalance: 0,
  role: "",
  preference: {
    currency: "",
    sentOrReceiveDigitalCurrency: false,
    receiveMerchantOrder: false,
    accountRecommendations: false,
    timeZone: "",
    twoFactorAuthentication: false,
  },
};

export const profileSlice = createSlice({
  name: "profile",
  initialState,
  reducers: {
    setProfile: (state, action) => {
      return { ...state, ...action.payload };
    },
    setPreferences: (state, action) => {
      return {
        ...state,
        preference: {
          ...state.preference,
          ...action.payload,
        },
      };
    },
  },
});

export const { setProfile } = profileSlice.actions;
export default profileSlice.reducer