import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Service } from "../types/service"; 

interface serviceState {
  service: Service[];
  loading: boolean;
  error: string | null;
}

const initialState: serviceState = {
  service: [],
  loading: false,
  error: null,
};

const serviceSlice = createSlice({
  name: "service",
  initialState,
  reducers: {
    setService: (state, action: PayloadAction<Service[]>) => {
      state.service = action.payload;
    },
    setLoading: (state, action: PayloadAction<boolean>) => {
      state.loading = action.payload;
    },
    setError: (state, action: PayloadAction<string | null>) => {
      state.error = action.payload;
    },
  },
});

export const { setService, setLoading, setError } = serviceSlice.actions;
export default serviceSlice.reducer;
