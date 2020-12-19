import {IJob} from "../IJob";
import {RouteComponentProps} from "react-router-dom";

export interface IProjectsRendererProps {
    projects: IJob[],
    routeProps: RouteComponentProps<any>
}

export interface IProjectsRendererState {
    queryString: string
}

export const projectsRendererInitialState: IProjectsRendererState = {
    queryString: ""
}