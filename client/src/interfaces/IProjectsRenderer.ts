import { RouteComponentProps } from 'react-router-dom';
import { IProject } from './projects/IProject';
import { HandleType, OnCheckboxChange, OnProjectClick } from './Functions';
import { FunctionComponent } from 'react';
import { IProjectRendererProps } from './projects/IProjectRenderer';

export interface IProjectsRendererProps {
    allSelected: boolean,
    changeAllSelected: HandleType<boolean> | null,
    onSelectAllClick: OnCheckboxChange<IProject[]> | null,
    onCheckboxClick: OnProjectClick | null,
    onProjectClick: OnProjectClick,
    projects: IProject[],
    ProjectRenderer: FunctionComponent<IProjectRendererProps>,
    routeProps: RouteComponentProps<any>,
    showSwitch: boolean
}

export interface IProjectsRendererState {
    queryString: string,
    selectSearch: string,
    selectedModules: string[]
}

export const projectsRendererInitialState: IProjectsRendererState = {
    queryString: '',
    selectSearch: '',
    selectedModules: []
};
