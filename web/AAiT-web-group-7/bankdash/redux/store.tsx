import { configureStore } from '@reduxjs/toolkit';
import { endpointsApiFunc } from './api/endpoints';

const store = configureStore({
  reducer: {
    [endpointsApiFunc.reducerPath]: endpointsApiFunc.reducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(endpointsApiFunc.middleware),
});

export default store;
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
