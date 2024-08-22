import { configureStore } from "@reduxjs/toolkit";
import {  persistStore, persistReducer } from "redux-persist";
import storage from "redux-persist/lib/storage"; // or 'redux-persist/lib/storage/session' for session storage
import { Persistor } from "redux-persist/es/types";
import userReducer from "./slice/userSlice";
import themeReducer from "./slice/themeSlice"; // Update the path to your slice

const persistConfig = {
	key: "root",
	storage,
};

const persistedReducer = persistReducer(persistConfig, themeReducer);

export const store = configureStore({
	reducer: {
		user: userReducer,
		theme: persistedReducer,
	},
	// other store configurations if needed
});

export const persistor: Persistor = persistStore(store);
