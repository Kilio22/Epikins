import { IProjectJobsMatchParams } from './IProjectJobs';
import { RouteComponentProps } from 'react-router-dom';
import { IWorkgroupsData } from '../IWorkgroupsData';
import { OnBuildClick, OnCheckboxChange } from '../Functions';

export interface IProjectJobsRendererProps {
    workgroupsData: IWorkgroupsData[],
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
    queryString: ''
};
