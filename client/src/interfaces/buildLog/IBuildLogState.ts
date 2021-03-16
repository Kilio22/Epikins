import { IBuildLogInfo } from './IBuildLogInfo';

export interface IBuildLogState {
    buildLogInfo: IBuildLogInfo | null,
    cities: string[],
    currentPage: number,
    isLoading: boolean,
    requestInProgressNb: number,
    projectString: string,
    selectedCity: string,
    showExportForm: boolean,
    starterString: string
}

export const LogInitialState: IBuildLogState = {
    buildLogInfo: null,
    cities: [],
    currentPage: 1,
    isLoading: false,
    requestInProgressNb: 0,
    projectString: '',
    selectedCity: '',
    showExportForm: false,
    starterString: ''
};
