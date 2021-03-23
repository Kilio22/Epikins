import { RouteComponentProps } from 'react-router-dom';
import { IWorkgroupsData } from '../IWorkgroupsData';
import { OnBuildClick, OnButtonClick, OnCheckboxChange } from '../Functions';

export interface IProjectJobsRendererProps {
    availableCities: string[] | undefined,
    workgroupsData: IWorkgroupsData[],
    isBuilding: boolean,
    selectedJobs: string[],
    selectedCity: string,
    onCheckboxChange: OnCheckboxChange<IWorkgroupsData>,
    onBuildClick: OnBuildClick,
    onForceUpdateClick: OnButtonClick,
    onGlobalBuildClick: OnBuildClick,
    routeProps: RouteComponentProps,
    onCitySelected: OnBuildClick
}

export interface IProjectJobsRendererState {
    queryString: string
}

export const jobsRendererInitialState: IProjectJobsRendererState = {
    queryString: ''
};
