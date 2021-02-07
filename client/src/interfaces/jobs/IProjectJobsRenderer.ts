import { IProjectJobsMatchParams } from './IProjectJobs';
import { RouteComponentProps } from 'react-router-dom';
import { IWorkgroupsData } from '../IWorkgroupsData';
import { OnBuildClick, OnCheckboxChange } from '../Functions';

export interface IProjectJobsRendererProps {
    availableCities: string[] | undefined,
    workgroupsData: IWorkgroupsData[],
    isBuilding: boolean,
    selectedJobs: string[],
    selectedCity: string,
    onCheckboxChange: OnCheckboxChange<IWorkgroupsData>,
    onBuildClick: OnBuildClick,
    onGlobalBuildClick: OnBuildClick,
    routeProps: RouteComponentProps<IProjectJobsMatchParams>,
    onCitySelected: OnBuildClick
}

export interface IProjectJobsRendererState {
    queryString: string
}

export const jobsRendererInitialState: IProjectJobsRendererState = {
    queryString: ''
};
