import { createSlice } from '@reduxjs/toolkit';
import User from '../../../type/user';
const initialState:User= {
  id: '',
  name: '',
  email: '',
  dateOfBirth: '',
  permanentAddress: '',
  postalCode: '',
  username: '',
  presentAddress: '',
  city: '',
  country: '',
  profilePicture: '',
  accountBalance: 0,
  role: '',
  preference: {
    currency: '',
    sentOrReceiveDigitalCurrency: false,
    receiveMerchantOrder: false,
    accountRecommendations: false,
    timeZone: '',
    twoFactorAuthentication: false,
  },
  
};

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setUser(state, action) {
      console.log('Setting user:', action.payload); // Log the payload
      return { ...state, ...action.payload.data }; // Access the `data` property
    },
    resetUser(state) {
      return initialState;
    }
  }
});


export const { setUser, resetUser } = userSlice.actions;
export default userSlice.reducer;
