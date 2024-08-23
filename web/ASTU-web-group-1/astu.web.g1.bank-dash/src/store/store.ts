import { activeLoanApi } from "@/lib/redux/slices/activeLoanSlice";
import { bankServiceApi } from "@/lib/redux/slices/bankService";
import { cardApi } from "@/lib/redux/slices/cardSlice";
import { transactionApi } from "@/lib/redux/slices/transactionSlice";
import uiSlice from "@/lib/redux/slices/uiSlice";
import { configureStore } from "@reduxjs/toolkit";

export const store = configureStore({
  reducer: {
    ui: uiSlice,
    [cardApi.reducerPath]: cardApi.reducer,
    [transactionApi.reducerPath]: transactionApi.reducer,
    [activeLoanApi.reducerPath]: activeLoanApi.reducer,
    [bankServiceApi.reducerPath]: bankServiceApi.reducer,
  },
  middleware: (getDefaultMiddleware) => {
    return getDefaultMiddleware()
      .concat(cardApi.middleware)
      .concat(transactionApi.middleware)
      .concat(activeLoanApi.middleware)
      .concat(bankServiceApi.middleware);
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;