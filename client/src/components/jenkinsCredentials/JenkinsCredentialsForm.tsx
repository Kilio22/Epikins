import React, { Component } from 'react';
import { Button, Form, Modal, Spinner } from 'react-bootstrap';
import { appInitialContext } from '../../interfaces/IAppContext';
import { authServiceObj } from '../../services/AuthService';
import EpikinsApiService from '../../services/EpikinsApiService';
import { userInitialState } from '../../interfaces/IUser';
import {
    AddJenkinsCredentialsFormInitialState,
    IJenkinsCredentialsFormProps,
    IJenkinsCredentialsFormState
} from '../../interfaces/jenkinsCredentials/IJenkinsCredentialsForm';
import { IApiJenkinsCredentials } from '../../interfaces/jenkinsCredentials/IApiJenkinsCredentials';

class JenkinsCredentialsForm extends Component<IJenkinsCredentialsFormProps, IJenkinsCredentialsFormState> {
    static contextType = appInitialContext;
    context!: React.ContextType<typeof appInitialContext>;

    constructor(props: IJenkinsCredentialsFormProps) {
        super(props);

        this.onSubmit = this.onSubmit.bind(this);
        this.addJenkinsCredentials = this.addJenkinsCredentials.bind(this);
        this.changeJenkinsCredentialsByProperty = this.changeJenkinsCredentialsByProperty.bind(this);
        this.resetUser = this.resetUser.bind(this);
        this.setErrorMessage = this.setErrorMessage.bind(this);
        this.checkFormValidity = this.checkFormValidity.bind(this);

        this.state = AddJenkinsCredentialsFormInitialState;
    }

    render() {
        return (
            <Modal show
                   onHide={() => this.props.changeJenkinsCredentialsStateByProperty('isAdding', false)}
                   size={'lg'}
                   centered>
                <Modal.Header closeButton>
                    <Modal.Title>
                        Add user
                    </Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <Form noValidate validated={this.state.validation} onSubmit={event => this.onSubmit(event)}>
                        <Form.Group>
                            <Form.Label>Jenkins login</Form.Label>
                            <Form.Control type={'text'}
                                          placeholder={'Enter jenkins login'}
                                          required
                                          onChange={(event => this.changeJenkinsCredentialsByProperty('login', event.target.value.trim()))}/>
                            <Form.Control.Feedback type={'invalid'}>
                                Please provide a jenkins login.
                            </Form.Control.Feedback>
                        </Form.Group>
                        <Form.Group>
                            <Form.Label>Associated API key</Form.Label>
                            <Form.Control type={'text'}
                                          placeholder={'Enter API key'}
                                          required
                                          onChange={(event => this.changeJenkinsCredentialsByProperty('apiKey', event.target.value.trim()))}/>
                            <Form.Control.Feedback type={'invalid'}>
                                Please provide an API key.
                            </Form.Control.Feedback>
                        </Form.Group>
                        <div className={'d-flex justify-content-center'}>
                            <Button variant={'primary'} type={'submit'} disabled={this.state.isLoading}>
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

    changeJenkinsCredentialsByProperty(key: keyof IApiJenkinsCredentials, value: any) {
        this.setState({
            ...this.state,
            jenkinsCredentials: {
                ...this.state.jenkinsCredentials,
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

    checkFormValidity(event: React.FormEvent<HTMLElement>): boolean {
        // @ts-ignore
        if (event.currentTarget.checkValidity() === false) {
            this.setState({
                ...this.state,
                validation: true
            });
            return false;
        }
        if (this.state.jenkinsCredentials.apiKey.trim() === '' || this.state.jenkinsCredentials.login.trim() === '') {
            this.setErrorMessage('Login or API key must not be formed of spaces only.');
            return false;
        }
        return true;
    }

    async onSubmit(event: React.FormEvent<HTMLElement>) {
        event.preventDefault();
        event.stopPropagation();
        if (!this.checkFormValidity(event)) {
            return;
        }
        this.setState({
            ...this.state,
            isLoading: true
        });
        await this.addJenkinsCredentials(this.state.jenkinsCredentials);
        this.setState({
            ...this.state,
            isLoading: false
        });
        await this.props.getJenkinsCredentials();
        this.props.changeJenkinsCredentialsStateByProperty('isAdding', false);
    }

    async addJenkinsCredentials(newCredentials: IApiJenkinsCredentials) {
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            this.resetUser();
            return;
        }

        const res = await EpikinsApiService.addJenkinsCredentials(newCredentials, accessToken);
        if (res === 409) {
            this.setErrorMessage('Cannot add credentials: credentials with the same login already exists.');
        } else if (res !== 201) {
            this.setErrorMessage('Cannot add credentials, please try to reload the page.');
        }
    }
}

export default JenkinsCredentialsForm;
