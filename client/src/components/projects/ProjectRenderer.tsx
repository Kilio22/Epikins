import { Button } from 'react-bootstrap';
import React from 'react';
import { IProjectRendererProps } from '../../interfaces/projects/IProjectRenderer';

const ProjectRenderer: React.FunctionComponent<IProjectRendererProps> = ({
                                                                             project,
                                                                             onCheckboxClick,
                                                                             onProjectClick
                                                                         }) => {
    return (
        <Button variant={'outline-primary'}
                className={'m-1 text-left'}
                block={true}
                onClick={() => onProjectClick(project)}>
            {
                onCheckboxClick &&
                <div>
                    <input
                        className={'jobs-checkbox mr-2'}
                        type={'checkbox'}
                        checked={project.checked}
                        onChange={(() => onCheckboxClick(project))}/>
                    {' '}
                </div>
            }
            <i className={'fas fa-folder'}/> {project.job.name}
        </Button>
    );
};

export default ProjectRenderer;
