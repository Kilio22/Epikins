import { IGroupData } from './IGroupData';
import { IAppState } from './IApp';
import { BaseSyntheticEvent } from 'react';
import { IUsersState } from './IUsers';

export type OnSignOutClick = () => void;
export type OnCheckboxChange = (checked: boolean, job: IGroupData) => void;
export type OnBuildClick = (visibility: string) => void;
export type ChangeAppStateByProperty = (propertyName: keyof IAppState, value: any, shouldCallback: boolean) => void;
export type OnJobClick = (event: BaseSyntheticEvent, url: string) => void;
export type OnButtonClick = () => void;
export type ChangeUsersStateByProperty = (propertyName: keyof IUsersState, value: any) => void;
