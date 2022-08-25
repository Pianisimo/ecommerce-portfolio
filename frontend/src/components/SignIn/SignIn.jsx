import React, {useState} from 'react';
import './SignIn.scss';
import FormInput from "../FormInput/FormInput";
import CustomButton from "../CustomButton/CustomButton";
import {useDispatch} from "react-redux";
import {logIn} from "../../redux/user.slice";
import {Snackbar} from "@mui/material";
import {SNACKBAR_SHORT_DURATION} from "../../utils/Constants";

const SignIn = () => {
    const dispatch = useDispatch();
    const [formData, setFormData] = useState({
        email: '',
        password: '',
    });

    const {email, password} = formData;

    const [snackbar, setSnackbar] = useState({
        action: false,
        message: '',
    })

    const handleClose = () => {
        setSnackbar({action: false, message: ''});
    };

    const handleSubmit = event => {
        event.preventDefault();
        dispatch(logIn({email, password}))
            .then(value => {
                if (value.type === 'users/login/rejected') {
                    setSnackbar({
                        action: true,
                        message: value.payload,
                    })
                }
            })
            .catch(reason => setSnackbar({
                action: true,
                message: reason.message,
            }))
    }

    const handleOnChange = event => {
        setFormData({...formData, [event.target.name]: event.target.value});
    }

    return (
        <div className='SignIn'>
            <h2>I already have an account</h2>
            <span>Sign in with your email and password</span>

            <form onSubmit={handleSubmit}>
                <FormInput handleChange={handleOnChange} name='email' type='email' value={email}
                    // required
                           label='email'/>
                <FormInput handleChange={handleOnChange} name='password' type='password'
                           value={password}
                    // required
                           label='password'/>
                <CustomButton type='submit'> Sign in </CustomButton>
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

export default SignIn;
