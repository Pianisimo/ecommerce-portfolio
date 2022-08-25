import React from 'react';
import './SignInSignUp.scss';
import SignIn from "../../components/SignIn/SignIn";
import SignUp from "../../components/SignUp/SignUp";

const SignInSignUp = () => (
    <div className="SignInSignUp">
        <SignIn></SignIn>
        <SignUp></SignUp>
    </div>
);


export default SignInSignUp;
