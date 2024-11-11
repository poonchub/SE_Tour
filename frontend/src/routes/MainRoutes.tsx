import { RouteObject } from "react-router-dom";
import MinimalLayout from "../layouts/MinimalLayout/MinimalLayout";
import Home from "../pages/home/Home";
import TourPackage from "../pages/tourPackage/TourPackage";
import Profile from "../pages/profile/Profile";

const MainRoutes = (): RouteObject => {
    return {
        path: "/",

        element: <MinimalLayout />,

        children: [
            {
                path: "/",
                element: <Home />,
            },
            {
                path: "/tour-package",
                element: <TourPackage />,
            },
            {
                path: "/profile",
                element: <Profile />,
            },
        ],
    };
};

export default MainRoutes;