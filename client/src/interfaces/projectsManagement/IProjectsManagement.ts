import { IProject } from '../projects/IProject';

export interface IProjectsManagementState {
    isLoading: boolean,
    projects: IProject[],
    selectedProject: IProject | null
}

export const projectsManagementInitialState: IProjectsManagementState = {
    isLoading: false,
    projects: [],
    selectedProject: null
};

export type ChangeProjectsManagementStateByProperty = (property: keyof IProjectsManagementState, value: any) => void
