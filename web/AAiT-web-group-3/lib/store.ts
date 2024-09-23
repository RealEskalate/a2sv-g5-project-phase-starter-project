import { setupListeners } from "@reduxjs/toolkit/query";
import { configureStore } from "@reduxjs/toolkit";
import { TypedUseSelectorHook, useSelector } from "react-redux";

import { dashboardAPI } from "./services/dashboard";
import { loanAPI } from "./services/loans";
import { serviceAPI } from "./services/services";
import { settingAPI } from "./services/settings";
import { transactionAPI } from "./services/transactions";

import navigationReducer from "@/lib/features/navigation/navigationSlice";
import authReducer from "@/lib/features/auth/authSlice";

export const store = configureStore({
  // TODO : Add your reducers here.
  reducer: {
    navigationReducer,
    authReducer,
    [dashboardAPI.reducerPath]: dashboardAPI.reducer,
    [loanAPI.reducerPath]: loanAPI.reducer,
    [serviceAPI.reducerPath]: serviceAPI.reducer,
    [settingAPI.reducerPath]: settingAPI.reducer,
    [transactionAPI.reducerPath]: transactionAPI.reducer,
  },
  // TODO: concatenate your middlewares here
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({}).concat([
      dashboardAPI.middleware,
      loanAPI.middleware,
      serviceAPI.middleware,
      settingAPI.middleware,
      transactionAPI.middleware,
    ]),
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;

setupListeners(store.dispatch);
