import { OnButtonClick } from '../Functions';
import { apiJenkinsCredentialsInitialState, IApiJenkinsCredentials } from './IApiJenkinsCredentials';
import { ChangeJenkinsCredentialsStateByProperty } from './IJenkinsCredentials';

export interface IJenkinsCredentialsFormProps {
    changeJenkinsCredentialsStateByProperty: ChangeJenkinsCredentialsStateByProperty,
    getJenkinsCredentials: OnButtonClick,
    jenkinsCredentials: string[]
}

export interface IJenkinsCredentialsFormState {
    jenkinsCredentials: IApiJenkinsCredentials,
    validation: boolean,
    isLoading: boolean
}

export const AddJenkinsCredentialsFormInitialState: IJenkinsCredentialsFormState = {
    jenkinsCredentials: apiJenkinsCredentialsInitialState,
    validation: false,
    isLoading: false
};
