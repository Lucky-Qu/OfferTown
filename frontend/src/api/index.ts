import axios from "axios";
import store from "../store";

/**
 * Axios 实例，携带统一的基础配置（基地址、超时时间等）供全局复用。
 */
const api = axios.create({
    baseURL: "http://localhost:7777",
    timeout: 5000,
});

/**
 * 请求拦截器：在请求发出前自动附加 JWT Token。
 */
api.interceptors.request.use(config => {
    const token = store.getState().auth.token;
    if (token && config.headers) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
});

export default api;
