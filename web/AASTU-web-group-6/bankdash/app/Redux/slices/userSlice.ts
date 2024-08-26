import { createSlice, PayloadAction, createAsyncThunk } from "@reduxjs/toolkit";
import UserValue from "@/types/UserValue";
import UserPreferenceValue from "@/types/UserPreferenceValue";

export interface YearlyInvestment {
  time: string;
  value: number;
}

export interface MonthlyRevenue {
  time: string;
  value: number;
}

export interface InvestmentData {
  totalInvestment: number;
  rateOfReturn: number;
  yearlyTotalInvestment: YearlyInvestment[];
  monthlyRevenue: MonthlyRevenue[];
}
interface UserState {
  user: UserValue | null;
  preferences: UserPreferenceValue | null;
  investment: InvestmentData;
  status: "idle" | "loading" | "succeeded" | "failed";
  error: string | null;
}

const initialState: UserState = {
  user: null,
  preferences: null,
  investment: {
    totalInvestment: 0,
    rateOfReturn: 0,
    yearlyTotalInvestment: [],
    monthlyRevenue: [],
  },
  status: "idle",
  error: null,
};

const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    setUserData(state, action: PayloadAction<UserValue>) {
      state.user = action.payload;
    },
    setPreferences(state, action: PayloadAction<UserPreferenceValue>) {
      state.preferences = action.payload;
    },
    setInvestment(state, action: PayloadAction<InvestmentData>) {
      state.investment = action.payload;
    },
    setStatus(state, action: PayloadAction<UserState["status"]>) {
      state.status = action.payload;
    },
    setError(state, action: PayloadAction<string | null>) {
      state.error = action.payload;
    },
    clearUserState(state) {
      state.user = null;
      state.preferences = null;
      state.status = "idle";
      state.error = null;
    },
  },
});

export const { setUserData, setPreferences, setInvestment, setStatus, setError, clearUserState } =
  userSlice.actions;

export default userSlice.reducer;

