import Fuse from 'fuse.js';
import React from 'react';
import { appInitialContext } from '../../interfaces/IAppContext';
import { TextField } from '@material-ui/core';
import { Col, Row } from 'react-bootstrap';
import { IStudentJob } from '../../interfaces/myProjects/IStudentJob';
import StudentProjectRenderer from './StudentProjectRenderer';
import { OnStudentProjectClick } from '../../interfaces/Functions';

const projectsFuseOptions: Fuse.IFuseOptions<IStudentJob> = {
    shouldSort: true,
    threshold: 0.4,
    keys: [ 'project.name' ]
};

interface IStudentProjectsRendererState {
    queryString: string,
}

interface IStudentProjectsRendererProps {
    projects: IStudentJob[],
    onStudentProjectClick: OnStudentProjectClick,
    showSwitch: boolean
}

const studentProjectsRendererInitialState: IStudentProjectsRendererState = {
    queryString: ''
};

class StudentProjectsRenderer extends React.Component<IStudentProjectsRendererProps, IStudentProjectsRendererState> {
    static contextType = appInitialContext;
    context!: React.ContextType<typeof appInitialContext>;

    constructor(props: IStudentProjectsRendererProps) {
        super(props);

        this.state = studentProjectsRendererInitialState;
        this.onSearchFieldChange = this.onSearchFieldChange.bind(this);
        this.fuseProjects = this.fuseProjects.bind(this);
    }

    render() {
        let projects: IStudentJob[] = this.props.projects;

        projects = this.fuseProjects(projects, this.props.projects, this.state.queryString);
        return (
            <div>
                <TextField placeholder={'Project name'} variant={'standard'}
                           color={'primary'}
                           onChange={(event => this.onSearchFieldChange(event.target.value.trim()))}
                           className={'ml-1'}
                           autoFocus={true}/>
                {
                    projects.length === 0 ?
                        <h2 className={'text-center'}>No projects to display</h2>
                        :
                        <Row className={'mt-3'}>
                            {
                                projects.map((project, id) => {
                                    return (
                                        <Col md={4} key={id}>
                                            {
                                                <StudentProjectRenderer job={project}
                                                                        onStudentProjectClick={this.props.onStudentProjectClick}/>
                                            }
                                        </Col>
                                    );
                                })
                            }
                        </Row>
                }
            </div>
        );
    }

    onSearchFieldChange(value: string) {
        this.setState({
            queryString: value
        });
    }

    fuseProjects(filteredProjects: IStudentJob[], originalProjectList: IStudentJob[], queryString: string) {
        if (queryString !== '') {
            const fuse = new Fuse(originalProjectList, projectsFuseOptions);
            const fuseResult = fuse.search(queryString);

            filteredProjects = fuseResult.map(fuseRes => {
                return fuseRes.item;
            });
        }
        return filteredProjects;
    }
}

export default StudentProjectsRenderer;
