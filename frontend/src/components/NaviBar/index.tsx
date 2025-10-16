import {useState} from "react";
import {useNavigate} from "react-router";
import {useDispatch, useSelector} from "react-redux";
import style from "./NaviBar.module.css";
import searchIcon from "../../assets/search_image.png";
import {clearToken, clearUsername} from "../../store/modules/user";
import type {AppDispatch, RootState} from "../../store";

/**
 * 顶部导航栏，负责展示基础导航入口、用户登录态与搜索入口。
 */
const NaviBar = () => {
    const navigate = useNavigate();
    const dispatch = useDispatch<AppDispatch>();
    const username = useSelector((state: RootState) => state.user.username ?? "");
    const isLogin = Boolean(username);
    const [isSearchOpen, setIsSearchOpen] = useState(false);
    const [inputContent, setInputContent] = useState("");

    return (
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
                        <img
                            alt={"搜索图标"}
                            src={searchIcon}
                            className={style.search_icon}
                            onClick={() => {
                                setIsSearchOpen(!isSearchOpen);
                            }}
                        />
                        <input
                            className={isSearchOpen ? style.search_input_active : style.search_input_inactive}
                            value={inputContent}
                            placeholder={"搜索题目..."}
                            onChange={e => {
                                setInputContent(e.target.value);
                                // TODO: 搜索功能在页面内处理
                            }}
                        />
                    </div>
                    <div className={isLogin ? style.inactive : style.feature_button} onClick={() => navigate("/register")}>
                        注册
                    </div>
                    <div className={isLogin ? style.inactive : style.feature_button} onClick={() => navigate("/login")}>
                        登录
                    </div>
                    <div className={!isLogin ? style.inactive : style.username}>
                        {username}
                        <span
                            className={style.feature_button}
                            onClick={() => {
                                dispatch(clearToken());
                                dispatch(clearUsername());
                                localStorage.removeItem("token");
                                localStorage.removeItem("username");
                            }}
                        >
                            登出
                        </span>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default NaviBar;
