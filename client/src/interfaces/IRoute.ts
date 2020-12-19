import {ComponentClass, FunctionComponent} from "react";
import Projects from "../components/projects/Projects";
import Home from "../components/Home";
import {RouteComponentProps} from "react-router-dom";
import ProjectJobs from "../components/projectJobs/ProjectJobs";

interface IRoute {
    path: string,
    component: ComponentClass<any> | FunctionComponent<any>
}

export interface IRouteProps<PARAMS> {
    routeProps: RouteComponentProps<PARAMS>
}

export const routePrefix: string = "/";

export const routes: IRoute[] = [
    {
        path: routePrefix,
        component: Home
    },
    {
        path: routePrefix + "projects",
        component: Projects
    },
    {
        path: routePrefix + "projects/:project",
        component: ProjectJobs
    }
]