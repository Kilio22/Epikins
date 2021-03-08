import React from 'react';
import { IRouteProps } from '../../interfaces/IRoute';
import Loading from '../Loading';
import EpikinsApiService from '../../services/EpikinsApiService';
import { authServiceObj } from '../../services/AuthService';
import { userInitialState } from '../../interfaces/IUser';
import { IBuildLogInfo } from '../../interfaces/log/IBuildLogInfo';

interface ILogState {
    buildLogInfo: IBuildLogInfo | null
    isLoading: boolean
}

const LogInitialState: ILogState = {
    buildLogInfo: null,
    isLoading: false
};

class Log extends React.Component<IRouteProps, ILogState> {
    constructor(props: IRouteProps) {
        super(props);

        this.state = LogInitialState;
    }

    async componentDidMount() {
        this.setState({
            ...this.state,
            isLoading: true
        });
        await this.getLog();
        this.setState({
            ...this.state,
            isLoading: false
        });
    }

    render() {
        return (
            <Loading/>
        );
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

    async getLog() {
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            this.resetUser();
            return;
        }

        const res = await EpikinsApiService.getLog(accessToken);
        if (res) {
            this.setState({
                ...this.state,
                buildLogInfo: res
            });
        } else {
            this.setErrorMessage('Cannot fetch data, please try to reload the page.');
        }
    }
}

export default Log;
