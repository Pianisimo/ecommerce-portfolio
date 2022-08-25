import React from 'react';
import './CheckoutItem.scss';
import {useDispatch} from "react-redux";
import {clearItemFromCart, addItem, removeItem} from "../../redux/cart.slice";

const CheckoutItem = ({cartItem}) => {
    const dispatch = useDispatch();
    const { id, name, imageUrl, price, quantity } = cartItem;

    console.log(id)

    return (
        <div className='CheckoutItem'>
            <div className='image-container'>
                <img src={imageUrl} alt='item' />
            </div>
            <span className='name'>{name}</span>
            <span className='quantity'>
                <div className='arrow' onClick={() => dispatch(removeItem(cartItem))}>&#10094;</div>
                <span className='value'>{quantity}</span>
                <div className='arrow' onClick={() => dispatch(addItem(cartItem))}>&#10095;</div>
            </span>
            <span className='price'>{price}</span>
            <div className='remove-button' onClick={() => dispatch(clearItemFromCart(id))}>
                &#10005;
            </div>
        </div>
    );
};

export default CheckoutItem;
