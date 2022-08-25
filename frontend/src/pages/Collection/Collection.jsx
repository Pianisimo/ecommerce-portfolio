import React from 'react';
import './Collection.scss';
import CollectionItem from "../../components/CollectionItem/CollectionItem";
import {useSelector} from "react-redux";

const Collection = ({match}) => {
    const {data} = useSelector(state => state.shop);

    return (
        <div className="Collection">
            <h2 className='title'>{match.params.collectionId.toUpperCase()}</h2>
            <div className='items'>
                {
                    data
                        .filter((item) => {
                            return match.params.collectionId.toLowerCase() === item.title.toLowerCase()
                        })[0].items
                        .map((item) => {
                            return (
                                <CollectionItem key={item.id} item={item}/>
                            );
                        })
                }
            </div>
        </div>
    );
};

export default Collection;
