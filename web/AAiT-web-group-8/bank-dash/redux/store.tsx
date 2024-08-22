import { configureStore } from '@reduxjs/toolkit';
import { endpointsApi } from './api/endpoints';

const store = configureStore({
  reducer: {
    [endpointsApi.reducerPath]: endpointsApi.reducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(endpointsApi.middleware),
});

export default store;
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
