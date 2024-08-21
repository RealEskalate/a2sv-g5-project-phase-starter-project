import { configureStore } from "@reduxjs/toolkit";
import userReducer from "./features/userSlice/userSlice";
import { transactionApi } from "../lib/service/TransactionService";
import { bankApi } from "../lib/service/BankService";
import { userApi } from "../lib/service/UserService";
import { loanApi } from "../lib/service/LoanService";
import { companyApi } from "../lib/service/CompanyService";
import { CreditCardInfoApi } from "../lib/service/CardService";
export const store = configureStore({
  reducer: {
    user: userReducer, // Add the user reducer
    [transactionApi.reducerPath]: transactionApi.reducer,
    [bankApi.reducerPath]: bankApi.reducer,
    [userApi.reducerPath]: userApi.reducer,
    [loanApi.reducerPath]: loanApi.reducer,
    [companyApi.reducerPath]: companyApi.reducer,
    [CreditCardInfoApi.reducerPath]: CreditCardInfoApi.reducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(
      transactionApi.middleware,
      bankApi.middleware,
      userApi.middleware,
      loanApi.middleware,
      companyApi.middleware,
      CreditCardInfoApi.middleware
    ),
});

export type AppStore = typeof store;
export type RootState = ReturnType<AppStore["getState"]>;
export type AppDispatch = AppStore["dispatch"];
