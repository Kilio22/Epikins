import React, { Component } from 'react';
import { IUsersState, usersInitialState } from '../../interfaces/IUsers';
import UsersToolbox from './UsersToolbox';
import { userInitialState } from '../../interfaces/IUser';
import { authServiceObj } from '../../services/AuthService';
import EpikinsApiService from '../../services/EpikinsApiService';
import { IApiUser } from '../../interfaces/IApiUser';
import { IRouteProps } from '../../interfaces/IRoute';
import { appInitialContext } from '../../interfaces/IAppContext';
import UsersTable from './UsersTable';
import AddUserForm from './AddUserForm';
import Loading from '../Loading';

class Users extends Component<IRouteProps<{}>, IUsersState> {
    static contextType = appInitialContext;
    context!: React.ContextType<typeof appInitialContext>;

    constructor(props: IRouteProps<{}>) {
        super(props);

        this.onSaveClick = this.onSaveClick.bind(this);
        this.onCancelClick = this.onCancelClick.bind(this);
        this.onEditClick = this.onEditClick.bind(this);
        this.onAddClick = this.onAddClick.bind(this);
        this.setErrorMessage = this.setErrorMessage.bind(this);
        this.resetUser = this.resetUser.bind(this);
        this.getUsers = this.getUsers.bind(this);
        this.getJenkinsCredentials = this.getJenkinsCredentials.bind(this);
        this.updateUser = this.updateUser.bind(this);
        this.updateUsers = this.updateUsers.bind(this);
        this.changeUsersStateByProperty = this.changeUsersStateByProperty.bind(this);

        this.state = usersInitialState;
    }

    async componentDidMount() {
        this.changeUsersStateByProperty('isLoading', true);
        await this.getUsers();
        await this.getJenkinsCredentials();
        this.changeUsersStateByProperty('isLoading', false);
    }

    render() {
        return (
            <appInitialContext.Consumer>
                {context => (
                    this.state.isLoading ?
                        <Loading/>
                        :
                        <div>
                            {
                                this.state.isAdding &&
                                <AddUserForm changeUsersStateByProperty={this.changeUsersStateByProperty}
                                             isAdding={this.state.isAdding}
                                             jenkinsCredentials={this.state.jenkinsCredentials}
                                             getUsers={this.getUsers}/>
                            }
                            <UsersToolbox isEditing={this.state.isEditing}
                                          isSaving={this.state.isSaving}
                                          onSaveClick={this.onSaveClick}
                                          onCancelClick={this.onCancelClick}
                                          onEditClick={this.onEditClick}
                                          onAddClick={this.onAddClick}/>
                            <UsersTable modifiedUsers={this.state.modifiedUsers}
                                        jenkinsCredentials={this.state.jenkinsCredentials}
                                        isEditing={this.state.isEditing}
                                        changeUsersStateByProperty={this.changeUsersStateByProperty}
                                        user={context.user}
                                        getUsers={this.getUsers}/>
                        </div>
                )}
            </appInitialContext.Consumer>
        );
    }

    async onSaveClick() {
        this.setState({
            ...this.state,
            isSaving: true
        });
        await this.updateUsers();
        this.setState({
            ...this.state,
            isEditing: false,
            isSaving: false
        });
    }

    onCancelClick() {
        this.setState({
            ...this.state,
            modifiedUsers: this.state.initialUsers.map(value => {
                return {
                    ...value,
                    roles: [ ...value.roles ]
                };
            }),
            isEditing: false
        });
    }

    onEditClick() {
        this.setState({
            ...this.state,
            isEditing: true
        });
    }

    onAddClick() {
        this.setState({
            ...this.state,
            isAdding: true
        });
    }

    setErrorMessage(message: string) {
        if (this.context.changeAppStateByProperty) {
            this.context.changeAppStateByProperty('errorMessage', message, true);
        }
    }

    resetUser() {
        if (this.context.changeAppStateByProperty != null) {
            this.context.changeAppStateByProperty('user', userInitialState, false);
        }
    }

    async getUsers() {
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            this.resetUser();
            return;
        }

        const data: IApiUser[] | null = await EpikinsApiService.getUsers(accessToken);
        if (data) {
            this.changeUsersStateByProperty('initialUsers', [ ...data ]);
            this.changeUsersStateByProperty('modifiedUsers', data.map(value => {
                return {
                    ...value,
                    roles: [ ...value.roles ]
                };
            }));
        } else {
            this.setErrorMessage('Cannot fetch data, please try to reload the page.');
        }
    }

    async getJenkinsCredentials() {
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            this.resetUser();
            return;
        }

        const data: string[] | null = await EpikinsApiService.getJenkinsCredentials(accessToken);
        if (data) {
            this.changeUsersStateByProperty('jenkinsCredentials', data);
        } else {
            this.setErrorMessage('Cannot fetch data, please try to reload the page.');
        }
    }

    async updateUser(updatedUser: IApiUser): Promise<boolean> {
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            this.resetUser();
            return false;
        }

        const res: boolean = await EpikinsApiService.updateUser(updatedUser, accessToken);
        if (!res) {
            this.setErrorMessage('Cannot update user, please try to reload the page.');
        }
        return res;
    }

    async updateUsers() {
        for (let initialUser of this.state.initialUsers) {
            const toFind = this.state.modifiedUsers.find((modifiedUser) => {
                return modifiedUser.email.localeCompare(initialUser.email) === 0;
            });
            let shouldBeUpdated = false;

            if (!toFind) {
                return;
            }
            if (initialUser.roles.length !== toFind.roles.length || initialUser.jenkinsLogin.localeCompare(toFind.jenkinsLogin) !== 0) {
                const res = await this.updateUser(toFind);
                if (!res)
                    return;
            }
            for (let role of toFind.roles) {
                if (!initialUser.roles.includes(role)) {
                    shouldBeUpdated = true;
                    return;
                }
            }
            if (shouldBeUpdated) {
                const res = await this.updateUser(toFind);
                if (!res)
                    return;
            }
        }
        await this.getUsers();
    }

    changeUsersStateByProperty(propertyName: keyof IUsersState, value: any) {
        this.setState({
            ...this.state,
            [propertyName]: value
        });
    }
}

export default Users;
