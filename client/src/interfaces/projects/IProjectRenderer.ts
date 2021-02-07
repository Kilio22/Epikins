import { IProject } from './IProject';
import { OnProjectClick } from '../Functions';

export interface IProjectRendererProps {
    onCheckboxClick: OnProjectClick | null,
    onProjectClick: OnProjectClick,
    project: IProject,
}
