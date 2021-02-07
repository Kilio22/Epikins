import React from 'react';
import { Col, Row } from 'react-bootstrap';
import Fuse from 'fuse.js';
import { TextField } from '@material-ui/core';
import {
    IProjectsRendererProps,
    IProjectsRendererState,
    projectsRendererInitialState
} from '../../interfaces/projects/IProjectsRenderer';
import { appInitialContext } from '../../interfaces/IAppContext';
import { IProject } from '../../interfaces/projects/IProject';
import Select, { InputActionMeta, ValueType } from 'react-select';
import { ISelectOption } from '../../interfaces/projects/ISelectOption';

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
        this.onSelectChange = this.onSelectChange.bind(this);
        this.getAvailableModules = this.getAvailableModules.bind(this);
        this.fuseProjects = this.fuseProjects.bind(this);
        this.onSelectSearchChange = this.onSelectSearchChange.bind(this);
    }

    render() {
        const availableModules = this.getAvailableModules(this.props.projects);
        let projects: IProject[] = this.fuseProjects(this.props.projects, this.state.queryString);
        let selectOptions: ISelectOption[] = [];

        projects = projects.filter((project) => {
            return this.state.selectedModules.length === 0 || this.state.selectedModules.includes(project.module);
        });
        for (let module of availableModules) {
            selectOptions.push({value: module, label: module});
        }
        return (
            <div>
                <div className={'d-flex d-flex-row'}>
                    <TextField placeholder={'Project name'} variant={'standard'}
                               color={'primary'}
                               onChange={(event => this.onSearchFieldChange(event.target.value.trim()))}
                               className={'ml-1 mt-1'}
                               autoFocus={true}/>
                    <Select
                        isMulti
                        inputValue={this.state.selectSearch}
                        onInputChange={this.onSelectSearchChange}
                        name="modules"
                        options={selectOptions}
                        onChange={(selectedOption: ValueType<ISelectOption, true>) => this.onSelectChange((selectedOption as ISelectOption[]))}
                        className="basic-multi-select ml-2 w-50"
                        classNamePrefix="select"
                        closeMenuOnSelect={this.state.selectSearch === ''}
                    />
                </div>
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

    onSelectSearchChange(value: string, actionMeta: InputActionMeta) {
        if (actionMeta.action === 'set-value')
            return;
        this.setState({
            ...this.state,
            selectSearch: value
        });
    }

    onSelectChange(modules: ISelectOption[]) {
        this.setState({
            ...this.state,
            selectedModules: modules.map(value => value.value)
        });
    }

    fuseProjects(originalProjectList: IProject[], queryString: string): IProject[] {
        if (queryString === '') {
            return originalProjectList;
        }
        const fuse = new Fuse(originalProjectList, projectsFuseOptions);
        const fuseResult = fuse.search(queryString);
        return fuseResult.map(fuseRes => {
            return fuseRes.item;
        });
    }
}

export default ProjectsRenderer;
