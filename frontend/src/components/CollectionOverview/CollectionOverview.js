import React from 'react';
import './CollectionOverview.scss';
import PreviewCollection from "../PreviewCollection/PreviewCollection";
import {useSelector} from "react-redux";

const CollectionOverview = () => {
    const {data} = useSelector(state => state.shop);

    return (
        <div className="CollectionOverview">
            {
                data.map(({id, ...otherCollectionProps}) => (
                    <PreviewCollection key={id} {...otherCollectionProps} />
                ))

            }
        </div>
    );
};

export default CollectionOverview;
