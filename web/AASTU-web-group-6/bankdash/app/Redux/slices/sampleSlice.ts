import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  data: [],
  loading: false,
  error: null,
};

const sampleSlice = createSlice({
  name: "sampledata",
  initialState,
  reducers: {
    // write actions
    setSample: (state, action) => {
      state.data = action.payload;
    },
  },
});

//export actions
export const { setSample } = sampleSlice.actions;

//export reducer
export default sampleSlice.reducer;
