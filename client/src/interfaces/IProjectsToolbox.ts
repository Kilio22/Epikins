import { HandleType, OnCheckboxChange, OnSelectChange, OnSelectSearchChange } from './Functions';
import { IProject } from './projects/IProject';
import { ISelectOption } from './projects/ISelectOption';

export interface IProjectsToolbox {
    allSelected: boolean,
    handleString: HandleType<string>,
    onSelectAllClick: OnCheckboxChange<IProject[]> | null,
    onSelectChange: OnSelectChange,
    onSelectSearchChange: OnSelectSearchChange,
    projects: IProject[],
    selectOptions: ISelectOption[],
    selectSearch: string
}
