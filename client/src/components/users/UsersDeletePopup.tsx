import * as React from 'react';
import { Button, Modal, Spinner } from 'react-bootstrap';
import {
    IUsersDeletePopupProps,
    IUsersDeletePopupState,
    usersDeletePopupInitialState
} from '../../interfaces/users/IUsersDeletePopup';


class UsersDeletePopup extends React.Component<IUsersDeletePopupProps, IUsersDeletePopupState> {
    constructor(props: IUsersDeletePopupProps) {
        super(props);

        this.state = usersDeletePopupInitialState;
    }

    render() {
        return (
            <Modal show size={'lg'} onHide={() => this.props.changeUsersStateByProperty('showDeletePopup', false)}
                   centered>
                <Modal.Header closeButton>
                    <Modal.Title>Are you sure?</Modal.Title>
                </Modal.Header>
                <Modal.Body>If you remove this person from this list, he will not be able to access to the app
                    anymore.</Modal.Body>
                <Modal.Footer>
                    <Button variant="secondary"
                            onClick={() => this.props.changeUsersStateByProperty('showDeletePopup', false)}
                            disabled={this.state.isLoading}>
                        Cancel
                    </Button>
                    <Button variant="danger" onClick={async () => {
                        this.setState({
                            isLoading: true
                        });
                        await this.props.onDeleteClick();
                    }}
                            disabled={this.state.isLoading}>
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

export default UsersDeletePopup;
