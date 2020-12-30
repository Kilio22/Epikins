import { IProject } from './IProject';

export interface IProjectsState {
    projects: IProject[],
    isLoading: boolean
}

export const projectsInitialState: IProjectsState = {
    projects: [],
    isLoading: false
};
