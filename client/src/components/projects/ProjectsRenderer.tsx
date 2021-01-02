import React from 'react';
import { Col, Form, Row } from 'react-bootstrap';
import Fuse from 'fuse.js';
import { TextField } from '@material-ui/core';
import {
    IProjectsRendererProps,
    IProjectsRendererState,
    projectsRendererInitialState
} from '../../interfaces/projects/IProjectsRenderer';
import { appInitialContext } from '../../interfaces/IAppContext';
import { IProject } from '../../interfaces/projects/IProject';

const projectsFuseOptions: Fuse.IFuseOptions<IProject> = {
    shouldSort: true,
    threshold: 0.4,
    keys: [ 'job.name' ]
};

class ProjectsRenderer extends React.Component<IProjectsRendererProps, IProjectsRendererState> {
    static contextType = appInitialContext;
    context!: React.ContextType<typeof appInitialContext>;

    constructor(props: IProjectsRendererProps) {
        super(props);

        this.state = projectsRendererInitialState;
        this.onSearchFieldChange = this.onSearchFieldChange.bind(this);
    }

    render() {
        let projects: IProject[] = this.props.projects;
        const fuse = new Fuse(this.props.projects, projectsFuseOptions);

        if (this.state.queryString !== '') {
            const fuseResult = fuse.search(this.state.queryString);

            projects = fuseResult.map(fuseRes => {
                return fuseRes.item;
            });
        }
        return (
            <div>
                <TextField placeholder={'Project name'} variant={'standard'}
                           color={'primary'}
                           onChange={(event => this.onSearchFieldChange(event.target.value.trim()))}
                           className={'ml-1'}
                           autoFocus={true}/>
                {
                    this.props.showSwitch &&
                    <Form className={'fu-switch p-1'}>
                        <Form.Check
                            type={'switch'}
                            id={'custom-switch'}
                            label={'Follow-up'}
                            checked={this.context.fuMode}
                            onChange={() => this.context.changeAppStateByProperty &&
                                this.context.changeAppStateByProperty('fuMode', !this.context.fuMode, false)}
                        />
                    </Form>
                }
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
                                                <this.props.ProjectRenderer onProjectClick={this.props.onProjectClick}
                                                                            project={project}/>
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

}

export default ProjectsRenderer;
