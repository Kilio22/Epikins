import React from 'react';
import { Col, Row } from 'react-bootstrap';
import Fuse from 'fuse.js';
import {
    IProjectsRendererProps,
    IProjectsRendererState,
    projectsRendererInitialState
} from '../interfaces/IProjectsRenderer';
import { appInitialContext } from '../interfaces/IAppContext';
import { IProject } from '../interfaces/projects/IProject';
import { InputActionMeta } from 'react-select';
import { ISelectOption } from '../interfaces/projects/ISelectOption';
import ProjectsToolbox from './ProjectsToolbox';

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
        if (this.props.changeAllSelected != null) {
            const isAllSelected = projects.every((project) => project.checked);
            if (isAllSelected !== this.props.allSelected) {
                this.props.changeAllSelected(isAllSelected);
            }
        }
        return (
            <div>
                <ProjectsToolbox
                    allSelected={this.props.allSelected}
                    handleString={this.onSearchFieldChange}
                    onSelectAllClick={this.props.onSelectAllClick}
                    onSelectChange={this.onSelectChange}
                    onSelectSearchChange={this.onSelectSearchChange}
                    projects={projects}
                    selectOptions={selectOptions}
                    selectSearch={this.state.selectSearch}/>
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
                                                <this.props.ProjectRenderer onCheckboxClick={this.props.onCheckboxClick}
                                                                            onProjectClick={this.props.onProjectClick}
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
