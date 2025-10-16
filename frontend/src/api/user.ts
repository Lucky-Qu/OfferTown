import api from "./index.ts";

const userRegister = async (username: string, password: string) => {
    const response = await api.post("/user/register", {username, password})
    return response.data;
}

const userLogin = async (username: string, password: string) => {
    const response = await api.post("/user/login", {username, password})
    return response.data;
}


export {userRegister, userLogin}