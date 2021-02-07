import { IProject } from '../projects/IProject';

export interface IProjectsManagementState {
    allSelected: boolean,
    isLoading: boolean,
    projects: IProject[],
    clickedProject: IProject | null
}

export const projectsManagementInitialState: IProjectsManagementState = {
    allSelected: false,
    isLoading: false,
    projects: [],
    clickedProject: null
};

export type ChangeProjectsManagementStateByProperty = (property: keyof IProjectsManagementState, value: any) => void
