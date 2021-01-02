export interface IApiJenkinsCredentials {
    login: string,
    apiKey: string
}

export const apiJenkinsCredentialsInitialState: IApiJenkinsCredentials = {
    login: '',
    apiKey: ''
};
