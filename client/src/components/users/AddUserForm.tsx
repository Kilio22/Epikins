import React, { Component } from 'react';
import { Button, Form, Modal, Spinner } from 'react-bootstrap';
import { roles } from './UsersTableHeader';
import { apiUserInitialState, IApiUser } from '../../interfaces/IApiUser';
import { appInitialContext } from '../../interfaces/IAppContext';
import { authServiceObj } from '../../services/AuthService';
import EpikinsApiService from '../../services/EpikinsApiService';
import { userInitialState } from '../../interfaces/IUser';
import { IAddUserFormProps, IAddUserFormState, usersFormInitialState } from '../../interfaces/users/IAddUserForm';

class AddUserForm extends Component<IAddUserFormProps, IAddUserFormState> {
    static contextType = appInitialContext;
    context!: React.ContextType<typeof appInitialContext>;

    constructor(props: IAddUserFormProps) {
        super(props);

        this.onSubmit = this.onSubmit.bind(this);
        this.addUser = this.addUser.bind(this);
        this.onCheckboxChange = this.onCheckboxChange.bind(this);
        this.changeUserByProperty = this.changeUserByProperty.bind(this);
        this.resetUser = this.resetUser.bind(this);
        this.setErrorMessage = this.setErrorMessage.bind(this);

        this.state = {
            ...usersFormInitialState,
            user: {
                ...apiUserInitialState,
                jenkinsLogin: this.props.jenkinsCredentials[0]
            }
        };
    }

    render() {
        return (
            <Modal show
                   onHide={() => this.props.changeUsersStateByProperty('isAdding', false)}
                   size="lg"
                   centered>
                <Modal.Header closeButton>
                    <Modal.Title>
                        Add user
                    </Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <Form noValidate validated={this.state.validation} onSubmit={event => this.onSubmit(event)}>
                        <Form.Group>
                            <Form.Label>Email address</Form.Label>
                            <Form.Control type="email"
                                          placeholder="Enter email"
                                          required
                                          onChange={(event => this.changeUserByProperty('email', event.target.value.trim()))}/>
                            <Form.Control.Feedback type="invalid">
                                Please provide a valid email.
                            </Form.Control.Feedback>
                        </Form.Group>
                        <Form.Group>
                            <Form.Label>Roles</Form.Label>
                            {
                                roles.map((role, index) => {
                                    return (
                                        <Form.Check type="checkbox" label={role} key={index}
                                                    checked={this.state.user.roles.includes(role.toLocaleLowerCase())}
                                                    onChange={() => this.onCheckboxChange(role.toLocaleLowerCase())}/>
                                    );
                                })
                            }
                        </Form.Group>
                        <Form.Group>
                            <Form.Label>Jenkins login</Form.Label>
                            <Form.Control as={'select'}
                                          onChange={event => this.changeUserByProperty('jenkinsLogin', event.target.value)}
                                          defaultValue={this.props.jenkinsCredentials[0]}>
                                {
                                    this.props.jenkinsCredentials.map((value, index) => {
                                        return (
                                            <option key={index}>
                                                {value}
                                            </option>
                                        );
                                    })
                                }
                            </Form.Control>
                        </Form.Group>
                        <div className={'d-flex justify-content-center'}>
                            <Button variant="primary" type="submit" disabled={this.state.isLoading}>
                                {
                                    this.state.isLoading ?
                                        <Spinner animation={'border'}/>
                                        :
                                        'Submit'
                                }
                            </Button>
                        </div>
                    </Form>
                </Modal.Body>
            </Modal>
        );
    }

    onCheckboxChange(role: string) {
        const idx = this.state.user.roles.indexOf(role);

        if (idx !== -1) {
            let roles = [ ...this.state.user.roles ];

            roles.splice(idx, 1);
            this.setState({
                ...this.state,
                user: {
                    ...this.state.user,
                    roles: roles
                }
            });
        } else {
            this.setState({
                ...this.state,
                user: {
                    ...this.state.user,
                    roles: this.state.user.roles.concat(role)
                }
            });
        }
    }

    changeUserByProperty(key: keyof IApiUser, value: any) {
        this.setState({
            ...this.state,
            user: {
                ...this.state.user,
                [key]: value
            }
        });
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

    async onSubmit(event: React.FormEvent<HTMLElement>) {
        event.preventDefault();
        event.stopPropagation();
        // @ts-ignore
        if (event.currentTarget.checkValidity() === false) {
            this.setState({
                ...this.state,
                validation: true
            });
            return;
        }

        this.setState({
            ...this.state,
            isLoading: true
        });
        await this.addUser(this.state.user);
        this.setState({
            ...this.state,
            isLoading: false
        });
        this.props.changeUsersStateByProperty('isAdding', false);
    }

    async addUser(newUser: IApiUser) {
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            this.resetUser();
            return;
        }

        const res = await EpikinsApiService.addUser(newUser, accessToken);
        if (res === 409) {
            this.setErrorMessage('Cannot add user: a user with the same email already exists.');
        } else if (res !== 201) {
            this.setErrorMessage('Cannot add user, please try to reload the page.');
        }
        await this.props.getUsers();
    }
}

export default AddUserForm;
