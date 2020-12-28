export interface IApiUser {
    email: string,
    roles: string[],
    jenkinsLogin: string
}

export const apiUserInitialState: IApiUser = {
    email: '',
    roles: [],
    jenkinsLogin: ''
};
