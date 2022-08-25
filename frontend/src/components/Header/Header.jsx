import React from 'react';
import './Header.scss';
import {Link} from "react-router-dom";
import {ReactComponent as Logo} from "../../assets/crown.svg";
import {useDispatch, useSelector} from "react-redux";
import {logOut} from "../../redux/user.slice";
import CartIcon from "../CartIcon/CartIcon";
import CartDropdown from "../CartDropdown/CartDropdown";

const Header = () => {
    const dispatch = useDispatch();
    const { currentUser } = useSelector(state => state.user);

    return (
        <div className="Header">
            <Link to='/' className='logo-container'>
                <Logo className='logo'/>
            </Link>

            <div className='options'>
                <Link className='option' to='/shop'>
                    SHOP
                </Link>
                <Link className='option' to='/shop'>
                    CONTACT
                </Link>

                {
                    currentUser
                        ? <div className='option' onClick={() => {
                            dispatch(logOut({}))
                        }}>SIGN OUT</div>
                        : <Link className='option' to='/signin'>SIGN IN</Link>
                }

                <CartIcon/>
            </div>
            <CartDropdown />
        </div>
    )
};

export default Header;
