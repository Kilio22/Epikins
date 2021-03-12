import { IBuildLogInfo } from './IBuildLogInfo';

export interface IBuildLogState {
    buildLogInfo: IBuildLogInfo | null,
    currentPage: number,
    isLoading: boolean,
    requestInProgressNb: number,
    projectString: string,
    starterString: string
}

export const LogInitialState: IBuildLogState = {
    buildLogInfo: null,
    currentPage: 1,
    isLoading: false,
    requestInProgressNb: 0,
    projectString: '',
    starterString: ''
};
