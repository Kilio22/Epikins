import React from 'react';
import EpikinsApiService from '../../services/EpikinsApiService';
import { apiBaseURI } from '../../Config';
import { IProjectJobsState, IProjectLocationState, projectJobsInitialState } from '../../interfaces/jobs/IProjectJobs';
import ProjectJobsRenderer from './ProjectJobsRenderer';
import { IWorkgroupsData } from '../../interfaces/IWorkgroupsData';
import { IRouteProps } from '../../interfaces/IRoute';
import { appInitialContext } from '../../interfaces/IAppContext';
import { authServiceObj } from '../../services/AuthService';
import Loading from '../Loading';
import { userInitialState } from '../../interfaces/IUser';
import { IProject } from '../../interfaces/projects/IProject';

class ProjectJobs extends React.Component<IRouteProps<{}, IProjectLocationState>, IProjectJobsState> {
    static contextType = appInitialContext;
    context!: React.ContextType<typeof appInitialContext>;
    private mounted = false;

    constructor(props: IRouteProps<{}, IProjectLocationState>) {
        super(props);

        if (props.routeProps.location?.state?.project) {
            this.state = {
                ...projectJobsInitialState,
                project: props.routeProps.location.state.project,
                selectedCity: props.routeProps.location.state.project.cities[0]
            };
        } else {
            this.state = projectJobsInitialState;
        }

        this.getProjectInformation = this.getProjectInformation.bind(this);
        this.getJobsByProject = this.getJobsByProject.bind(this);
        this.onCitySelected = this.onCitySelected.bind(this);
        this.onCheckboxChange = this.onCheckboxChange.bind(this);
        this.onBuildClick = this.onBuildClick.bind(this);
        this.onGlobalBuildClick = this.onGlobalBuildClick.bind(this);
    }

    async componentDidMount() {
        this.mounted = true;
        this.setState({...this.state, isLoading: true});
        if (!this.state.project) {
            await this.getProjectInformation();
        }
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
                <ProjectJobsRenderer selectedCity={this.state.selectedCity} onCitySelected={this.onCitySelected}
                                     availableCities={this.state.project?.cities}
                                     workgroupsData={this.state.workgroupsData}
                                     isBuilding={this.state.isBuilding}
                                     selectedJobs={this.state.selectedJobs}
                                     onCheckboxChange={this.onCheckboxChange}
                                     onBuildClick={this.onBuildClick}
                                     onGlobalBuildClick={this.onGlobalBuildClick}
                                     routeProps={this.props.routeProps}/>
        );
    }

    async onCitySelected(city: string) {
        this.setState({
            ...this.state,
            selectedCity: city,
            isLoading: true
        });
        await this.getJobsByProject(true);
        this.setState({...this.state, isLoading: false});
    }

    async getProjectInformation() {
        if (this.state.project == null) {
            return;
        }
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            if (this.context.changeAppStateByProperty != null) {
                this.context.changeAppStateByProperty('user', userInitialState, false);
            }
            return;
        }

        const res: IProject | null = await EpikinsApiService.getProjectInformation(this.state.project.job.name, this.state.project.module, accessToken);
        if (!this.mounted) {
            return;
        }
        if (res) {
            this.setState({
                ...this.state,
                project: res,
                selectedCity: res.cities[0]
            });
        } else {
            if (this.context.changeAppStateByProperty) {
                this.context.changeAppStateByProperty('errorMessage',
                    'Cannot fetch data, please try to reload the page.', true);
            }
        }
    }

    async getJobsByProject(shouldCallback: boolean) {
        if (!this.state.project) {
            return;
        }
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            if (this.context.changeAppStateByProperty != null) {
                this.context.changeAppStateByProperty('user', userInitialState, false);
            }
            return;
        }

        const res: IWorkgroupsData[] | null = await EpikinsApiService.getWorkgroupsData(
            apiBaseURI + '/projects/' + this.state.project.module + '/' + this.state.project.job.name + '/' + this.state.selectedCity, accessToken);

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
        if (!this.state.project) {
            return;
        }
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
            selectedJobs, this.state.project.job.name, visibility, this.state.selectedCity, this.state.project.module, accessToken
        ));
    }

    async onGlobalBuildClick(visibility: string) {
        if (!this.state.project) {
            return;
        }
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

        await this.handleBuildResponse(await EpikinsApiService.globalBuild(this.state.project.job.name, visibility, this.state.selectedCity, this.state.project.module, accessToken));
    }
}

export default ProjectJobs;
