import { configureStore } from "@reduxjs/toolkit";
import { transactionApi } from "./service/TransactionService";
import { bankApi } from "./service/BankService";
import { userApi } from "./service/UserService";
import { loanApi } from "./service/LoanService";
import { companyApi } from "./service/CompanyService";
import { investmentApi } from "./service/InvestmentServices";

export const store = () => {
  return configureStore({
    reducer: {
      [transactionApi.reducerPath]: transactionApi.reducer,
      [bankApi.reducerPath]: bankApi.reducer,
      [userApi.reducerPath]: userApi.reducer,
      [loanApi.reducerPath]: loanApi.reducer,
      [companyApi.reducerPath]: companyApi.reducer,
      // Add your reducer path this way
    },
    middleware: (getDefaultMiddleWare) =>
      getDefaultMiddleWare().concat(
        transactionApi.middleware,
        bankApi.middleware,
        userApi.middleware,
        loanApi.middleware,
        companyApi.middleware
      ), // Add the middleware beside the transactionAPi.middleware by adding a comma
  });
};

export type AppStore = ReturnType<typeof store>;
export type RootState = ReturnType<AppStore["getState"]>;
export type AppDispatch = AppStore["dispatch"];
