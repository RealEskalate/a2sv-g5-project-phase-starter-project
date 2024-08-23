import { configureStore } from "@reduxjs/toolkit";
import { setupListeners } from "@reduxjs/toolkit/query";
import { authApi } from "./api/authApi";
import { transactionsApi } from "./api/transactionsApi";
import { serviceApi } from "./api/serviceApi";
import { settingApi } from "./api/settingApi";
import { loansApi } from "./api/loansApi";
import { cardsApi } from './api/cardsApi';
import transactionsReducer from "./slices/transactionsSlice";
import loansReducer from "./slices/loansSlice";
import settingReducer from "./slices/settingSlice";
import serviceReducer from "./slices/serviceSlice";
import cardsReducer from './slices/cardsSlice';
import layoutReducer from './slices/layoutSlice'; // Import the layout reducer

export const store = configureStore({
  reducer: {
    [authApi.reducerPath]: authApi.reducer,
    [transactionsApi.reducerPath]: transactionsApi.reducer,
    [serviceApi.reducerPath]: serviceApi.reducer,
    [settingApi.reducerPath]: settingApi.reducer,
    [loansApi.reducerPath]: loansApi.reducer,
    [cardsApi.reducerPath]: cardsApi.reducer,
    transactions: transactionsReducer,
    service: serviceReducer,
    setting: settingReducer,
    loans: loansReducer,
    cards: cardsReducer,
    layout: layoutReducer, // Add the layout reducer to the store
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(
      authApi.middleware,
      transactionsApi.middleware,
      serviceApi.middleware,
      settingApi.middleware,
      loansApi.middleware,
      cardsApi.middleware
    ),
});

setupListeners(store.dispatch);

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export default store;
