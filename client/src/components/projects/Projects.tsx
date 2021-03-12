import React, { Component } from 'react';
import 'bootstrap/dist/css/bootstrap.css';
import { IProjectsState, projectsInitialState } from '../../interfaces/projects/IProjects';
import ProjectsRenderer from '../ProjectsRenderer';
import { IRouteProps, routePrefix } from '../../interfaces/IRoute';
import { appInitialContext } from '../../interfaces/IAppContext';
import Loading from '../Loading';
import ProjectRenderer from './ProjectRenderer';
import { authServiceObj } from '../../services/AuthService';
import { userInitialState } from '../../interfaces/IUser';
import EpikinsApiService from '../../services/EpikinsApiService';
import { IProject } from '../../interfaces/projects/IProject';

class Projects extends Component<IRouteProps, IProjectsState> {
    static contextType = appInitialContext;
    context!: React.ContextType<typeof appInitialContext>;

    constructor(props: IRouteProps) {
        super(props);

        this.getProjects = this.getProjects.bind(this);
        this.onProjectClick = this.onProjectClick.bind(this);

        this.state = projectsInitialState;
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
                <ProjectsRenderer
                    allSelected={false}
                    changeAllSelected={null}
                    onCheckboxClick={null}
                    onSelectAllClick={null}
                    projects={this.state.projects}
                    routeProps={this.props.routeProps}
                    onProjectClick={this.onProjectClick}
                    ProjectRenderer={ProjectRenderer}
                    showSwitch={true}/>
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
            let sortedProjects: IProject[] = res.sort((a, b) => {
                return a.job.name.localeCompare(b.job.name);
            });
            sortedProjects = sortedProjects.filter((project) => {
                return project.cities.length !== 0;
            });
            sortedProjects.forEach((project) => {
                project.epikinsProjectURL = routePrefix + 'projects/' + project.module + '/' + project.job.name;
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

    onProjectClick(project: IProject) {
        this.props.routeProps.history.push(project.epikinsProjectURL, {
            project
        });
    }
}

export default Projects;
