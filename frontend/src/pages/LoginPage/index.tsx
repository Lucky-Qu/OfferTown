import NaviBar from "../../components/NaviBar";
import style from "./LoginPage.module.css";
import {useNavigate} from "react-router";
import {useState} from "react";
import {userLogin} from "../../api/user.ts";
import store from "../../store";
import {setToken, setUsername} from "../../store/modules/user";

const LoginPage = () => {
    const navi = useNavigate()
    const [inputUsername, setInputUsername] = useState("")
    const [password, setPassword] = useState("")
    const [errMsg, setErrMsg] = useState("")
    const login = (username: string, password: string) => {
        userLogin(username, password).
        then(res=>{
            if (res.code !== 1001){
                setErrMsg(res.message)
                return
            }
            const token = res.token
            store.dispatch(setToken(token))
            store.dispatch(setUsername(username))
            localStorage.setItem("token", token)
            localStorage.setItem("username", username)
            setErrMsg("登录成功，正在跳转...")
            navi("/")
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
                        登录账号
                    </div>
                    <input
                        type="text"
                        placeholder="请输入用户名"
                        className={style.input}
                        value={inputUsername}
                        onChange={(e) => {setInputUsername(e.target.value)}}
                    />
                    <input
                        type="password"
                        placeholder="请输入密码"
                        className={style.input}
                        value={password}
                        onChange={(e) => {setPassword(e.target.value)}}
                    />
                    <div className={style.err_msg}>
                        {errMsg}
                    </div>
                    <button className={style.submit_button} onClick={() => {login(inputUsername, password)}}>
                        登录
                    </button>
                    <div className={style.go_login}>
                        还没有账号？<span className={style.go_register_link} onClick={() => {navi("/register")}}>前往注册</span>
                    </div>
                </div>
            </div>
        </div>
    )
}
export default LoginPage