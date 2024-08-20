import { configureStore } from "@reduxjs/toolkit";
import { setupListeners } from "@reduxjs/toolkit/query";
import { authApi } from "./api/authApi";
import { transactionsApi } from "./api/transactionsApi";
import { serviceApi } from "./api/serviceApi";
import { settingApi } from "./api/settingApi";
import { loansApi } from "./api/loansApi";
import transactionsReducer from "./slices/transactionsSlice";
import loansReducer from "./slices/loansSlice";
import settingReducer from "./slices/settingSlice";
import serviceReducer from "./slices/serviceSlice";

export const store = configureStore({
  reducer: {
    [authApi.reducerPath]: authApi.reducer,
    [transactionsApi.reducerPath]: transactionsApi.reducer,
    [serviceApi.reducerPath]: serviceApi.reducer,
    [settingApi.reducerPath]: settingApi.reducer,
    [loansApi.reducerPath]: loansApi.reducer,
    transactions: transactionsReducer, 
    service: serviceReducer, 
    setting: settingReducer, 
    loans: loansReducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(
      authApi.middleware,
      transactionsApi.middleware,
      serviceApi.middleware,
      settingApi.middleware,
      loansApi.middleware
    ),
});

setupListeners(store.dispatch);

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export default store;
