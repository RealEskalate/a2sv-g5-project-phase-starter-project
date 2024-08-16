import { configureStore } from "@reduxjs/toolkit";
import { transactionApi } from "./service/TransactionService";

export const store = () => {
  return configureStore({
    reducer: {
      [transactionApi.reducerPath]: transactionApi.reducer,
      // Add your reducer path this way
    },
    middleware: (getDefaultMiddleWare) =>
      getDefaultMiddleWare().concat(transactionApi.middleware), // Add the middleware beside the transactionAPi.middleware by adding a comma
  });
};

export type AppStore = ReturnType<typeof store>;
export type RootState = ReturnType<AppStore["getState"]>;
export type AppDispatch = AppStore["dispatch"];
