import {IJob} from "../IJob";

export interface IProjectsState {
    projects: IJob[],
    isLoading: boolean
}

export const projectsInitialState: IProjectsState = {
    projects: [],
    isLoading: false
}
