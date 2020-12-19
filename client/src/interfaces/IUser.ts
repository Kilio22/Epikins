export interface IUser {
    name: string,
    email: string,
    isLoggedIn: boolean,
    canAccess: boolean
}

export const userInitialState: IUser = {
    name: "",
    email: "",
    isLoggedIn: false,
    canAccess: false
};