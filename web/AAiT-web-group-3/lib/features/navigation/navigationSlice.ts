import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { navigationValue } from "@/types";

const initialState: { value: navigationValue } = {
  value: {
    activePage: "Dashboard",
    toggle: false,
  } as navigationValue,
};

const sidebarSlice = createSlice({
  name: "navigation",
  initialState,
  reducers: {
    updateActivePage(state, action: PayloadAction<string>) {
      state.value.activePage = action.payload;
    },
    updateToggle(state, action: PayloadAction<boolean>) {
      state.value.toggle = action.payload;
    },
  },
});

export const { updateActivePage,updateToggle } = sidebarSlice.actions;

export default sidebarSlice.reducer;
