import {createSlice} from "@reduxjs/toolkit";

/**
 * 认证 Token 相关状态与 reducer。
 * 负责缓存当前用户的登录令牌。
 */
const authSlice = createSlice({
    name: "auth",
    initialState: {
        token: localStorage.getItem("token"),
    },
    reducers: {
        setToken: (state, action) => {
            state.token = action.payload;
        },
        clearToken: state => {
            state.token = "";
        },
    },
});

/**
 * 用户基本信息（目前仅用户名）的状态管理。
 */
const userSlice = createSlice({
    name: "user",
    initialState: {
        username: localStorage.getItem("username"),
    },
    reducers: {
        setUsername: (state, action) => {
            state.username = action.payload;
        },
        clearUsername: state => {
            state.username = "";
        },
    },
});

export const {setToken, clearToken} = authSlice.actions;
export const {setUsername, clearUsername} = userSlice.actions;
export const userSliceReducer = userSlice.reducer;
export const authSliceReducer = authSlice.reducer;
