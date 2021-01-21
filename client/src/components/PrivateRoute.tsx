import React, { ComponentClass, FunctionComponent } from 'react';
import Unauthorized from './Unauthorized';
import { Redirect, RouteComponentProps } from 'react-router-dom';
import { routePrefix } from '../interfaces/IRoute';
import NavBar from './NavBar';
import Footer from './Footer';
import { appInitialContext } from '../interfaces/IAppContext';
import { Alert, Container } from 'react-bootstrap';
import { OnSignOutClick } from '../interfaces/Functions';

interface IPrivateRoute {
    component: ComponentClass<any> | FunctionComponent<any>
    onSignOutClick: OnSignOutClick,
    routeProps: RouteComponentProps,
    routeRole: string
}

const PrivateRoute: React.FunctionComponent<IPrivateRoute> = ({
                                                                  component: Component,
                                                                  onSignOutClick,
                                                                  routeProps,
                                                                  routeRole
                                                              }) => {
    return (
        <appInitialContext.Consumer>
            {context => (
                context.user.isLoggedIn ?
                    <div className={'d-flex flex-column min-vh-100'}>
                        <NavBar routeProps={routeProps} onSignOutClick={onSignOutClick} user={context.user}/>
                        <div className={'d-flex flex-grow-1'}>
                            {
                                (routeRole === '' || context.user.roles.includes(routeRole)) ?
                                    <Container className={'mt-3'}>
                                        {
                                            context.errorMessage &&
                                            <Alert variant={'danger'}>{context.errorMessage}</Alert>
                                        }
                                        <Component
                                            routeProps={routeProps}/>
                                    </Container>
                                    :
                                    <Unauthorized/>
                            }
                        </div>
                        <Footer/>
                    </div>
                    :
                    <Redirect
                        to={{
                            pathname: routePrefix + 'login',
                            state: {from: routeProps.location}
                        }}
                    />
            )}
        </appInitialContext.Consumer>

    );
};

export default PrivateRoute;
