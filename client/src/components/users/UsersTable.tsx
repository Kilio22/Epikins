import * as React from 'react';
import { Component } from 'react';
import { IApiUser } from '../../interfaces/IApiUser';
import { Table } from 'react-bootstrap';
import UsersTableBody from './UsersTableBody';
import UsersTableHeader from './UsersTableHeader';
import { IUsersTableProps } from '../../interfaces/users/IUsersTable';
import { authServiceObj } from '../../services/AuthService';
import EpikinsApiService from '../../services/EpikinsApiService';
import { userInitialState } from '../../interfaces/IUser';

class UsersTable extends Component<IUsersTableProps> {
    constructor(props: IUsersTableProps) {
        super(props);

        this.onCheckboxClick = this.onCheckboxClick.bind(this);
        this.onDeleteClick = this.onDeleteClick.bind(this);
        this.resetUser = this.resetUser.bind(this);
        this.setErrorMessage = this.setErrorMessage.bind(this);
    }

    render() {
        return (
            <Table bordered responsive={'md'} striped className={'mt-3'}>
                <UsersTableHeader/>
                <UsersTableBody user={this.props.user}
                                modifiedUsers={this.props.modifiedUsers}
                                jenkinsCredentials={this.props.jenkinsCredentials}
                                isEditing={this.props.isEditing}
                                onCheckboxClick={this.onCheckboxClick}
                                onDeleteClick={this.onDeleteClick}/>
            </Table>
        );
    }

    resetUser() {
        if (this.context.changeAppStateByProperty != null) {
            this.context.changeAppStateByProperty('user', userInitialState, false);
        }
    }

    setErrorMessage(message: string) {
        if (this.context.changeAppStateByProperty) {
            this.context.changeAppStateByProperty('errorMessage', message, true);
        }
    }

    onCheckboxClick(modifiedUsers: IApiUser[], modifiedUser: IApiUser, modifiedUserIdx: number, currentRole: string) {
        const userRoleIdx = modifiedUser.roles.indexOf(currentRole.toLocaleLowerCase());

        if (userRoleIdx === -1) {
            modifiedUsers[modifiedUserIdx].roles.push(currentRole.toLocaleLowerCase());
        } else {
            modifiedUsers[modifiedUserIdx].roles.splice(userRoleIdx, 1);
        }
        this.props.changeUsersStateByProperty('modifiedUsers', modifiedUsers);
    }

    async onDeleteClick(email: string) {
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            this.resetUser();
            return;
        }

        const res = await EpikinsApiService.deleteUser(email, accessToken);
        if (!res) {
            this.setErrorMessage('Cannot delete user, please try to reload the page.');
        }
        await this.props.getUsers();
    }
}

export default UsersTable;
