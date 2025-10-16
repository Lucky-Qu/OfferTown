import {useState} from "react";
import {useNavigate} from "react-router";
import NaviBar from "../../components/NaviBar";
import style from "./LoginPage.module.css";
import {userLogin} from "../../api/user.ts";
import store from "../../store";
import {setToken, setUsername} from "../../store/modules/user";

/**
 * 调用后端登录接口并处理响应。
 * @param setErrMsg 用于更新错误描述的 setter
 */
const useLogin = (setErrMsg: (message: string) => void) => {
    const navigate = useNavigate();
    return async (username: string, password: string) => {
        try {
            const res = await userLogin(username, password);
            if (res.code !== 1001) {
                setErrMsg(res.message);
                return;
            }
            const token = res.token;
            store.dispatch(setToken(token));
            store.dispatch(setUsername(username));
            localStorage.setItem("token", token);
            localStorage.setItem("username", username);
            setErrMsg("登录成功，正在跳转...");
            navigate("/");
        } catch (error) {
            setErrMsg((error as Error).message);
        }
    };
};

/**
 * 登录页面：支持账号密码登录并在成功后跳转至主页。
 */
const LoginPage = () => {
    const [inputUsername, setInputUsername] = useState("");
    const [password, setPassword] = useState("");
    const [errMsg, setErrMsg] = useState("");
    const navigate = useNavigate();
    const login = useLogin(setErrMsg);

    return (
        <div className={style.main_page}>
            <NaviBar/>
            <div className={style.main_content}>
                <div className={style.card}>
                    <div className={style.card_title}>
                        登录账号
                    </div>
                    <input
                        type="text"
                        placeholder="请输入用户名"
                        className={style.input}
                        value={inputUsername}
                        onChange={event => {
                            setInputUsername(event.target.value);
                        }}
                    />
                    <input
                        type="password"
                        placeholder="请输入密码"
                        className={style.input}
                        value={password}
                        onChange={event => {
                            setPassword(event.target.value);
                        }}
                    />
                    <div className={style.err_msg}>
                        {errMsg}
                    </div>
                    <button
                        className={style.submit_button}
                        onClick={() => {
                            login(inputUsername, password);
                        }}
                    >
                        登录
                    </button>
                    <div className={style.go_login}>
                        还没有账号？
                        <span
                            className={style.go_register_link}
                            onClick={() => {
                                navigate("/register");
                            }}
                        >
                            前往注册
                        </span>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default LoginPage;
