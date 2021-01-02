import { RouteComponentProps } from 'react-router-dom';
import { IProject } from './IProject';
import { OnProjectClick } from '../Functions';
import { FunctionComponent } from 'react';
import { IProjectRendererProps } from './IProjectRenderer';

export interface IProjectsRendererProps {
    projects: IProject[],
    routeProps: RouteComponentProps<any>,
    onProjectClick: OnProjectClick,
    ProjectRenderer: FunctionComponent<IProjectRendererProps>,
    showSwitch: boolean
}

export interface IProjectsRendererState {
    queryString: string
}

export const projectsRendererInitialState: IProjectsRendererState = {
    queryString: ''
};
