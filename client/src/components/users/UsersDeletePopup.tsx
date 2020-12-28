import * as React from 'react';
import { Button, Modal } from 'react-bootstrap';
import { IUsersDeletePopupProps } from '../../interfaces/users/IUsersDeletePopup';

const UsersDeletePopup: React.FunctionComponent<IUsersDeletePopupProps> = ({
                                                                               onDeleteClick,
                                                                               changeUsersStateByProperty
                                                                           }) => {
    return (
        <Modal show size={'lg'} onHide={() => changeUsersStateByProperty('isDeleting', false)} centered>
            <Modal.Header closeButton>
                <Modal.Title>Are you sure?</Modal.Title>
            </Modal.Header>
            <Modal.Body>If you remove this person from this list, he will not be able to access to the app
                anymore.</Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={() => changeUsersStateByProperty('isDeleting', false)}>
                    Cancel
                </Button>
                <Button variant="danger" onClick={onDeleteClick}>
                    Delete
                </Button>
            </Modal.Footer>
        </Modal>
    );
};

export default UsersDeletePopup;
