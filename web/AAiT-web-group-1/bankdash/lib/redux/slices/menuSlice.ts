import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface MenuState {
  isSidebarVisible: boolean;
}

const initialState: MenuState = {
  isSidebarVisible: false, // initial value
};

const menuSlice = createSlice({
  name: 'menu',
  initialState,
  reducers: {
    toggleSidebar: (state) => {
      state.isSidebarVisible = !state.isSidebarVisible;
    },
    setSidebarVisibility: (state, action: PayloadAction<boolean>) => {
      state.isSidebarVisible = action.payload;
    },
  },
});

export const { toggleSidebar, setSidebarVisibility } = menuSlice.actions;
export default menuSlice.reducer;
