import React from 'react';
import './CartIcon.scss';
import {ReactComponent as ShoppingIcon} from '../../assets/shopping-bag.svg';
import {useDispatch, useSelector} from "react-redux";
import {changeToOpposite} from "../../redux/cart.slice";

const CartIcon = () => {
    const dispatch = useDispatch();
    const {cartItems} = useSelector(state => state.cart);

    return (
        <div className="CartIcon" onClick={() => {
            dispatch(changeToOpposite())
        }
        }>
            <ShoppingIcon className='shopping-icon'/>
            <span className='item-count'>{
                cartItems.reduce((totalQuantity, cartItem) => {
                    return totalQuantity + cartItem.quantity
                }, 0)
            }</span>
        </div>
    );
};

export default CartIcon;
