import React, { Component } from 'react';
import 'bootstrap/dist/css/bootstrap.css';
import { IProjectsState, projectsInitialState } from '../../interfaces/projects/IProjects';
import EpikinsApiService from '../../services/EpikinsApiService';
import ProjectsRenderer from './ProjectsRenderer';
import { IRouteProps, routePrefix } from '../../interfaces/IRoute';
import { apiBaseURI } from '../../Config';
import { IJob } from '../../interfaces/IJob';
import { appInitialContext } from '../../interfaces/IAppContext';
import { authServiceObj } from '../../services/AuthService';
import Loading from '../Loading';
import { userInitialState } from '../../interfaces/IUser';

class Projects extends Component<IRouteProps<{}>, IProjectsState> {
    static contextType = appInitialContext;
    context!: React.ContextType<typeof appInitialContext>;

    constructor(props: IRouteProps<{}>) {
        super(props);

        this.state = projectsInitialState;
        this.getProjects = this.getProjects.bind(this);
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
                <ProjectsRenderer projects={this.state.projects} routeProps={this.props.routeProps}/>
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

        const res: IJob[] | null = await EpikinsApiService.getJobs(apiBaseURI + 'projects', accessToken);
        if (res) {
            const sortedJobs: IJob[] = res.sort((a, b) => {
                return a.name.localeCompare(b.name);
            });
            sortedJobs.forEach((job) => {
                job.epikinsJobURL = routePrefix + 'projects/' + job.name;
            });
            this.setState({
                ...this.state,
                projects: sortedJobs
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
}

export default Projects;
