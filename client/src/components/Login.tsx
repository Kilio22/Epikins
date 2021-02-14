import React, { Component } from 'react';
import loginWith from '../assets/login_with_microsoft.png';
import { Redirect } from 'react-router-dom';
import { ILoginProps, ILoginState, loginInitialState } from '../interfaces/ILogin';
import { IUser, userInitialState } from '../interfaces/IUser';
import EpikinsApiService from '../services/EpikinsApiService';
import { routePrefix } from '../interfaces/IRoute';
import { appInitialContext } from '../interfaces/IAppContext';
import { authServiceObj } from '../services/AuthService';
import { AccountInfo, AuthenticationResult } from '@azure/msal-browser';
import Loading from './Loading';
import { IApiUser } from '../interfaces/users/IApiUser';

class Login extends Component<ILoginProps, ILoginState> {
    static contextType = appInitialContext;
    context!: React.ContextType<typeof appInitialContext>;

    constructor(props: any) {
        super(props);

        this.state = {
            ...loginInitialState,
            isConnecting: true
        };

        this.getUserProfile = this.getUserProfile.bind(this);
        this.login = this.login.bind(this);
    }

    async componentDidMount() {
        let res: AuthenticationResult | null = null;

        try {
            res = await authServiceObj.handleRedirectPromise();
        } catch (e) {
            console.log(e);
            this.setState(loginInitialState);
            return;
        }
        const account: AccountInfo | null = authServiceObj.handleResponse(res);
        if (account) {
            await this.getUserProfile();
        } else {
            this.setState(loginInitialState);
        }
    }

    render() {
        return (
            <appInitialContext.Consumer>
                {context => (
                    context.user.isLoggedIn ?
                        this.props.routeProps.location.state?.from ?
                            <Redirect to={{pathname: this.props.routeProps.location.state.from.pathname}}/>
                            :
                            <Redirect to={{pathname: routePrefix}}/>
                        :
                        this.state.isConnecting ?
                            <Loading/>
                            :
                            <div className={'h-100 d-flex justify-content-center align-items-center'}>
                                <input type={'image'} src={loginWith} alt={'login button'}
                                       onClick={this.login}/>
                            </div>
                )}
            </appInitialContext.Consumer>
        );
    }

    async login() {
        this.setState({
            ...this.state,
            isConnecting: true
        });
        try {
            await authServiceObj.login();
        } catch (e) {
            console.log(e);
            this.setState(loginInitialState);
        }
    }

    async getUserProfile() {
        let newUser: IUser = userInitialState;
        let accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            if (this.context.changeAppStateByProperty != null) {
                this.context.changeAppStateByProperty('user', userInitialState, false);
            }
            this.setState(loginInitialState);
            return;
        }

        try {
            const decodedJWT: any = await authServiceObj.verifyAccessToken(accessToken);
            if (decodedJWT === null) {
                this.setState(loginInitialState);
                return;
            }

            newUser = {
                ...newUser,
                name: decodedJWT.name,
                email: decodedJWT.email,
                isLoggedIn: true
            };
        } catch (err) {
            console.log(err);
            this.setState(loginInitialState);
            return;
        }

        const res: IApiUser | null = await EpikinsApiService.login(accessToken);
        if (res) {
            newUser = {
                ...newUser,
                roles: res.roles
            };
        }

        if (this.context.changeAppStateByProperty) {
            this.context.changeAppStateByProperty('user', newUser, false);
        }
    }
}

export default Login;
