import React, { Component } from 'react';
import JenkinsCredentialsToolbox from './JenkinsCredentialsToolbox';
import { userInitialState } from '../../interfaces/IUser';
import { authServiceObj } from '../../services/AuthService';
import EpikinsApiService from '../../services/EpikinsApiService';
import { IRouteProps } from '../../interfaces/IRoute';
import { appInitialContext } from '../../interfaces/IAppContext';
import JenkinsCredentialsTable from './JenkinsCredentialsTable';
import Loading from '../Loading';
import JenkinsCredentialsForm from './JenkinsCredentialsForm';
import JenkinsCredentialsDeletePopup from './JenkinsCredentialsDeletePopup';
import {
    IJenkinsCredentialsState,
    jenkinsCredentialsInitialState
} from '../../interfaces/jenkinsCredentials/IJenkinsCredentials';

class JenkinsCredentials extends Component<IRouteProps<{}>, IJenkinsCredentialsState> {
    static contextType = appInitialContext;
    context!: React.ContextType<typeof appInitialContext>;

    constructor(props: IRouteProps<{}>) {
        super(props);

        this.onAddClick = this.onAddClick.bind(this);
        this.setErrorMessage = this.setErrorMessage.bind(this);
        this.resetUser = this.resetUser.bind(this);
        this.getJenkinsCredentials = this.getJenkinsCredentials.bind(this);
        this.changeJenkinsCredentialsStateByProperty = this.changeJenkinsCredentialsStateByProperty.bind(this);
        this.onDeleteClick = this.onDeleteClick.bind(this);
        this.onFirstDeleteClick = this.onFirstDeleteClick.bind(this);

        this.state = jenkinsCredentialsInitialState;
    }

    async componentDidMount() {
        this.changeJenkinsCredentialsStateByProperty('isLoading', true);
        await this.getJenkinsCredentials();
        this.changeJenkinsCredentialsStateByProperty('isLoading', false);
    }

    render() {
        return (
            this.state.isLoading ?
                <Loading/>
                :
                <div>
                    {
                        this.state.isAdding &&
                        <JenkinsCredentialsForm
                            changeJenkinsCredentialsStateByProperty={this.changeJenkinsCredentialsStateByProperty}
                            jenkinsCredentials={this.state.jenkinsCredentials}
                            getJenkinsCredentials={this.getJenkinsCredentials}/>
                    }
                    {
                        this.state.isDeleting &&
                        <JenkinsCredentialsDeletePopup onDeleteClick={this.onDeleteClick}
                                                       changeJenkinsCredentialStateByProperty={this.changeJenkinsCredentialsStateByProperty}/>
                    }
                    <JenkinsCredentialsToolbox onAddClick={this.onAddClick}/>
                    <JenkinsCredentialsTable jenkinsCredentials={this.state.jenkinsCredentials}
                                             onFirstDeleteClick={this.onFirstDeleteClick}/>
                </div>
        );
    }

    onAddClick() {
        this.setState({
            ...this.state,
            isAdding: true
        });
    }

    async onDeleteClick() {
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            this.resetUser();
            return;
        }

        const res = await EpikinsApiService.deleteJenkinsCredentials(this.state.toDelete, accessToken);
        if (!res) {
            this.setErrorMessage('Cannot delete credentials, please try to reload the page.');
        }
        this.setState({
            ...this.state,
            isDeleting: false,
            toDelete: ''
        });
        await this.getJenkinsCredentials();
    }

    onFirstDeleteClick(toDelete: string) {
        this.setState({
            ...this.state,
            isDeleting: true,
            toDelete: toDelete
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

    async getJenkinsCredentials() {
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            this.resetUser();
            return;
        }

        const data: string[] | null = await EpikinsApiService.getJenkinsCredentials(accessToken);
        if (data) {
            this.changeJenkinsCredentialsStateByProperty('jenkinsCredentials', data);
        } else {
            this.setErrorMessage('Cannot fetch data, please try to reload the page.');
        }
    }

    changeJenkinsCredentialsStateByProperty(propertyName: keyof IJenkinsCredentialsState, value: any) {
        this.setState({
            ...this.state,
            [propertyName]: value
        });
    }
}

export default JenkinsCredentials;
