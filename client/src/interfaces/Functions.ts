import { IAppState } from './IApp';
import { BaseSyntheticEvent } from 'react';
import { IUsersState } from './users/IUsers';
import { IProject } from './projects/IProject';
import { IStudentJob } from './myProjects/IStudentJob';
import { IMyProjectsState } from '../components/myProjects/MyProjects';
import { InputActionMeta } from 'react-select';
import { ISelectOption } from './projects/ISelectOption';

export type ChangeAppStateByProperty = (propertyName: keyof IAppState, value: any, shouldCallback: boolean) => void;
export type ChangeMyProjectsStateByProperty = (key: keyof IMyProjectsState, value: any) => void;
export type ChangeUsersStateByProperty = (propertyName: keyof IUsersState, value: any) => void;

export type HandleType<T> = (value: T) => void;
export type OnBuildClick = (visibility: string) => void;
export type OnButtonClick = () => void;
export type OnCheckboxChange<T> = (checked: boolean, value: T) => void;
export type OnJobClick = (event: BaseSyntheticEvent, url: string) => void;
export type OnProjectClick = (project: IProject) => void;
export type OnSelectChange = (modules: ISelectOption[]) => void;
export type OnSelectSearchChange = (value: string, actionMeta: InputActionMeta) => void;
export type OnSignOutClick = () => void;
export type OnStudentProjectClick = (project: IStudentJob) => void;
