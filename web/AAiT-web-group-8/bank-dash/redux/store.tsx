import { configureStore } from '@reduxjs/toolkit';
import { loanApi } from './api/active-loan-controller';
import { authApi } from './api/authentication-controller';
import { bankserviceApi } from './api/bank-service-controller';
import { cardApi } from './api/card-controller';
import { companyApi } from './api/company-controller';
import { userApi } from './api/user-controller';
import { transactionApi } from './api/transaction-controller';


const store = configureStore({
  reducer: {
    [loanApi.reducerPath]: loanApi.reducer,
    [authApi.reducerPath]: authApi.reducer,
    [userApi.reducerPath]: userApi.reducer,
    [cardApi.reducerPath]: cardApi.reducer,
    [transactionApi.reducerPath]: transactionApi.reducer,
    [bankserviceApi.reducerPath]: bankserviceApi.reducer,
    [companyApi.reducerPath]: companyApi.reducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(loanApi.middleware, authApi.middleware, userApi.middleware, cardApi.middleware, transactionApi.middleware, bankserviceApi.middleware, companyApi.middleware),
});

export default store;
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;