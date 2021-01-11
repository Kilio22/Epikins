import React from 'react';
import EpikinsApiService from '../../services/EpikinsApiService';
import { apiBaseURI } from '../../Config';
import {
    IProjectJobsMatchParams,
    IProjectJobsState,
    projectJobsInitialState
} from '../../interfaces/jobs/IProjectJobs';
import ProjectJobsRenderer from './ProjectJobsRenderer';
import { IWorkgroupsData } from '../../interfaces/IWorkgroupsData';
import { IRouteProps } from '../../interfaces/IRoute';
import { appInitialContext } from '../../interfaces/IAppContext';
import { authServiceObj } from '../../services/AuthService';
import Loading from '../Loading';
import { userInitialState } from '../../interfaces/IUser';

class ProjectJobs extends React.Component<IRouteProps<IProjectJobsMatchParams>, IProjectJobsState> {
    static contextType = appInitialContext;
    context!: React.ContextType<typeof appInitialContext>;
    private mounted = false;

    constructor(props: IRouteProps<IProjectJobsMatchParams>) {
        super(props);

        this.state = projectJobsInitialState;

        this.getJobsByProject = this.getJobsByProject.bind(this);
        this.onCheckboxChange = this.onCheckboxChange.bind(this);
        this.onBuildClick = this.onBuildClick.bind(this);
        this.onGlobalBuildClick = this.onGlobalBuildClick.bind(this);
    }

    async componentDidMount() {
        this.mounted = true;
        this.setState({...this.state, isLoading: true});
        await this.getJobsByProject(true);
        this.setState({...this.state, isLoading: false});
    }

    componentWillUnmount() {
        this.mounted = false;
    }

    render() {
        return (
            this.state.isLoading ?
                <Loading/>
                :
                <ProjectJobsRenderer workgroupsData={this.state.workgroupsData}
                                     isBuilding={this.state.isBuilding}
                                     selectedJobs={this.state.selectedJobs}
                                     onCheckboxChange={this.onCheckboxChange}
                                     onBuildClick={this.onBuildClick}
                                     onGlobalBuildClick={this.onGlobalBuildClick}
                                     routeProps={this.props.routeProps}/>
        );
    }

    async getJobsByProject(shouldCallback: boolean) {
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            if (this.context.changeAppStateByProperty != null) {
                this.context.changeAppStateByProperty('user', userInitialState, false);
            }
            return;
        }

        const res: IWorkgroupsData[] | null = await EpikinsApiService.getWorkgroupsData(
            apiBaseURI + 'projects/' + this.props.routeProps.match.params.project + '/REN', accessToken);

        if (!this.mounted) {
            return;
        }
        if (res) {
            const sortedWorkgroupsData: IWorkgroupsData[] = res.sort((a, b) => {
                return a.mongoWorkgroupData.name.localeCompare(b.mongoWorkgroupData.name);
            });
            if (shouldCallback) {
                this.setState({
                    ...this.state,
                    workgroupsData: sortedWorkgroupsData
                }, () => setTimeout(() => {
                    if (this.mounted) {
                        this.getJobsByProject(true);
                    }
                }, 5000));
            } else {
                this.setState({
                    ...this.state,
                    workgroupsData: sortedWorkgroupsData
                });
            }
        } else {
            this.setState({
                ...this.state,
                workgroupsData: []
            });
            if (this.context.changeAppStateByProperty) {
                this.context.changeAppStateByProperty('errorMessage',
                    'Cannot fetch data, please try to reload the page.', true);
            }
        }
    }

    onCheckboxChange(checked: boolean, groupJob: IWorkgroupsData) {
        if (this.state.selectedJobs.includes(groupJob.mongoWorkgroupData.name)) {
            this.setState({
                ...this.state,
                selectedJobs: this.state.selectedJobs.filter((value) => {
                    return value !== groupJob.mongoWorkgroupData.name;
                })
            });
        } else {
            this.setState({
                ...this.state,
                selectedJobs: this.state.selectedJobs.concat(groupJob.mongoWorkgroupData.name)
            });
        }
    }

    async handleBuildResponse(res: boolean) {
        if (res) {
            await this.getJobsByProject(false);
        } else {
            if (this.context.changeAppStateByProperty) {
                this.context.changeAppStateByProperty('errorMessage',
                    'Cannot build jobs, see console for more infos', true);
            }
        }
        this.setState({
            ...this.state,
            isBuilding: false
        });
    }

    async onBuildClick(visibility: string) {
        const selectedJobs: string[] = this.state.selectedJobs;
        this.setState({
            ...this.state,
            isBuilding: true,
            selectedJobs: []
        });

        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            if (this.context.changeAppStateByProperty != null) {
                this.context.changeAppStateByProperty('user', userInitialState, false);
            }
            return;
        }

        await this.handleBuildResponse(await EpikinsApiService.buildJobs(
            selectedJobs, this.props.routeProps.match.params.project, visibility, this.context.fuMode, accessToken
        ));
    }

    async onGlobalBuildClick(visibility: string) {
        this.setState({
            ...this.state,
            isBuilding: true,
            selectedJobs: []
        });

        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            if (this.context.changeAppStateByProperty != null) {
                this.context.changeAppStateByProperty('user', userInitialState, false);
            }
            return;
        }

        await this.handleBuildResponse(await EpikinsApiService.globalBuild(this.props.routeProps.match.params.project, visibility, accessToken));
    }
}

export default ProjectJobs;
