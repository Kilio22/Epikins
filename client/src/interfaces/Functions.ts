import { IWorkgroupsData } from './IWorkgroupsData';
import { IAppState } from './IApp';
import { BaseSyntheticEvent } from 'react';
import { IUsersState } from './users/IUsers';
import { IProject } from './projects/IProject';
import { IStudentJob } from './myProjects/IStudentJob';
import { IMyProjectsState } from '../components/myProjects/MyProjects';

export type OnSignOutClick = () => void;
export type OnCheckboxChange = (checked: boolean, job: IWorkgroupsData) => void;
export type OnBuildClick = (visibility: string) => void;
export type ChangeAppStateByProperty = (propertyName: keyof IAppState, value: any, shouldCallback: boolean) => void;
export type OnJobClick = (event: BaseSyntheticEvent, url: string) => void;
export type OnButtonClick = () => void;
export type ChangeUsersStateByProperty = (propertyName: keyof IUsersState, value: any) => void;
export type OnFirstDeleteClick = (toDelete: string) => void;
export type OnProjectClick = (project: IProject) => void;
export type OnStudentProjectClick = (project: IStudentJob) => void;
export type ChangeMyProjectsStateByProperty = (key: keyof IMyProjectsState, value: any) => void;
