import {configureStore} from "@reduxjs/toolkit";
import {authSliceReducer, userSliceReducer} from "./modules/user"


const store = configureStore(
    {
        reducer:{
            auth: authSliceReducer,
            user: userSliceReducer,
        }
    }
)

export default store