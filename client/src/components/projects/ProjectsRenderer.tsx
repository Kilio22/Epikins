import React from 'react';
import { Col, Form, Row } from 'react-bootstrap';
import Fuse from 'fuse.js';
import { NativeSelect, TextField } from '@material-ui/core';
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
        this.onModuleSelected = this.onModuleSelected.bind(this);
        this.getAvailableModules = this.getAvailableModules.bind(this);
        this.fuseProjects = this.fuseProjects.bind(this);
    }

    render() {
        let projects: IProject[] = this.props.projects;
        const availableModules = this.getAvailableModules(projects);

        projects = this.fuseProjects(projects, this.props.projects, this.state.queryString);
        projects = projects.filter((project) => {
            return this.state.selectedModule === 'All' || project.module === this.state.selectedModule;
        });
        return (
            <div>
                <TextField placeholder={'Project name'} variant={'standard'}
                           color={'primary'}
                           onChange={(event => this.onSearchFieldChange(event.target.value.trim()))}
                           className={'ml-1'}
                           autoFocus={true}/>
                <NativeSelect className={'ml-2'} variant={'standard'}
                              onChange={(event => {
                                  this.onModuleSelected(event.target.value);
                              })}
                              defaultValue={this.state.selectedModule}>
                    {
                        projects.length !== 0 &&
                        availableModules.map((value, index) => {
                            return (
                                <option key={index}>
                                    {value}
                                </option>
                            );
                        })
                    }
                </NativeSelect>
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

    getAvailableModules(projects: IProject[]) {
        let availableModules = projects.map((project) => {
            return project.module;
        });

        availableModules.push('All');
        availableModules = Array.from(new Set<string>(availableModules));
        availableModules = availableModules.sort((a, b) => {
            return a.localeCompare(b);
        });
        return availableModules;
    }

    onSearchFieldChange(value: string) {
        this.setState({
            queryString: value
        });
    }

    onModuleSelected(module: string) {
        this.setState({
            ...this.state,
            selectedModule: module
        });
    }

    fuseProjects(filteredProjects: IProject[], originalProjectList: IProject[], queryString: string) {
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

export default ProjectsRenderer;
