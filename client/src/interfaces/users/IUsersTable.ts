import { IUser } from '../IUser';
import { IApiUser } from '../IApiUser';
import { ChangeUsersStateByProperty, OnButtonClick } from '../Functions';


export interface IUsersTableProps {
    user: IUser,
    modifiedUsers: IApiUser[],
    jenkinsCredentials: string[],
    isEditing: boolean,
    changeUsersStateByProperty: ChangeUsersStateByProperty,
    getUsers: OnButtonClick
}
