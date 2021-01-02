import React, { Component } from 'react';
import { IProject } from '../../interfaces/projects/IProject';
import { Button, Form, Modal, Spinner } from 'react-bootstrap';
import { userInitialState } from '../../interfaces/IUser';
import { appInitialContext } from '../../interfaces/IAppContext';
import { authServiceObj } from '../../services/AuthService';
import EpikinsApiService from '../../services/EpikinsApiService';
import { IProjectFormProps, IProjectFormState } from '../../interfaces/projectsManagement/IProjectForm';

class ProjectForm extends Component<IProjectFormProps, IProjectFormState> {
    static contextType = appInitialContext;
    context!: React.ContextType<typeof appInitialContext>;

    constructor(props: IProjectFormProps) {
        super(props);

        this.changeProjectByProperty = this.changeProjectByProperty.bind(this);
        this.onSubmit = this.onSubmit.bind(this);
        this.changeProjectBuildLimit = this.changeProjectBuildLimit.bind(this);
        this.resetUser = this.resetUser.bind(this);
        this.setErrorMessage = this.setErrorMessage.bind(this);

        this.state = {
            isLoading: false,
            project: {
                ...this.props.project
            }
        };
    }

    render() {
        return (
            <Modal show
                   onHide={() => this.props.changeProjectsManagementStateByProperty('selectedProject', null)}
                   size={'lg'}
                   centered>
                <Modal.Header closeButton>
                    <Modal.Title>
                        Change project build limit
                    </Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <span>
                        <span className={'font-weight-bold'}>Associated module</span> {this.props.project.module}
                    </span>
                    <Form className={'mt-2'} noValidate onSubmit={event => this.onSubmit(event)}>
                        <Form.Group>
                            <Form.Label>Build limit</Form.Label>
                            <Form.Control type={'number'}
                                          min={0}
                                          required
                                          onChange={(event => this.changeProjectByProperty('buildLimit', Number(event.target.value)))}
                                          value={this.state.project.buildLimit}/>
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
        if (this.props.project.buildLimit === this.state.project.buildLimit) {
            this.props.changeProjectsManagementStateByProperty('selectedProject', null);
            return;
        }
        this.setState({
            ...this.state,
            isLoading: true
        });
        await this.changeProjectBuildLimit();
        this.setState({
            ...this.state,
            isLoading: false
        });
        await this.props.getProjects();
        this.props.changeProjectsManagementStateByProperty('selectedProject', null);
    }

    async changeProjectBuildLimit() {
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            this.resetUser();
            return;
        }

        const res = await EpikinsApiService.changeProjectBuildLimit(this.state.project, accessToken);
        if (!res) {
            this.setErrorMessage('Cannot change build limit, please try to reload the page.');
        }
    }

    changeProjectByProperty(key: keyof IProject, value: any) {
        this.setState({
            ...this.state,
            project: {
                ...this.state.project,
                [key]: value
            }
        });
    }
}

export default ProjectForm;
