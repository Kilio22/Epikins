import { IUser } from '../IUser';
import { IApiUser } from './IApiUser';
import { ChangeUsersStateByProperty, HandleType } from '../Functions';

export interface IUsersTableProps {
    connectedUser: IUser,
    users: IApiUser[],
    jenkinsCredentials: string[],
    isEditing: boolean,
    changeUsersStateByProperty: ChangeUsersStateByProperty,
    onFirstDeleteClick: HandleType<string>
}
