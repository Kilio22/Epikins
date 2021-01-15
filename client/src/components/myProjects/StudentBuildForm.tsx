import React, { Component } from 'react';
import { appInitialContext } from '../../interfaces/IAppContext';
import { Button, Modal, Spinner } from 'react-bootstrap';
import { userInitialState } from '../../interfaces/IUser';
import { ChangeMyProjectsStateByProperty, OnButtonClick } from '../../interfaces/Functions';
import { IStudentJob } from '../../interfaces/myProjects/IStudentJob';

interface IStudentBuildFormProps {
    changeMyProjectsStateByProperty: ChangeMyProjectsStateByProperty,
    getStudentProjects: OnButtonClick,
    startBuild: OnButtonClick,
    selectedProject: IStudentJob,
}

interface IStudentBuildFormState {
    isLoading: boolean
}

const StudentBuildFormInitialState: IStudentBuildFormState = {
    isLoading: false
};

class StudentBuildForm extends Component<IStudentBuildFormProps, IStudentBuildFormState> {
    static contextType = appInitialContext;
    context!: React.ContextType<typeof appInitialContext>;

    constructor(props: IStudentBuildFormProps) {
        super(props);

        this.onBuildClick = this.onBuildClick.bind(this);
        this.resetUser = this.resetUser.bind(this);
        this.setErrorMessage = this.setErrorMessage.bind(this);

        this.state = StudentBuildFormInitialState;
    }

    render() {
        return (
            <Modal show
                   onHide={() => this.props.changeMyProjectsStateByProperty('showForm', false)}
                   size={'lg'}
                   centered>
                <Modal.Header closeButton>
                    <Modal.Title>
                        Start a build
                    </Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <h4>Do you really want to start a build for {this.props.selectedProject.project.name} project?</h4>
                </Modal.Body>
                <Modal.Footer>
                    <div className={'d-flex justify-content-center'}>
                        {
                            this.state.isLoading ?
                                <Spinner animation={'border'}/>
                                :
                                <div>
                                    <Button variant="secondary"
                                            onClick={() => this.props.changeMyProjectsStateByProperty('showForm', false)}
                                            className={'mr-2'}>
                                        Cancel
                                    </Button>
                                    <Button variant={'primary'} onClick={this.onBuildClick}>
                                        Start
                                    </Button>
                                </div>
                        }
                    </div>
                </Modal.Footer>
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

    async onBuildClick() {
        this.setState({
            ...this.state,
            isLoading: true
        });
        await this.props.startBuild();
        await this.props.getStudentProjects();
        this.setState({
            ...this.state,
            isLoading: false
        });
        this.props.changeMyProjectsStateByProperty('showForm', false);
    }
}

export default StudentBuildForm;
