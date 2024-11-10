import { useRoutes, RouteObject } from "react-router-dom";
import MainRoutes from "./MainRoutes";

function ConfigRoutes() {

    // const isLoggedIn = localStorage.getItem("isLogin") === "true";
    // const loggedinByOwner = localStorage.getItem("loginByOwner") == "true";

    let routes: RouteObject[] = [];

    // if (isLoggedIn && loggedinByOwner) {
    //     routes = [OwnerRoutes()];
    // } 
    // else if (isLoggedIn){
    //     routes = [CustomerRoutes()]
    // }
    // else {
    //     routes = [MainRoutes()];
    // }
    routes = [MainRoutes()];

    return useRoutes(routes);
}

export default ConfigRoutes;