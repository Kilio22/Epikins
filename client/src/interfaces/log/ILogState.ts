import { IBuildLogInfo } from './IBuildLogInfo';

export interface ILogState {
    buildLogInfo: IBuildLogInfo | null,
    currentPage: number,
    isLoading: boolean,
    requestInProgressNb: number,
    projectString: string,
    starterString: string
}

export const LogInitialState: ILogState = {
    buildLogInfo: null,
    currentPage: 1,
    isLoading: false,
    requestInProgressNb: 0,
    projectString: '',
    starterString: ''
};
