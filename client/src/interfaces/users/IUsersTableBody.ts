import { IApiUser } from '../IApiUser';
import { IUser } from '../IUser';

type OnCheckboxClick = (modifiedUsers: IApiUser[], modifiedUser: IApiUser, modifiedUserIdx: number, currentRole: string) => void;
type OnDeleteClick = (email: string) => void;

export interface IUsersTableBodyProps {
    connectedUser: IUser,
    users: IApiUser[],
    jenkinsCredentials: string[],
    isEditing: boolean,
    onCheckboxClick: OnCheckboxClick,
    onDeleteClick: OnDeleteClick
}
