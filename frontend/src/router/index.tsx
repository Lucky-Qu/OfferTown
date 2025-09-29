import {createBrowserRouter} from "react-router";
import MainPage from "../pages/MainPage";
import NotFoundPage from "../pages/NotFoundPage";

const router = createBrowserRouter(
    [
        {
            path: "/",
            element: <MainPage />
        },

        {
            path: "*",
            element: <NotFoundPage />
        }
    ]
)

export default router