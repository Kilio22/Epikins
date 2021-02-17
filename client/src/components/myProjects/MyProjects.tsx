import React from 'react';
import { IRouteProps } from '../../interfaces/IRoute';
import { appInitialContext } from '../../interfaces/IAppContext';
import { authServiceObj } from '../../services/AuthService';
import { userInitialState } from '../../interfaces/IUser';
import EpikinsApiService from '../../services/EpikinsApiService';
import { IStudentJob } from '../../interfaces/myProjects/IStudentJob';
import Loading from '../Loading';
import StudentProjectsRenderer from './StudentProjectsRenderer';
import StudentBuildForm from './StudentBuildForm';

export interface IMyProjectsState {
    isLoading: boolean,
    projects: IStudentJob[],
    showForm: boolean,
    selectedProject: IStudentJob | null
}

const myProjectsInitialState: IMyProjectsState = {
    isLoading: false,
    projects: [],
    showForm: false,
    selectedProject: null
};

class MyProjects extends React.Component<IRouteProps, IMyProjectsState> {
    static contextType = appInitialContext;
    context!: React.ContextType<typeof appInitialContext>;

    constructor(props: IRouteProps) {
        super(props);

        this.getStudentProjects = this.getStudentProjects.bind(this);
        this.changeMyProjectsStateByProperty = this.changeMyProjectsStateByProperty.bind(this);
        this.onStudentProjectClick = this.onStudentProjectClick.bind(this);
        this.startBuild = this.startBuild.bind(this);

        this.state = myProjectsInitialState;
    }

    async componentDidMount() {
        this.setState({
            ...this.state,
            isLoading: true
        });
        await this.getStudentProjects();
        this.setState({
            ...this.state,
            isLoading: false
        });
    }

    render() {
        return (
            this.state.isLoading ?
                <Loading/>
                :
                <div>
                    {
                        this.state.showForm && this.state.selectedProject &&
                        <StudentBuildForm changeMyProjectsStateByProperty={this.changeMyProjectsStateByProperty}
                                          selectedProject={this.state.selectedProject} startBuild={this.startBuild}
                                          getStudentProjects={this.getStudentProjects}/>
                    }
                    <StudentProjectsRenderer projects={this.state.projects}
                                             showSwitch={true} onStudentProjectClick={this.onStudentProjectClick}/>
                </div>
        );
    }

    async getStudentProjects() {
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            if (this.context.changeAppStateByProperty != null) {
                this.context.changeAppStateByProperty('user', userInitialState, false);
            }
            return;
        }

        const res: IStudentJob[] | null = await EpikinsApiService.getStudentJobs(accessToken);
        if (res) {
            let sortedProjects: IStudentJob[] = res.filter((project) => project.mongoWorkgroupData.remainingBuilds !== 0)
                                                   .sort((a, b) => {
                                                       return a.project.name.localeCompare(b.project.name);
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

    async startBuild() {
        if (!this.state.selectedProject) {
            return;
        }
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            if (this.context.changeAppStateByProperty != null) {
                this.context.changeAppStateByProperty('user', userInitialState, false);
            }
            return;
        }

        const selectedProject = this.state.selectedProject;
        const res: boolean = await EpikinsApiService.buildStudent(selectedProject?.city,
            selectedProject.mongoWorkgroupData.name,
            selectedProject.project.name,
            selectedProject.project.module,
            accessToken);
        if (!res) {
            if (this.context.changeAppStateByProperty) {
                this.context.changeAppStateByProperty('errorMessage', 'Cannot start build, please try to reload the page.', true);
            }
        }
        return res;
    }

    changeMyProjectsStateByProperty(key: keyof IMyProjectsState, value: any) {
        this.setState({
            ...this.state,
            [key]: value
        });
    }

    onStudentProjectClick(project: IStudentJob) {
        this.setState({
            ...this.state,
            showForm: true,
            selectedProject: project
        });
    }
}

export default MyProjects;
