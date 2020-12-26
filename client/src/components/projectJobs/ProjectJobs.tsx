import React from 'react';
import EpikinsApiService from '../../services/EpikinsApiService';
import { apiBaseURI } from '../../Config';
import {
    IProjectJobsMatchParams,
    IProjectJobsState,
    projectJobsInitialState
} from '../../interfaces/jobs/IProjectJobs';
import ProjectJobsRenderer from './ProjectJobsRenderer';
import { IGroupData } from '../../interfaces/IGroupData';
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
                <ProjectJobsRenderer groupsData={this.state.groupsData}
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

        const res: IGroupData[] | null = await EpikinsApiService.getGroupsData(
            apiBaseURI + 'projects/' + this.props.routeProps.match.params.project, accessToken);
        if (res) {
            const newGroupsOfJobs: IGroupData[] = res.sort((a, b) => {
                return a.groupJob.job.name.localeCompare(b.groupJob.job.name);
            });
            if (shouldCallback) {
                this.setState({
                    ...this.state,
                    groupsData: newGroupsOfJobs
                }, () => setTimeout(() => {
                    if (this.mounted) {
                        this.getJobsByProject(true);
                    }
                }, 5000));
            } else {
                this.setState({
                    ...this.state,
                    groupsData: newGroupsOfJobs
                });
            }
        } else {
            this.setState({
                ...this.state,
                groupsData: []
            });
            if (this.context.changeAppStateByProperty) {
                this.context.changeAppStateByProperty('errorMessage',
                    'Cannot fetch data, please try to reload the page.', true);
            }
        }
    }

    onCheckboxChange(checked: boolean, groupJob: IGroupData) {
        if (this.state.selectedJobs.includes(groupJob.groupJob.job.name)) {
            this.setState({
                ...this.state,
                selectedJobs: this.state.selectedJobs.filter((value) => {
                    return value !== groupJob.groupJob.job.name;
                })
            });
        } else {
            this.setState({
                ...this.state,
                selectedJobs: this.state.selectedJobs.concat(groupJob.groupJob.job.name)
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
