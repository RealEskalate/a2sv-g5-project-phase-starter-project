import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { RandomInvestmentData, settingPutUserResponse } from "../types/setting"; 

interface settingState {
  setting: settingPutUserResponse[];
  loading: boolean;
  error: string | null;
}

const initialState: settingState = {
  setting: [],
  loading: false,
  error: null,
};

const settingSlice = createSlice({
  name: "setting",
  initialState,
  reducers: {
    setSetting: (state, action: PayloadAction<settingPutUserResponse[]>) => {
      state.setting = action.payload;
    },
    setLoading: (state, action: PayloadAction<boolean>) => {
      state.loading = action.payload;
    },
    setError: (state, action: PayloadAction<string | null>) => {
      state.error = action.payload;
    },
  },
});

export const { setSetting, setLoading, setError} = settingSlice.actions;
export default settingSlice.reducer;
