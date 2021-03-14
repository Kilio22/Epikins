import { IApiUser } from './IApiUser';

export interface IUsersState {
    initialUsers: IApiUser[],
    modifiedUsers: IApiUser[],
    jenkinsCredentials: string[],
    isEditing: boolean,
    isSaving: boolean,
    isAdding: boolean,
    showDeletePopup: boolean,
    toDelete: string,
    isLoading: boolean
}

export const usersInitialState: IUsersState = {
    initialUsers: [],
    modifiedUsers: [],
    jenkinsCredentials: [],
    isEditing: false,
    isSaving: false,
    isAdding: false,
    showDeletePopup: false,
    toDelete: '',
    isLoading: false
};
