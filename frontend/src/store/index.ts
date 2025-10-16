import {configureStore} from "@reduxjs/toolkit";
import {authSliceReducer, userSliceReducer} from "./modules/user";

/**
 * 应用全局状态仓库。
 * 将认证信息与用户信息模块组合在一起。
 */
const store = configureStore({
    reducer: {
        auth: authSliceReducer,
        user: userSliceReducer,
    },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export default store;
