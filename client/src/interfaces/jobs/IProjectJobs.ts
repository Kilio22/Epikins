import { IWorkgroupsData } from '../IWorkgroupsData';

export interface IProjectJobsMatchParams {
    project: string
}

export interface IProjectJobsState {
    isBuilding: boolean
    isLoading: boolean,
    workgroupsData: IWorkgroupsData[],
    selectedJobs: string[]
}

export const projectJobsInitialState: IProjectJobsState = {
    isBuilding: false,
    isLoading: false,
    workgroupsData: [],
    selectedJobs: []
};
