import React from 'react';
import { Button } from 'react-bootstrap';
import { IProjectRendererProps } from '../../interfaces/projects/IProjectRenderer';

const ProjectBuildLimitRenderer: React.FunctionComponent<IProjectRendererProps> = ({
                                                                                       onCheckboxClick,
                                                                                       project,
                                                                                       onProjectClick
                                                                                   }) => {
    return (
        <Button variant={'outline-primary'}
                className={'m-1 text-left d-flex align-items-center'}
                block={true}
                onClick={() => onProjectClick(project)}>
            {
                onCheckboxClick &&
                <div className={'d-flex align-items-center'}>
                    <input
                        className={'jobs-checkbox mr-2'}
                        type={'checkbox'}
                        checked={project.checked}
                        onChange={(() => onCheckboxClick(project))}/>
                    {' '}
                </div>
            }
            <i className={'fas fa-folder mr-1'}/> {project.job.name}
            <div className={'ml-auto mr-1'}>
                [{project.buildLimit !== 0 ?
                <span className={'font-weight-bold build-limit-green'}>{project.buildLimit}</span>
                :
                <span className={'build-limit'}>{project.buildLimit}</span>}]
            </div>
        </Button>
    );
};

export default ProjectBuildLimitRenderer;
