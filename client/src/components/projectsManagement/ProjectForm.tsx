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
        let numberOfSelectedProjects = 0;

        for (let project of this.props.projects) {
            if (project.checked) {
                numberOfSelectedProjects++;
            }
        }
        return (
            <Modal show
                   onHide={() => this.props.changeProjectsManagementStateByProperty('clickedProject', null)}
                   size={'lg'}
                   centered>
                <Modal.Header closeButton>
                    <Modal.Title>
                        {
                            this.props.project.checked ?
                                'Change multiple projects build limit'
                                :
                                'Change project build limit'
                        }
                    </Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    {
                        this.props.project.checked ?
                            <span>
                                <span
                                    className={'font-weight-bold'}>Number of projects selected</span> {numberOfSelectedProjects}
                            </span>
                            :
                            <span>
                                <span
                                    className={'font-weight-bold'}>Associated module</span> {this.props.project.module}
                            </span>
                    }
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
        if (this.props.project.buildLimit === this.state.project.buildLimit && !this.props.project.checked) {
            this.props.changeProjectsManagementStateByProperty('clickedProject', null);
            return;
        }
        this.setState({
            ...this.state,
            isLoading: true
        });
        if (this.props.project.checked) {
            await this.changeSelectedProjectsBuildLimit();
            this.props.onSelectAllClick(false, []);
        } else {
            await this.changeProjectBuildLimit(this.state.project);
        }
        this.setState({
            ...this.state,
            isLoading: false
        });
        await this.props.getProjects();
        this.props.changeProjectsManagementStateByProperty('clickedProject', null);
    }

    async changeProjectBuildLimit(project: IProject) {
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            this.resetUser();
            return;
        }

        const res = await EpikinsApiService.changeProjectBuildLimit(project, accessToken);
        if (!res) {
            this.setErrorMessage('Cannot change build limit, please try to reload the page.');
        }
    }

    async changeSelectedProjectsBuildLimit() {
        for (let project of this.props.projects) {
            if (project.checked) {
                project.buildLimit = this.state.project.buildLimit;
                await this.changeProjectBuildLimit(project);
            }
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
