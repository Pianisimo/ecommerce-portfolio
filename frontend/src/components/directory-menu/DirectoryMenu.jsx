import React from "react";
import './DirectoryMenu.scss'
import MenuItem from "../menu-item/MenuItem";
import {useSelector} from "react-redux";

const DirectoryMenu = () => {
    const {sections} = useSelector(state => state.directory);
    return (
        <div className='DirectoryMenu'>
            {
                sections.map(({id, ...otherSectionProps}) => (
                    <MenuItem key={id} {...otherSectionProps}/>
                ))
            }
        </div>
    )
}

export default DirectoryMenu;
