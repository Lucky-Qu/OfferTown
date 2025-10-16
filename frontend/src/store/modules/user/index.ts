import {createSlice} from "@reduxjs/toolkit";


const authSlice = createSlice(
    {
        name: "auth",
        initialState: {
            token: localStorage.getItem("token")
        },
        reducers:{
            setToken: (state, action) => {
                state.token = action.payload
            },
            clearToken: (state) => {
                state.token = ""
            }
        }
    }
)

const userSlice = createSlice({
    name: "user",
    initialState: {
        username: localStorage.getItem("username"),
    },
    reducers: {
        setUsername: (state, action) => {
            state.username = action.payload
        },
        clearUsername: (state) => {
            state.username = ""
        }
    }
})
export const {setToken, clearToken} = authSlice.actions
export const {setUsername, clearUsername} = userSlice.actions
export const userSliceReducer = userSlice.reducer
export const authSliceReducer = authSlice.reducer