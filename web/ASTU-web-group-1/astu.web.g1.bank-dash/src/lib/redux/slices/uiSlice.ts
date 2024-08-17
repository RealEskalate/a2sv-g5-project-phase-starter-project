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
  },
});

export const { toggleHamburgerMenu } = uiSlice.actions;
export default uiSlice.reducer;
