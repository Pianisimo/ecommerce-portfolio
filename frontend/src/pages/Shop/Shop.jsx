import React from 'react';
import './Shop.scss';
import CollectionOverview from "../../components/CollectionOverview/CollectionOverview";
import {Route} from "react-router-dom";
import Collection from "../Collection/Collection";

const Shop = ({match}) => {
    return (<div className='Shop'>
        <Route exact path={`${match.path}`} component={CollectionOverview}/>
        <Route path={`${match.path}/:collectionId`} component={Collection} />
    </div>);
}

export default Shop;
