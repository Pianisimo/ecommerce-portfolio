import React, {useState} from 'react';
import './SignUp.scss';

import FormInput from "../FormInput/FormInput";
import CustomButton from "../CustomButton/CustomButton";
import {useDispatch, useSelector} from "react-redux";
import {signUp} from "../../redux/user.slice";
import {SNACKBAR_SHORT_DURATION} from "../../utils/Constants";
import {CircularProgress, Snackbar} from "@mui/material";

const SignUp = () => {
    const dispatch = useDispatch();
    const [formData, setFormData] = useState({
        firstName: '',
        lastName: '',
        email: '',
        password: '',
        confirmPassword: '',
    });

    const {firstName, lastName, email, password, confirmPassword} = formData;

    const [snackbar, setSnackbar] = useState({
        action: false,
        message: '',
    })

    const {loading} = useSelector(state => state.user)

    const handleClose = () => {
        setSnackbar({action: false, message: ''});
    };

    const handleSubmit = async event => {
        event.preventDefault();

        if (password !== confirmPassword) {
            setSnackbar({
                action: true,
                message: 'Password does not match',
            })
            return;
        }

        dispatch(signUp({firstName, lastName, email, password}))
            .then(value => {
                if (value.type === 'users/signup/rejected') {
                    setSnackbar({
                        action: true,
                        message: value.payload,
                    })
                }
            })
            .catch(reason => setSnackbar({
                action: true,
                message: reason.payload,
            }))

    }

    const handleOnChange = event => {
        setFormData({...formData, [event.target.name]: event.target.value});
    }

    return (
        <div className='SignUp'>
            <h2 className='title'>I do not have an account</h2>
            <span>Sign up with your email and password</span>
            <form className='sign-up-form' onSubmit={handleSubmit}>
                <FormInput handleChange={handleOnChange} name='email' type='email' value={email}
                           required autocomplete="email"
                           label='email'/>
                <FormInput handleChange={handleOnChange} name='firstName' type='text'
                           value={firstName}
                           required autocomplete='given-name'
                           label='first name'/>
                <FormInput handleChange={handleOnChange} name='lastName' type='text'
                           value={lastName}
                           required autocomplete='family-name'
                           label='last name'/>
                <FormInput handleChange={handleOnChange} name='password' type='password'
                           value={password}
                           required autocomplete='new-password'
                           label='password'/>

                <FormInput handleChange={handleOnChange} name='confirmPassword' type='password'
                           value={confirmPassword}
                           required autocomplete='new-password'
                           label='retype password'/>
                <div className='buttons'>
                    {
                        loading
                            ? <CircularProgress />
                            : <CustomButton type='submit'> Sign up </CustomButton>
                    }
                </div>
            </form>
            <Snackbar
                open={snackbar.action}
                autoHideDuration={SNACKBAR_SHORT_DURATION}
                onClose={handleClose}
                message={snackbar.message}
                anchorOrigin={{horizontal: "center", vertical: "bottom"}}
            />
        </div>
    )
}

export default SignUp;
