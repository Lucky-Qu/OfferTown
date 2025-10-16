import axios from "axios";
import store from "../store";

const api = axios.create({
    baseURL: "http://localhost:7777",
    timeout: 5000,
})
// 请求拦截器
api.interceptors.request.use(
    (config) => {
        const token = store.getState().auth.token
        if (token && config.headers) {
            config.headers.Authorization = `Bearer ${token}`
        }
        return config;
    },
)


export default api