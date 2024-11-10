import { RouteObject } from "react-router-dom";
import MinimalLayout from "../layouts/MinimalLayout/MinimalLayout";
import Home from "../pages/home/Home";

const MainRoutes = (): RouteObject => {
    return {
        path: "/",

        element: <MinimalLayout />,

        children: [
            {
                path: "/",
                element: <Home />,
            },
        ],
    };
};

export default MainRoutes;