import './App.css';
import HomePage from "./pages/homepage/HomePage";
import {Redirect, Route, Switch} from "react-router-dom";
import Shop from "./pages/Shop/Shop";
import Header from "./components/Header/Header";
import SignInSignUp from "./pages/SignInSignUp/SignInSignUp";
import {useDispatch, useSelector} from "react-redux";
import {isAuth} from "./redux/user.slice";
import Checkout from "./pages/Checkout/Checkout";
import {useEffect} from "react";

const App = () => {
    const dispatch = useDispatch();
    const {currentUser} = useSelector(state => state.user);

    useEffect(() => {
        const dispatchFunc = async () => {
            try {
                const response = await dispatch(isAuth({}))
                if (response.type === 'users/isauth/rejected') {
                    console.log(response.payload)
                }
            } catch (e) {
                console.log(e)
            }
        }

        dispatchFunc();
    }, [])

    return (
        <div>
            <Header/>
            <Switch>
                <Route exact path='/' component={HomePage}/>
                <Route path='/shop' component={Shop}/>
                <Route exact path='/checkout' component={Checkout}/>
                <Route exact path='/signin' render={() => {
                    return currentUser
                        ? (<Redirect to='/'/>)
                        : (<SignInSignUp/>)
                }}/>
            </Switch>
        </div>
    );
}

export default App;
