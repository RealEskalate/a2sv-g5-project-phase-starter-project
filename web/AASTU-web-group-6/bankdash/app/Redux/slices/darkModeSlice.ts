// store/darkModeSlice.ts
import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface DarkModeState {
  darkMode: boolean;
}

const initialState: DarkModeState = {
  //initially based on browser preference
  // darkMode: window.matchMedia("(prefers-color-scheme: dark)").matches,
  darkMode: false,
};

const darkModeSlice = createSlice({
  name: "darkMode",
  initialState,
  reducers: {
    toggleDarkMode(state) {
      state.darkMode = !state.darkMode;
    },
    setDarkMode(state, action: PayloadAction<boolean>) {
      state.darkMode = action.payload;
    },
  },
});

export const { toggleDarkMode, setDarkMode } = darkModeSlice.actions;

export default darkModeSlice.reducer;
