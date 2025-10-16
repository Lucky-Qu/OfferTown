import style from "./NaviBar.module.css";
import searchIcon from "../../assets/search_image.png";
import {useNavigate} from "react-router";
import {useState} from "react";
import store from "../../store";
import {clearToken, clearUsername} from "../../store/modules/user";

const NaviBar = () => {
    const navigate = useNavigate()
    const [isSearchOpen, setIsSearchOpen] = useState(false)
    const [inputContent, setInputContent] = useState("");
    const username = store.getState().user.username
    const [isLogin, setIsLogin] = useState(username !== null);
    return (
        <>
                <div className={style.page}>
                    <div className={style.header}>
                        <div className={style.header_left_bar}>
                            <div className={style.logo_name} onClick={() => navigate("/")}>
                                OfferTown
                            </div>
                            <div className={style.feature_button} onClick={() => navigate("/learn")}>
                                学习
                            </div>
                            <div className={style.feature_button} onClick={() => navigate("/question")}>
                                题库
                            </div>
                        </div>
                        <div className={style.header_right_bar}>
                            <div className={style.search}>
                                <img alt={"搜索图标"} src={searchIcon} className={style.search_icon} onClick={() => {setIsSearchOpen(!isSearchOpen)}}></img>
                                <input className={isSearchOpen ? style.search_input_active : style.search_input_inactive}
                                       value={inputContent}
                                       placeholder={"搜索题目..."}
                                       onChange={e => {
                                    setInputContent(e.target.value)
                                    //TODO 搜索功能
                                }}/>
                            </div>
                            <div className={isLogin?style.inactive:style.feature_button} onClick={() => navigate("/register")}>
                                注册
                            </div>
                            <div className={isLogin?style.inactive:style.feature_button} onClick={() => navigate("/login")}>
                                登录
                            </div>
                            <div className={!isLogin?style.inactive:style.username}>
                                {username}
                                <span className={style.feature_button}
                                onClick={() => {
                                    store.dispatch(clearToken())
                                    store.dispatch(clearUsername())
                                    localStorage.removeItem("token")
                                    localStorage.removeItem("username")
                                    setIsLogin(!isLogin)
                                }}>
                                登出
                                </span>
                            </div>
                        </div>
                    </div>
                </div>
        </>
    )
}

export default NaviBar