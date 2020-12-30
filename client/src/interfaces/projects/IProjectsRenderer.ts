import { RouteComponentProps } from 'react-router-dom';
import { IProject } from './IProject';

export interface IProjectsRendererProps {
    projects: IProject[],
    routeProps: RouteComponentProps<any>
}

export interface IProjectsRendererState {
    queryString: string
}

export const projectsRendererInitialState: IProjectsRendererState = {
    queryString: ''
};
