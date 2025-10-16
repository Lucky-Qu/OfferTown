import {useState} from "react";
import {useNavigate} from "react-router";
import NaviBar from "../../components/NaviBar";
import style from "./RegisterPage.module.css";
import {userRegister} from "../../api/user.ts";

/**
 * 注册逻辑：封装校验与后端交互。
 * @param setErrMsg 用于更新错误描述的 setter
 */
const useRegister = (setErrMsg: (message: string) => void) => {
    const navigate = useNavigate();
    return async (username: string, password: string, confirmPassword: string) => {
        if (confirmPassword !== password) {
            setErrMsg("两次输入的密码不一致");
            return;
        }
        setErrMsg("");
        try {
            const res = await userRegister(username, password);
            if (res.code !== 1001) {
                setErrMsg(res.message);
                return;
            }
            setErrMsg("注册成功，正在跳转登录...");
            navigate("/login");
        } catch (error) {
            setErrMsg((error as Error).message);
        }
    };
};

/**
 * 注册页面：提供账号注册与快速跳转至登录页的入口。
 */
const RegisterPage = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [checkPassword, setCheckPassword] = useState("");
    const [errMsg, setErrMsg] = useState("");
    const navigate = useNavigate();
    const register = useRegister(setErrMsg);

    return (
        <div className={style.main_page}>
            <NaviBar/>
            <div className={style.main_content}>
                <div className={style.card}>
                    <div className={style.card_title}>
                        注册账号
                    </div>
                    <input
                        type="text"
                        placeholder="用户名:长度为2～20位,只能包含中英文,数字,-,_,."
                        className={style.input}
                        value={username}
                        onChange={event => {
                            setUsername(event.target.value);
                        }}
                    />
                    <input
                        type="password"
                        placeholder="密码:长度为8～20位，可以包含A-Z,a-z,!-/"
                        className={style.input}
                        value={password}
                        onChange={event => {
                            setPassword(event.target.value);
                        }}
                    />
                    <input
                        type="password"
                        placeholder="请二次输入密码"
                        className={style.input}
                        value={checkPassword}
                        onChange={event => {
                            setCheckPassword(event.target.value);
                        }}
                    />
                    <div className={style.err_msg}>
                        {errMsg}
                    </div>
                    <button
                        className={style.submit_button}
                        onClick={() => {
                            register(username, password, checkPassword);
                        }}
                    >
                        注册
                    </button>
                    <div className={style.go_register}>
                        已有账号？
                        <span
                            className={style.go_login_link}
                            onClick={() => {
                                navigate("/login");
                            }}
                        >
                            前往登录
                        </span>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default RegisterPage;
