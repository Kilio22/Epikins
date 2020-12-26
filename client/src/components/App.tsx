import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import { routePrefix, routes } from '../interfaces/IRoute';
import Login from './Login';
import PrivateRoute from './PrivateRoute';
import NotFound from './NotFound';
import { appInitialContext } from '../interfaces/IAppContext';
import { authServiceObj } from '../services/AuthService';
import { appInitialState, IAppState } from '../interfaces/IApp';

class App extends Component<{}, IAppState> {
    constructor(props: any) {
        super(props);

        this.state = appInitialState;

        this.changeAppStateByProperty = this.changeAppStateByProperty.bind(this);
        this.onSignOutClick = this.onSignOutClick.bind(this);
    }

    render() {
        return (
            <appInitialContext.Provider value={{
                changeAppStateByProperty: this.changeAppStateByProperty,
                fuMode: this.state.fuMode,
                user: this.state.user,
                errorMessage: this.state.errorMessage
            }}>
                <BrowserRouter>
                    <Switch>
                        <Route exact={true} path={routePrefix + 'login'} render={(props) => {
                            return (
                                <Login
                                    routeProps={props}/>
                            );
                        }}/>
                        {
                            routes.map(((value, index) => {
                                return <Route exact={true} path={value.path} render={(props) => {
                                    return (
                                        <PrivateRoute component={value.component}
                                                      onSignOutClick={this.onSignOutClick}
                                                      routeProps={props}
                                                      routeRole={value.role}
                                        />
                                    );
                                }} key={index}/>;
                            }))
                        }
                        <Route path={'*'} render={NotFound}/>
                    </Switch>
                </BrowserRouter>
            </appInitialContext.Provider>
        );
    }

    changeAppStateByProperty(propertyName: keyof IAppState, value: any, shouldCallback: boolean) {
        if (shouldCallback) {
            this.setState({
                ...this.state,
                [propertyName]: value
            }, () => setTimeout(() => this.setState({
                ...this.state,
                [propertyName]: appInitialState[propertyName]
            }), 5000));
        } else {
            this.setState({
                ...this.state,
                [propertyName]: value
            });
        }
    }

    async onSignOutClick() {
        await authServiceObj.logout();
    }
}

export default App;
