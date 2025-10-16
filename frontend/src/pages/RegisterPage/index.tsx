import NaviBar from "../../components/NaviBar";
import style from "./RegisterPage.module.css"
import {useNavigate} from "react-router";
import {useState} from "react";
import {userRegister} from "../../api/user.ts";

const RegisterPage = () => {
    const navi = useNavigate()
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const [checkPassword, setCheckPassword] = useState("")
    const [errMsg, setErrMsg] = useState("")
    const register = (username: string, password: string, checkPassword: string) => {
        if (checkPassword !== password) {
            setErrMsg("两次输入的密码不一致")
            return
        }
        setErrMsg("")
        userRegister(username, password).
        then((res) => {
            if (res.code !== 1001){
                setErrMsg(res.message)
                return
            }
            setErrMsg("注册成功，正在跳转登录...")
            navi("/login")
        }).
        catch(err => {
            setErrMsg(err.message)
        })
    }
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
                        onChange={(e) => {setUsername(e.target.value)}}
                    />
                    <input
                        type="password"
                        placeholder="密码:长度为8～20位，可以包含A-Z,a-z,!-/"
                        className={style.input}
                        value={password}
                        onChange={(e) => {setPassword(e.target.value)}}
                    />
                    <input
                        type="password"
                        placeholder="请二次输入密码"
                        className={style.input}
                        value={checkPassword}
                        onChange={(e) => {setCheckPassword(e.target.value)}}
                    />
                    <div className={style.err_msg}>
                        {errMsg}
                    </div>
                    <button className={style.submit_button} onClick={() => {register(username, password, checkPassword)}}>
                        注册
                    </button>
                    <div className={style.go_register}>
                        已有账号？<span className={style.go_login_link} onClick={() => {navi("/login")}}>前往登录</span>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default RegisterPage