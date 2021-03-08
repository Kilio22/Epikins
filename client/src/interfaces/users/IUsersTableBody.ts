import { IApiUser } from './IApiUser';
import { HandleType } from '../Functions';

type OnCheckboxClick = (modifiedUsers: IApiUser[], modifiedUser: IApiUser, modifiedUserIdx: number, currentRole: string) => void;

export interface IUsersTableBodyProps {
    users: IApiUser[],
    jenkinsCredentials: string[],
    isEditing: boolean,
    onCheckboxClick: OnCheckboxClick,
    onFirstDeleteClick: HandleType<string>
}
