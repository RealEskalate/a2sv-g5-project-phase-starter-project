import { configureStore, applyMiddleware } from "@reduxjs/toolkit";
import { persistStore, persistReducer } from "redux-persist";
import storage from "redux-persist/lib/storage"; 
import { Persistor } from "redux-persist/es/types";
import userReducer from "./slice/userSlice";
import themeReducer from "./slice/themeSlice"; // Update the path to your slice
import createSagaMiddleware from "redux-saga";
import userSaga from "./userSaga";

const persistConfig = {
	key: "root",
	storage,
};

const sagaMiddleware = createSagaMiddleware();
const persistedReducer = persistReducer(persistConfig, themeReducer);

export const store = configureStore({
	reducer: {
		user: userReducer,
		theme: persistedReducer,
	},
	middleware: (getDefaultMiddleware) =>
		getDefaultMiddleware({
			serializableCheck: false,
		}).concat(sagaMiddleware),
});

export const persistor: Persistor = persistStore(store);
sagaMiddleware.run(userSaga);

export type RootState = ReturnType<typeof store.getState>;

export default store;