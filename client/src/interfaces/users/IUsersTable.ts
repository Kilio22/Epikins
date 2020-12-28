import { IUser } from '../IUser';
import { IApiUser } from './IApiUser';
import { ChangeUsersStateByProperty, OnFirstDeleteClick } from '../Functions';

export interface IUsersTableProps {
    connectedUser: IUser,
    users: IApiUser[],
    jenkinsCredentials: string[],
    isEditing: boolean,
    changeUsersStateByProperty: ChangeUsersStateByProperty,
    onFirstDeleteClick: OnFirstDeleteClick
}
