import React, { Component } from 'react';
import { IRouteProps } from '../../interfaces/IRoute';
import { appInitialContext } from '../../interfaces/IAppContext';
import { authServiceObj } from '../../services/AuthService';
import { userInitialState } from '../../interfaces/IUser';
import { IProject } from '../../interfaces/projects/IProject';
import EpikinsApiService from '../../services/EpikinsApiService';
import Loading from '../Loading';
import ProjectsRenderer from '../ProjectsRenderer';
import ProjectBuildLimitRenderer from './ProjectBuildLimitRenderer';
import ProjectForm from './ProjectForm';
import {
    IProjectsManagementState,
    projectsManagementInitialState
} from '../../interfaces/projectsManagement/IProjectsManagement';

class ProjectsManagement extends Component<IRouteProps, IProjectsManagementState> {
    static contextType = appInitialContext;
    context!: React.ContextType<typeof appInitialContext>;

    constructor(props: IRouteProps) {
        super(props);

        this.changeAllSelected = this.changeAllSelected.bind(this);
        this.onCheckboxClick = this.onCheckboxClick.bind(this);
        this.onSelectAllClick = this.onSelectAllClick.bind(this);
        this.getProjects = this.getProjects.bind(this);
        this.onProjectClick = this.onProjectClick.bind(this);
        this.updateProjects = this.updateProjects.bind(this);
        this.changeProjectsManagementStateByProperty = this.changeProjectsManagementStateByProperty.bind(this);

        this.state = projectsManagementInitialState;
    }

    async componentDidMount() {
        this.setState({...this.state, isLoading: true});
        await this.getProjects(false);
        this.setState({...this.state, isLoading: false});
    }

    render() {
        return (
            this.state.isLoading ?
                <Loading/>
                :
                <div>
                    {
                        this.state.clickedProject !== null &&
                        <ProjectForm onSelectAllClick={this.onSelectAllClick}
                                     project={this.state.clickedProject}
                                     projects={this.state.projects}
                                     changeProjectsManagementStateByProperty={this.changeProjectsManagementStateByProperty}
                                     getProjects={() => this.getProjects(false)}/>
                    }
                    <ProjectsRenderer allSelected={this.state.allSelected}
                                      changeAllSelected={this.changeAllSelected}
                                      onForceUpdateClick={() => this.updateProjects(true)}
                                      onSelectAllClick={this.onSelectAllClick}
                                      onCheckboxClick={this.onCheckboxClick}
                                      projects={this.state.projects}
                                      routeProps={this.props.routeProps}
                                      onProjectClick={this.onProjectClick}
                                      ProjectRenderer={ProjectBuildLimitRenderer}
                                      showSwitch={false}/>
                </div>
        );
    }

    async getProjects(forceUpdate: boolean) {
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            if (this.context.changeAppStateByProperty != null) {
                this.context.changeAppStateByProperty('user', userInitialState, false);
            }
            return;
        }

        const res: IProject[] | null = await EpikinsApiService.getProjects(forceUpdate, accessToken);
        if (res) {
            const sortedProjects: IProject[] = res.sort((a, b) => {
                return a.job.name.localeCompare(b.job.name);
            });
            sortedProjects.forEach((project) => {
                project.checked = false;
            });
            this.setState({
                ...this.state,
                projects: sortedProjects
            });
        } else {
            this.setState({
                ...this.state,
                projects: []
            });
            if (this.context.changeAppStateByProperty) {
                this.context.changeAppStateByProperty('errorMessage', 'Cannot fetch data, please try to reload the page.', true);
            }
        }
    }

    async updateProjects(forceUpdate: boolean) {
        this.setState({...this.state, isLoading: true});
        await this.getProjects(forceUpdate);
        this.setState({...this.state, isLoading: false});
    }

    onSelectAllClick(checked: boolean, projects: IProject[]) {
        let allProjects = this.state.projects;

        allProjects.forEach((project) => {
            project.checked = false;
        });
        projects.forEach((project) => {
            project.checked = checked;
        });
        this.setState({
            ...this.state,
            allSelected: checked,
            projects: [ ...this.state.projects ]
        });
    }

    onCheckboxClick(checkedProject: IProject) {
        checkedProject.checked = !checkedProject.checked;
        this.setState({
            ...this.state,
            allSelected: this.state.projects.every((project) => project.checked)
        });
    }

    changeAllSelected(value: boolean) {
        this.setState({
            ...this.state,
            allSelected: value
        });
    }

    onProjectClick(project: IProject) {
        if (project.checked) {
            this.setState({
                ...this.state,
                clickedProject: project
            });
        } else {
            this.setState({
                ...this.state,
                clickedProject: project,
                projects: this.state.projects.map((project) => {
                    return {
                        ...project,
                        checked: false
                    };
                })
            });
        }
    }

    changeProjectsManagementStateByProperty(property: keyof IProjectsManagementState, value: any) {
        this.setState({
            ...this.state,
            [property]: value
        });
    }

}

export default ProjectsManagement;
