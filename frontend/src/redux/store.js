import {logger} from "redux-logger/src";
import {configureStore} from "@reduxjs/toolkit";
import userSlice from "./user.slice";
import cartSlice from "./cart.slice";
import directorySlice from "./directory.slice";
import shopSlice from "./shop.slice";
import {persistReducer, persistStore} from 'redux-persist'
import storage from 'redux-persist/lib/storage'

const persistConfig = {
    key: 'root',
    storage,
}

const cartPersistReducer = persistReducer(persistConfig, cartSlice)

export const store = configureStore({
    reducer: {
        user: userSlice,
        cart: cartPersistReducer,
        directory: directorySlice,
        shop: shopSlice
    },
    middleware: (getDefaultMiddleware) => {
        if (process.env.NODE_ENV === 'development') {
            return getDefaultMiddleware().concat(logger);
        }

        return getDefaultMiddleware();
    },
})

export const persistor = persistStore(store)

export default { store, persistor };
