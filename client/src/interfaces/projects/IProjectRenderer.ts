import { IProject } from './IProject';
import { OnProjectClick } from '../Functions';

export interface IProjectRendererProps {
    project: IProject
    onProjectClick: OnProjectClick
}
