import { configureStore } from "@reduxjs/toolkit";
import { setupListeners } from "@reduxjs/toolkit/query";
import formSlice from "./slices/formSlice";
import { authApi } from "./api/authApi";

export const store = configureStore({
    reducer: {
        [authApi.reducerPath]: authApi.reducer,
        form: formSlice
    },
    middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(authApi.middleware),

})

setupListeners(store.dispatch)