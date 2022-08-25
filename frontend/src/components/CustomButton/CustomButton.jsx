import React from 'react';
import './CustomButton.scss';

const CustomButton = ({children, inverted, ...otherProps}) => (
    <button className={`${inverted ? 'inverted' : ''} CustomButton`} {...otherProps}>
        {children}
    </button>
)

export default CustomButton;
