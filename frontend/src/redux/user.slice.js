import {createAsyncThunk, createSlice} from '@reduxjs/toolkit'
import {BACKEND_URL} from "../index";

const initialState = {
    currentUser: null,
    shouldFetchAuth: true,
}

const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(isAuth.fulfilled, (state, action) => {
                state.currentUser = action.payload
                state.shouldFetchAuth = false
            })
            .addCase(isAuth.rejected, (state, action) => {
                state.currentUser = null
                state.shouldFetchAuth = false
            })
            .addCase(logOut.fulfilled, (state, action) => {
                state.currentUser = null
            })
            .addCase(logOut.rejected, (state, action) => {
                state.currentUser = null;
            })
            .addCase(logIn.fulfilled, (state, action) => {
                state.currentUser = action.payload
            })
            .addCase(logIn.rejected, (state, action) => {
                state.currentUser = null
            })
            .addCase(signUp.fulfilled, (state, action) => {
                state.currentUser = action.payload
            })
            .addCase(signUp.rejected, (state, action) => {
                state.currentUser = null
            })
    },
})

export const isAuth = createAsyncThunk(
    'users/isauth',
    async ({}, thunkAPI) => {
        const requestOptions = {
            method: 'GET',
            credentials: 'include',
        };

        try {
            const response = await fetch(BACKEND_URL + '/users/auth', requestOptions);
            if (response.status === 200) {
                const data = await response.json();
                return data
            } else {
                const data = await response.text();
                return thunkAPI.rejectWithValue(data)
            }
        } catch (e) {
            return thunkAPI.rejectWithValue(e)
        }
    }
)

export const logOut = createAsyncThunk(
    'users/logout',
    async ({}, thunkAPI) => {
        const requestOptions = {
            method: 'GET',
            credentials: 'include',
        };

        try {
            const response = await fetch(BACKEND_URL + '/users/logout', requestOptions);
            if (response.status === 200) {
                const data = await response.text();
                return data
            } else {
                const data = await response.text();
                return thunkAPI.rejectWithValue(data)
            }
        } catch (e) {
            return thunkAPI.rejectWithValue(e)
        }

    }
)

export const logIn = createAsyncThunk(
    'users/login',
    async ({email, password}, thunkAPI) => {
        const payload = {email, password}
        const requestOptions = {
            method: 'POST',
            credentials: 'include',
            body: JSON.stringify(payload)
        };

        try {
            const response = await fetch(BACKEND_URL + '/users/login', requestOptions);
            if (response.status === 200) {
                const data = await response.json();
                return data
            } else {
                const data = await response.text();
                return thunkAPI.rejectWithValue(data)
            }
        } catch (e) {
            return thunkAPI.rejectWithValue(e)
        }

    }
)

export const signUp = createAsyncThunk(
    'users/signup',
    async ({firstName, lastName, email, password}, thunkAPI) => {
        const payload = {
            first_name: firstName,
            last_name: lastName,
            email,
            password
        }
        const requestOptions = {
            method: 'POST',
            credentials: 'include',
            body: JSON.stringify(payload)
        };

        try {
            const response = await fetch(BACKEND_URL + '/users/signup', requestOptions);
            if (response.status === 201) {
                const data = await response.json();
                return data
            } else {
                const data = await response.text();
                return thunkAPI.rejectWithValue(data)
            }
        } catch (e) {
            return thunkAPI.rejectWithValue(e)
        }
    }
)

// export const {} = userSlice.actions
export default userSlice.reducer