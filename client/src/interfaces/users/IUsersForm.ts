import { ChangeUsersStateByProperty, OnButtonClick } from '../Functions';
import { apiUserInitialState, IApiUser } from './IApiUser';

export interface IUsersFormProps {
    changeUsersStateByProperty: ChangeUsersStateByProperty,
    getUsers: OnButtonClick,
    jenkinsCredentials: string[]
}

export interface IUsersFormState {
    user: IApiUser
    validation: boolean,
    isLoading: boolean
}

export const usersFormInitialState: IUsersFormState = {
    user: apiUserInitialState,
    validation: false,
    isLoading: false
};
