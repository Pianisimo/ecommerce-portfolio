import React from 'react';
import './CartDropdown.scss';
import CustomButton from "../CustomButton/CustomButton";
import {useDispatch, useSelector} from "react-redux";
import CartItem from "../CartItem/CartItem";
import {withRouter} from "react-router-dom";
import {changeToOpposite} from "../../redux/cart.slice";

const CartDropdown = ({history}) => {
    const dispatch = useDispatch();
    const {hidden, cartItems} = useSelector(state => state.cart);

    return (
        <div>
            {hidden
                ? (<div/>)
                : (<div className="CartDropdown">
                    <div className='cart-items'>
                        {
                            cartItems.length
                                ? cartItems.map(cartItem => {
                                    return (
                                        <CartItem key={cartItem.id} item={cartItem}/>
                                    )
                                })
                                : <span className='empty-message'>Your cart is empty</span>
                        }
                    </div>
                    <CustomButton onClick={() => {
                        dispatch(changeToOpposite())
                        return history.push('/checkout');
                    }}>
                        GO TO CHECKOUT
                    </CustomButton>
                </div>)
            }
        </div>
    )
};

export default withRouter(CartDropdown);
