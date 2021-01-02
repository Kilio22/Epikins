import { IProject } from '../projects/IProject';
import { OnButtonClick } from '../Functions';
import { ChangeProjectsManagementStateByProperty } from './IProjectsManagement';

export interface IProjectFormProps {
    project: IProject,
    changeProjectsManagementStateByProperty: ChangeProjectsManagementStateByProperty,
    getProjects: OnButtonClick
}

export interface IProjectFormState {
    project: IProject,
    isLoading: boolean
}
