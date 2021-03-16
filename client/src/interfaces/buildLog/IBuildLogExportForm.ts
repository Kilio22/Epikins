import { IBuildLogState } from './IBuildLogState';

export type ChangeBuildLogStateByProperty = (key: keyof IBuildLogState, value: any) => void;

export interface IBuildLogExportFormProps {
    changeBuildLogStateByProperty: ChangeBuildLogStateByProperty,
    cities: string[]
}

export interface IBuildLogExportFormState {
    endDate: Date
    isLoading: boolean,
    selectedCity: string,
    startDate: Date,
    projectValue: string
}

export const buildLogExportFormInitialState: IBuildLogExportFormState = {
    endDate: new Date(),
    isLoading: false,
    selectedCity: '',
    startDate: new Date(),
    projectValue: ''
};
