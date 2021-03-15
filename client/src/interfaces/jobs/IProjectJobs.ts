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
    isBuilding: boolean,
    isLoading: boolean,
    selectedCity: string,
    selectedJobs: string[]
    project: IProject | null,
    workgroupsData: IWorkgroupsData[],
}

export const projectJobsInitialState: IProjectJobsState = {
    isBuilding: false,
    isLoading: false,
    selectedCity: '',
    selectedJobs: [],
    project: null,
    workgroupsData: []
};
