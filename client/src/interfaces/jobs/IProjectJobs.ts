import {IGroupData} from "../IGroupData";

export interface IProjectJobsMatchParams {
    project: string
}

export interface IProjectJobsState {
    isBuilding: boolean
    isLoading: boolean,
    groupsData: IGroupData[],
    selectedJobs: string[]
}

export const projectJobsInitialState: IProjectJobsState = {
    isBuilding: false,
    isLoading: false,
    groupsData: [],
    selectedJobs: []
}