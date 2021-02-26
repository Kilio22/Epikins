import React from 'react';
import { Toast } from 'react-bootstrap';
import { OnButtonClick } from '../interfaces/Functions';

interface BuildStartedToastProps {
    onClose: OnButtonClick
}

const BuildStartedToast: React.FunctionComponent<BuildStartedToastProps> = ({onClose}) => {
    return (
        <Toast show={true} onClose={onClose} delay={6000} autohide className={'build-toast'} animation={true}>
            <Toast.Header>
                <i className={'fas fa-check-square mr-1 check-build'}/>
                <strong className={'mr-auto'}>Build started</strong>
            </Toast.Header>
            <Toast.Body>Be patient, you'll get the result soon.</Toast.Body>
        </Toast>
    );
};

export default BuildStartedToast;
