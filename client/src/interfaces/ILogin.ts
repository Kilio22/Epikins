import {RouteComponentProps} from "react-router-dom";
import {StaticContext} from "react-router";
import {AccountInfo} from "@azure/msal-browser";

export interface ILoginState {
    isConnecting: boolean,
    account: AccountInfo | null
}

export interface ILoginProps {
    routeProps: RouteComponentProps<{}, StaticContext, any>
}

export const loginInitialState: ILoginState = {
    isConnecting: false,
    account: null
}
