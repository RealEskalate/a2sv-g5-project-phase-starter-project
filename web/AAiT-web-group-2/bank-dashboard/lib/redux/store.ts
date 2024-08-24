import { bankApi } from '@/lib/redux/api/bankApi';
import { configureStore } from '@reduxjs/toolkit'


export const store = configureStore({
  reducer: {
    [bankApi.reducerPath]: bankApi.reducer
  },
  middleware: getDefaultMiddleware =>
    getDefaultMiddleware()
      .concat(bankApi.middleware)
})


export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;