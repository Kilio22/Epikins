import * as React from 'react';
import { Button, Modal } from 'react-bootstrap';
import { IJenkinsCredentialsDeletePopupProps } from '../../interfaces/IJenkinsCredentialsTable/IJenkinsCredentialsDeletePopup';

const JenkinsCredentialsDeletePopup: React.FunctionComponent<IJenkinsCredentialsDeletePopupProps> = ({
                                                                                                         onDeleteClick,
                                                                                                         changeJenkinsCredentialStateByProperty
                                                                                                     }) => {
    return (
        <Modal show size={'lg'} onHide={() => changeJenkinsCredentialStateByProperty('isDeleting', false)} centered>
            <Modal.Header closeButton>
                <Modal.Title>Are you sure?</Modal.Title>
            </Modal.Header>
            <Modal.Body>If you delete these credentials, you will not be able to get back these and every user
                associated with these credentials will not be able to access to projects and start builds.</Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={() => changeJenkinsCredentialStateByProperty('isDeleting', false)}>
                    Cancel
                </Button>
                <Button variant="danger" onClick={onDeleteClick}>
                    Delete
                </Button>
            </Modal.Footer>
        </Modal>
    );
};

export default JenkinsCredentialsDeletePopup;
