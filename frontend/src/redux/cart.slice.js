import {createSlice} from '@reduxjs/toolkit'

const initialState = {
    hidden: true,
    cartItems: []
}

const cartSlice = createSlice({
    name: 'cart',
    initialState,
    reducers: {
        changeToOpposite(state) {
            state.hidden = !state.hidden
        },

        changeTo(state, action) {
            state.hidden = action.payload
        },

        addItem(state, action) {
            const existingCartItem = state.cartItems.find(cartItem => {
                return cartItem.id === action.payload.id;
            })

            if (existingCartItem) {
                state.cartItems = state.cartItems.map(cartItem => {
                    return cartItem.id === action.payload.id
                        ? {...cartItem, quantity: cartItem.quantity + 1}
                        : cartItem
                })
            } else {
                state.cartItems = [...state.cartItems, {...action.payload, quantity: 1}]
            }
        },

        removeItem(state, action) {
            const helper = state.cartItems.map((cartItem) => {
                if (cartItem.id === action.payload.id) {
                    return {...cartItem, quantity: cartItem.quantity - 1}
                } else {
                    return {...cartItem}
                }
            })

            state.cartItems = helper.filter(cartItem => {
                return cartItem.quantity !== 0;
            })
        },

        clearItemFromCart(state, action) {
            state.cartItems = state.cartItems.filter(cartItem => {
                return cartItem.id !== action.payload;
            })
        },
    },
})


export const {changeToOpposite, changeTo, addItem, removeItem, clearItemFromCart} = cartSlice.actions
export default cartSlice.reducer
