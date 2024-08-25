import { createSlice } from '@reduxjs/toolkit';

const initialState = {
  hamburgerMenu: false,
};

export const uiSlice = createSlice({
  name: 'ui',
  initialState,
  reducers: {
    toggleHamburgerMenu: (state) => {
      state.hamburgerMenu = !state.hamburgerMenu;
    },
    hideMenu: (state) => {
      state.hamburgerMenu = false;
    },
  },
});

export const { toggleHamburgerMenu, hideMenu } = uiSlice.actions;
export default uiSlice.reducer;
