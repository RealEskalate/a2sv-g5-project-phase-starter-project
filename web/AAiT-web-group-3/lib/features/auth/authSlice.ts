import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { authValue } from "@/types";

const initialState: {
  value: authValue;
} = {
  value: {
    accessToken: "",
    errorMessage: "",
    isLoading: false,
  } as authValue,
};

const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    updateErrorMessage(state, action: PayloadAction<string | null>) {
      state.value.errorMessage = action.payload;
    },
    updateIsLoading(state, action: PayloadAction<boolean>) {
      state.value.isLoading = action.payload;
    },
    updateAccessToken(state, action: PayloadAction<string>) {
      state.value.accessToken = action.payload;
    },
  },
});

export const { updateErrorMessage, updateIsLoading, updateAccessToken } =
  authSlice.actions;

export default authSlice.reducer;
