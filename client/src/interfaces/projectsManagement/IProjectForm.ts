import { IProject } from '../projects/IProject';
import { OnButtonClick, OnCheckboxChange } from '../Functions';
import { ChangeProjectsManagementStateByProperty } from './IProjectsManagement';

export interface IProjectFormProps {
    onSelectAllClick: OnCheckboxChange<IProject[]>
    changeProjectsManagementStateByProperty: ChangeProjectsManagementStateByProperty,
    getProjects: OnButtonClick
    project: IProject,
    projects: IProject[],
}

export interface IProjectFormState {
    project: IProject,
    isLoading: boolean
}
