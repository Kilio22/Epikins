import { IUser, userInitialState } from './IUser';
import { ChangeAppStateByProperty } from './Functions';
import React from 'react';

export interface IAppContext {
    changeAppStateByProperty: ChangeAppStateByProperty | null,
    user: IUser,
    errorMessage: string | null
}

export const appInitialContext: React.Context<IAppContext> = React.createContext<IAppContext>({
    changeAppStateByProperty: null,
    user: userInitialState,
    errorMessage: null
});
