import { configureStore } from "@reduxjs/toolkit";
import { setupListeners } from "@reduxjs/toolkit/query";
import menuReducer from './slices/menuSlice';
import formSlice from "./slices/formSlice";
import { authApi } from "./api/authApi";
import { serviceApi } from "./api/servicesApi";

export const store = configureStore({
    reducer: {
        [authApi.reducerPath]: authApi.reducer,
        [serviceApi.reducerPath]: serviceApi.reducer,
        form: formSlice,
        menu: menuReducer,
    },
    middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(authApi.middleware).concat(serviceApi.middleware),

})
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

setupListeners(store.dispatch)
export default store;
