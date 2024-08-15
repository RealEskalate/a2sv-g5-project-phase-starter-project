import { configureStore } from 'react-redux'
import authReducer from './features/auth-slice'

export const store = configureStore ({
    reducer: {
        authReducer,

    },
});

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch