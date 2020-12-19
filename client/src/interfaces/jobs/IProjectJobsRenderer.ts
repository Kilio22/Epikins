import {IProjectJobsMatchParams} from "./IProjectJobs";
import {RouteComponentProps} from "react-router-dom";
import {IGroupData} from "../IGroupData";
import {OnBuildClick, OnCheckboxChange} from "../Functions";

export interface IProjectJobsRendererProps {
    groupsData: IGroupData[],
    isBuilding: boolean,
    selectedJobs: string[],
    onCheckboxChange: OnCheckboxChange,
    onBuildClick: OnBuildClick,
    onGlobalBuildClick: OnBuildClick,
    routeProps: RouteComponentProps<IProjectJobsMatchParams>
}

export interface IProjectJobsRendererState {
    queryString: string
}

export const jobsRendererInitialState: IProjectJobsRendererState = {
    queryString: ""
}