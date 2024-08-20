import { configureStore, combineReducers } from "@reduxjs/toolkit";
import { persistStore, persistReducer } from "redux-persist";
import storage from "redux-persist/lib/storage";
import { useDispatch, useSelector, TypedUseSelectorHook } from "react-redux";
import {
  FLUSH,
  REHYDRATE,
  PAUSE,
  PERSIST,
  PURGE,
  REGISTER,
} from "redux-persist";
import cardReducer from "../slices/cardSlice"; // Import the cardSlice reducer
import transactionReducer from "../slices/TransactionSlice";
import userReducer from "../slices/userSlice";
import darkModeReducer from "../slices/darkModeSlice";
const rootReducer = combineReducers({
  cards: cardReducer,
  transactions: transactionReducer,
  user: userReducer,
  darkMode: darkModeReducer,
  // Add other reducers here if needed
});

const persistConfig = {
  key: "root",
  storage,
  whitelist: ["cards", "transactions", "darkMode"], // Persist the cards slice
};

const persistedReducer = persistReducer(persistConfig, rootReducer);

export const store = configureStore({
  reducer: persistedReducer,
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: {
        ignoredActions: [FLUSH, REHYDRATE, PAUSE, PERSIST, PURGE, REGISTER],
      },
    }),
});

export const persistor = persistStore(store);

// defining type
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

// creating custom hooks for useDispatch and useSelector
export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;
