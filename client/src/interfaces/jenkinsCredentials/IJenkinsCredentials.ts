export type ChangeJenkinsCredentialsStateByProperty = (propertyName: keyof IJenkinsCredentialsState, value: any) => void;

export interface IJenkinsCredentialsState {
    jenkinsCredentials: string[],
    isAdding: boolean,
    showDeletePopup: boolean,
    toDelete: string,
    isLoading: boolean
}

export const jenkinsCredentialsInitialState: IJenkinsCredentialsState = {
    jenkinsCredentials: [],
    isAdding: false,
    showDeletePopup: false,
    toDelete: '',
    isLoading: false
};
