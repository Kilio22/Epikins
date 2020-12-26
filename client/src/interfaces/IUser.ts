export interface IUser {
    name: string,
    email: string,
    isLoggedIn: boolean,
    roles: string[]
}

export const userInitialState: IUser = {
    name: '',
    email: '',
    isLoggedIn: false,
    roles: []
};
