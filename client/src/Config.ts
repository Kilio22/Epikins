import { Configuration } from '@azure/msal-browser';

export const apiBaseURI: string = process.env.REACT_APP_API_BASE_URI;

export const appId: string = process.env.REACT_APP_APP_ID;
export const redirectUri: string = process.env.REACT_APP_REDIRECT_URI;
export const apiScope: string = process.env.REACT_APP_API_SCOPE;

export const tenantJWKSURI: string = 'https://login.microsoftonline.com/' + process.env.REACT_APP_TENANT_ID + '/discovery/v2.0/keys';
export const IssuerURL: string = 'https://login.microsoftonline.com/' + process.env.REACT_APP_TENANT_ID + '/v2.0';

export const connectionScopes: string[] = [
    apiScope,
    'offline_access'
];
export const msalConfig: Configuration = {
    auth: {
        clientId: appId,
        redirectUri: redirectUri,
        authority: 'https://login.microsoftonline.com/' + process.env.REACT_APP_TENANT_ID
    },
    cache: {
        cacheLocation: 'localStorage',
        storeAuthStateInCookie: true
    }
    /*system: {
     loggerOptions: {
     loggerCallback: (level: LogLevel, message: string, containsPii: boolean): void => {
     if (containsPii) {
     return;
     }
     switch (level) {
     case LogLevel.Error:
     console.log(message);
     return;
     case LogLevel.Info:
     console.log(message);
     return;
     case LogLevel.Verbose:
     console.log(message);
     return;
     case LogLevel.Warning:
     console.log(message);
     return;
     }
     }
     }
     }*/
};
