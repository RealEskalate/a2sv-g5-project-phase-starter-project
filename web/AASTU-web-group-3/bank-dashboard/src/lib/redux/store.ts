import { configureStore } from "@reduxjs/toolkit";
import { setupListeners } from "@reduxjs/toolkit/query";
import { authApi } from "./api/authApi";
import { transactionsApi } from "./api/transactionsApi";
import { loansApi } from "./api/loansApi";
import transactionsReducer from "./slices/transactionsSlice";
import loansReducer from "./slices/loansSlice";

export const store = configureStore({
  reducer: {
    [authApi.reducerPath]: authApi.reducer,
    [transactionsApi.reducerPath]: transactionsApi.reducer,
    [loansApi.reducerPath]: loansApi.reducer,
    transactions: transactionsReducer, 
    loans: loansReducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(authApi.middleware, transactionsApi.middleware, loansApi.middleware),
});

setupListeners(store.dispatch);

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export default store;
