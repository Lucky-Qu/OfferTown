import api from "./index.ts";

/**
 * 发送注册请求。
 * @param username 需要注册的用户名
 * @param password 注册时填写的密码
 */
const userRegister = async (username: string, password: string) => {
    const response = await api.post("/user/register", {username, password});
    return response.data;
};

/**
 * 发送登录请求。
 * @param username 登录用户名
 * @param password 登录密码
 */
const userLogin = async (username: string, password: string) => {
    const response = await api.post("/user/login", {username, password});
    return response.data;
};

export {userRegister, userLogin};
