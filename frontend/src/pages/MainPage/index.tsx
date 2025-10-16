import NaviBar from "../../components/NaviBar";
import style from "./MainPage.module.css";

/**
 * 首页入口，后续可拓展推荐内容等信息。
 */
const MainPage = () => {
    return (
        <div className={style.page}>
            <NaviBar/>
            主页面
        </div>
    );
};

export default MainPage;
