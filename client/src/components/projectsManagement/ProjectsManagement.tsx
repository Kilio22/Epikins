import React, { Component } from 'react';
import { IRouteProps, routePrefix } from '../../interfaces/IRoute';
import { appInitialContext } from '../../interfaces/IAppContext';
import { authServiceObj } from '../../services/AuthService';
import { userInitialState } from '../../interfaces/IUser';
import { IProject } from '../../interfaces/projects/IProject';
import EpikinsApiService from '../../services/EpikinsApiService';
import Loading from '../Loading';
import ProjectsRenderer from '../projects/ProjectsRenderer';
import ProjectBuildLimitRenderer from './ProjectBuildLimitRenderer';
import ProjectForm from './ProjectForm';
import {
    IProjectsManagementState,
    projectsManagementInitialState
} from '../../interfaces/projectsManagement/IProjectsManagement';

class ProjectsManagement extends Component<IRouteProps<{}>, IProjectsManagementState> {
    static contextType = appInitialContext;
    context!: React.ContextType<typeof appInitialContext>;

    constructor(props: IRouteProps<{}>) {
        super(props);

        this.getProjects = this.getProjects.bind(this);
        this.onProjectClick = this.onProjectClick.bind(this);
        this.changeProjectsManagementStateByProperty = this.changeProjectsManagementStateByProperty.bind(this);

        this.state = projectsManagementInitialState;
    }

    async componentDidMount() {
        this.setState({...this.state, isLoading: true});
        await this.getProjects();
        this.setState({...this.state, isLoading: false});
    }

    render() {
        return (
            this.state.isLoading ?
                <Loading/>
                :
                <div>
                    {
                        this.state.selectedProject !== null &&
                        <ProjectForm project={this.state.selectedProject}
                                     changeProjectsManagementStateByProperty={this.changeProjectsManagementStateByProperty}
                                     getProjects={this.getProjects}/>
                    }
                    <ProjectsRenderer projects={this.state.projects}
                                      routeProps={this.props.routeProps}
                                      onProjectClick={this.onProjectClick}
                                      ProjectRenderer={ProjectBuildLimitRenderer}
                                      showSwitch={false}/>
                </div>
        );
    }

    async getProjects() {
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            if (this.context.changeAppStateByProperty != null) {
                this.context.changeAppStateByProperty('user', userInitialState, false);
            }
            return;
        }

        const res: IProject[] | null = await EpikinsApiService.getProjects(accessToken);
        if (res) {
            const sortedProjects: IProject[] = res.sort((a, b) => {
                return a.job.name.localeCompare(b.job.name);
            });
            sortedProjects.forEach((project) => {
                project.epikinsProjectURL = routePrefix + 'projects/' + project.job.name;
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

    onProjectClick(project: IProject) {
        this.setState({
            ...this.state,
            selectedProject: project
        });
    }

    changeProjectsManagementStateByProperty(property: keyof IProjectsManagementState, value: any) {
        this.setState({
            ...this.state,
            [property]: value
        });
    }

}

export default ProjectsManagement;