import { IApiUser } from './IApiUser';
import { ChangeUsersStateByProperty, HandleType } from '../Functions';

export interface IUsersTableProps {
    users: IApiUser[],
    jenkinsCredentials: string[],
    isEditing: boolean,
    changeUsersStateByProperty: ChangeUsersStateByProperty,
    onFirstDeleteClick: HandleType<string>
}
