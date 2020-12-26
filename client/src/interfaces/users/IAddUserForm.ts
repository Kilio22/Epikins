import { ChangeUsersStateByProperty, OnButtonClick } from '../Functions';
import { apiUserInitialState, IApiUser } from '../IApiUser';

export interface IAddUserFormProps {
    changeUsersStateByProperty: ChangeUsersStateByProperty,
    getUsers: OnButtonClick,
    isAdding: boolean,
    jenkinsCredentials: string[]
}

export interface IAddUserFormState {
    user: IApiUser
    validation: boolean,
    isLoading: boolean
}

export const usersFormInitialState: IAddUserFormState = {
    user: apiUserInitialState,
    validation: false,
    isLoading: false
};
