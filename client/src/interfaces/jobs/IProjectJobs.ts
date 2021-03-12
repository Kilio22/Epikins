import { IWorkgroupsData } from '../IWorkgroupsData';
import { IProject } from '../projects/IProject';

export interface IProjectJobsLocationState {
    project: IProject | null
}

export interface IProjectJobsMatchParams {
    module: string,
    project: string
}

export interface IProjectJobsState {
    project: IProject | null,
    isBuilding: boolean,
    isLoading: boolean,
    workgroupsData: IWorkgroupsData[],
    selectedCity: string,
    selectedJobs: string[]
}

export const projectJobsInitialState: IProjectJobsState = {
    project: null,
    isBuilding: false,
    isLoading: false,
    workgroupsData: [],
    selectedCity: '',
    selectedJobs: []
};
