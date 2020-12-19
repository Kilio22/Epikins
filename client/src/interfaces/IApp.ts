import {IUser, userInitialState} from "./IUser";

export interface IAppState {
    user: IUser,
    fuMode: boolean,
    errorMessage: string | null
}

export let appInitialState: IAppState = {
    user: userInitialState,
    fuMode: false,
    errorMessage: null
}