import { OnButtonClick } from '../Functions';
import { ChangeJenkinsCredentialsStateByProperty } from './IJenkinsCredentials';

export interface IJenkinsCredentialsDeletePopupProps {
    onDeleteClick: OnButtonClick,
    changeJenkinsCredentialStateByProperty: ChangeJenkinsCredentialsStateByProperty
}
