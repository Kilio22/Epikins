export type ChangeJenkinsCredentialsStateByProperty = (propertyName: keyof IJenkinsCredentialsState, value: any) => void;

export interface IJenkinsCredentialsState {
    jenkinsCredentials: string[],
    isAdding: boolean,
    isDeleting: boolean,
    toDelete: string,
    isLoading: boolean
}

export const jenkinsCredentialsInitialState: IJenkinsCredentialsState = {
    jenkinsCredentials: [],
    isAdding: false,
    isDeleting: false,
    toDelete: '',
    isLoading: false
};
