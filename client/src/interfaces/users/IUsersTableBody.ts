import { IApiUser } from './IApiUser';
import { IUser } from '../IUser';
import { OnFirstDeleteClick } from '../Functions';

type OnCheckboxClick = (modifiedUsers: IApiUser[], modifiedUser: IApiUser, modifiedUserIdx: number, currentRole: string) => void;

export interface IUsersTableBodyProps {
    connectedUser: IUser,
    users: IApiUser[],
    jenkinsCredentials: string[],
    isEditing: boolean,
    onCheckboxClick: OnCheckboxClick,
    onFirstDeleteClick: OnFirstDeleteClick
}
