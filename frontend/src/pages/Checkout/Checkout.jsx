import React from 'react';
import './Checkout.scss';
import {useSelector} from "react-redux";
import CheckoutItem from "../../components/CheckoutItem/CheckoutItem";
import StripeButton from "../../components/StripeButton/StripeButton";

const Checkout = () => {
    const {cartItems} = useSelector(state => state.cart);
    const total = cartItems.reduce((totalQuantity, cartItem) => {
        return totalQuantity + (cartItem.quantity * cartItem.price)
    }, 0)

    return (
        <div className='Checkout'>
            <div className='checkout-header'>
                <div className='header-block'>
                    <span>Product</span>
                </div>
                <div className='header-block'>
                    <span>Description</span>
                </div>
                <div className='header-block'>
                    <span>Quantity</span>
                </div>
                <div className='header-block'>
                    <span>Price</span>
                </div>
                <div className='header-block'>
                    <span>Remove</span>
                </div>
            </div>
            {cartItems.map(cartItem => (
                <CheckoutItem key={cartItem.id} cartItem={cartItem}/>
            ))}
            <div className='total'>TOTAL: ${total}</div>
            <StripeButton price={total} />
        </div>
    );
};

export default Checkout;
