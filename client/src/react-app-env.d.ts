/// <reference types="react-scripts" />
declare namespace NodeJS {
    interface ProcessEnv {
        REACT_APP_APP_ID: string
        REACT_APP_TENANT_ID: string
        REACT_APP_REDIRECT_URI: string
        REACT_APP_API_BASE_URI: string
        REACT_APP_API_SCOPE: string
    }
}
