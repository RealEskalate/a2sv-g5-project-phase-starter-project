import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface LayoutState {
  ishidden: boolean;
  activeItem: string;
}

const initialState: LayoutState = {
  ishidden: false,
  activeItem: 'Dashboard',
};

const layoutSlice = createSlice({
  name: 'layout',
  initialState,
  reducers: {
    toggleSidebar: (state) => {
      state.ishidden = !state.ishidden;
    },
    setActiveItem: (state, action: PayloadAction<string>) => {
      state.activeItem = action.payload;
    },
  },
});

export const { toggleSidebar, setActiveItem } = layoutSlice.actions;

export default layoutSlice.reducer;
