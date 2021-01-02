import { Button } from 'react-bootstrap';
import React from 'react';
import { IProjectRendererProps } from '../../interfaces/projects/IProjectRenderer';

const ProjectRenderer: React.FunctionComponent<IProjectRendererProps> = ({project, onProjectClick}) => {
    return (
        <Button variant={'outline-primary'}
                className={'m-1 text-left'}
                block={true}
                onClick={() => onProjectClick(project)}>
            <i className={'fas fa-folder'}/> {project.job.name}
        </Button>
    );
};

export default ProjectRenderer;
