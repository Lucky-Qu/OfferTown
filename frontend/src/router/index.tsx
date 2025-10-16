import {createBrowserRouter} from "react-router";
import MainPage from "../pages/MainPage";
import NotFoundPage from "../pages/NotFoundPage";
import LoginPage from "../pages/LoginPage";
import RegisterPage from "../pages/RegisterPage";
import LearnPage from "../pages/LearnPage";
import QuestionPage from "../pages/QuestionPage";

/**
 * 应用路由配置：集中维护各页面的路径映射。
 */
const router = createBrowserRouter([
    {
        path: "/",
        element: <MainPage/>,
    },
    {
        path: "/login",
        element: <LoginPage/>,
    },
    {
        path: "/register",
        element: <RegisterPage/>,
    },
    {
        path: "/learn",
        element: <LearnPage/>,
    },
    {
        path: "/question",
        element: <QuestionPage/>,
    },
    {
        path: "*",
        element: <NotFoundPage/>,
    },
]);

export default router;
