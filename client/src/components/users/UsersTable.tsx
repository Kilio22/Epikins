import * as React from 'react';
import { Component } from 'react';
import { IApiUser } from '../../interfaces/users/IApiUser';
import { Table } from 'react-bootstrap';
import UsersTableBody from './UsersTableBody';
import UsersTableHeader from './UsersTableHeader';
import { IUsersTableProps } from '../../interfaces/users/IUsersTable';

class UsersTable extends Component<IUsersTableProps> {
    constructor(props: IUsersTableProps) {
        super(props);

        this.onCheckboxClick = this.onCheckboxClick.bind(this);
    }

    render() {
        return (
            <Table bordered responsive={'md'} striped className={'mt-3'}>
                <UsersTableHeader/>
                <UsersTableBody connectedUser={this.props.connectedUser}
                                users={this.props.users}
                                jenkinsCredentials={this.props.jenkinsCredentials}
                                isEditing={this.props.isEditing}
                                onCheckboxClick={this.onCheckboxClick}
                                onFirstDeleteClick={this.props.onFirstDeleteClick}/>
            </Table>
        );
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
}

export default UsersTable;
