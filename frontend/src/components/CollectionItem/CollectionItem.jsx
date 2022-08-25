import React from 'react';
import './CollectionItem.scss';
import CustomButton from "../CustomButton/CustomButton";
import {useDispatch} from "react-redux";
import {addItem} from "../../redux/cart.slice";

const CollectionItem = ({item}) => {
    const {id, name, price, imageUrl} = item
    const dispatch = useDispatch();

    return (
        <div className="CollectionItem">
            <div
                className='image'
                style={{
                    backgroundImage: `url(${imageUrl})`
                }}>
            </div>
            <div className='collection-footer'>
                <span className='name'>{name}</span>
                <span className='price'>{price}</span>
            </div>
            <CustomButton inverted onClick={() => {
                dispatch(addItem({...item}))
            }
            }>ADD TO CART</CustomButton>
        </div>
    );
};

export default CollectionItem;
