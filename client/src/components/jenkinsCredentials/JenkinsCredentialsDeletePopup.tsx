import * as React from 'react';
import { Button, Modal, Spinner } from 'react-bootstrap';
import { IJenkinsCredentialsDeletePopupProps } from '../../interfaces/jenkinsCredentials/IJenkinsCredentialsDeletePopup';

interface IJenkinsCredentialsDeletePopupState {
    isLoading: boolean
}

const jenkinsCredentialsDeletePopupInitialState: IJenkinsCredentialsDeletePopupState = {
    isLoading: false
};

class JenkinsCredentialsDeletePopup extends React.Component<IJenkinsCredentialsDeletePopupProps, IJenkinsCredentialsDeletePopupState> {
    constructor(props: IJenkinsCredentialsDeletePopupProps) {
        super(props);

        this.state = jenkinsCredentialsDeletePopupInitialState;
    }

    render() {
        return (
            <Modal show size={'lg'}
                   onHide={() => this.props.changeJenkinsCredentialStateByProperty('showDeletePopup', false)} centered>
                <Modal.Header closeButton>
                    <Modal.Title>Are you sure?</Modal.Title>
                </Modal.Header>
                <Modal.Body>If you delete these credentials, you will not be able to get back these and every user
                    associated with these credentials will not be able to access to projects and start
                    builds.</Modal.Body>
                <Modal.Footer>
                    <Button variant="secondary"
                            onClick={() => this.props.changeJenkinsCredentialStateByProperty('showDeletePopup', false)}
                            disabled={this.state.isLoading}>
                        Cancel
                    </Button>
                    <Button variant="danger" onClick={async () => {
                        this.setState({
                            isLoading: true
                        });
                        await this.props.onDeleteClick();
                    }} disabled={this.state.isLoading}>
                        {
                            this.state.isLoading ?
                                <Spinner animation={'border'}/>
                                :
                                'Delete'
                        }
                    </Button>
                </Modal.Footer>
            </Modal>
        );
    }
}

export default JenkinsCredentialsDeletePopup;
