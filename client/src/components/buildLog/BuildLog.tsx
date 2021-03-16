import React from 'react';
import { IRouteProps } from '../../interfaces/IRoute';
import Loading from '../Loading';
import BuildLogCard from './BuildLogCard';
import BuildLogFooter from './BuildLogFooter';
import { userInitialState } from '../../interfaces/IUser';
import { authServiceObj } from '../../services/AuthService';
import EpikinsApiService from '../../services/EpikinsApiService';
import { IBuildLogState, LogInitialState } from '../../interfaces/buildLog/IBuildLogState';
import BuildLogExportForm from './BuildLogExportForm';
import BuildLogHeader from './BuildLogHeader';

class BuildLog extends React.Component<IRouteProps, IBuildLogState> {
    constructor(props: IRouteProps) {
        super(props);

        this.changeBuildLogStateByProperty = this.changeBuildLogStateByProperty.bind(this);
        this.getCities = this.getCities.bind(this);
        this.getLog = this.getLog.bind(this);
        this.onPageClick = this.onPageClick.bind(this);
        this.resetUser = this.resetUser.bind(this);
        this.setErrorMessage = this.setErrorMessage.bind(this);
        this.updateBuildLog = this.updateBuildLog.bind(this);

        this.state = LogInitialState;
    }

    async componentDidMount() {
        this.setState({
            ...this.state,
            isLoading: true,
            requestInProgressNb: this.state.requestInProgressNb + 1
        });
        await this.getLog(this.state.selectedCity, this.state.currentPage, this.state.projectString, this.state.starterString);
        await this.getCities();
        this.setState({
            ...this.state,
            isLoading: this.state.requestInProgressNb === 0,
            requestInProgressNb: this.state.requestInProgressNb - 1
        });
    }

    render() {
        return (
            <div className={'h-100 d-flex flex-column'}>
                <BuildLogHeader cities={this.state.cities}
                                changeBuildLogStateByProperty={this.changeBuildLogStateByProperty}
                                currentPage={this.state.currentPage}
                                isLoading={this.state.isLoading}
                                projectString={this.state.projectString}
                                selectedCity={this.state.selectedCity}
                                starterString={this.state.starterString}
                                updateBuildLog={this.updateBuildLog}/>
                {
                    this.state.showExportForm &&
                    <BuildLogExportForm changeBuildLogStateByProperty={this.changeBuildLogStateByProperty}
                                        cities={this.state.cities}/>
                }
                {
                    this.state.isLoading ?
                        <Loading/>
                        :
                        <div>
                            {
                                this.state.buildLogInfo && this.state.buildLogInfo.totalPage !== 0 ?
                                    <div className={'d-flex flex-column'}>
                                        {
                                            this.state.buildLogInfo.buildLog.map(((value, index) => {
                                                return (
                                                    <BuildLogCard buildLog={value} key={index}/>
                                                );
                                            }))
                                        }
                                        <BuildLogFooter currentPage={this.state.currentPage}
                                                        totalPage={this.state.buildLogInfo.totalPage}
                                                        onPageClick={this.onPageClick}/>
                                    </div>
                                    :
                                    <h2 className={'text-center'}>No build log to display</h2>
                            }
                        </div>
                }
            </div>
        );
    }

    async getCities() {
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            this.resetUser();
            return;
        }

        const res = await EpikinsApiService.getCities(accessToken);
        if (res) {
            this.setState({
                ...this.state,
                cities: res.sort(((a, b) => a.localeCompare(b)))
            });
        } else {
            this.setErrorMessage('Cannot fetch cities list, please try to reload the page.');
        }
    }

    async updateBuildLog(city: string, page: number, projectString: string, starterString: string) {
        this.setState({
            ...this.state,
            isLoading: true,
            requestInProgressNb: this.state.requestInProgressNb + 1
        });
        await this.getLog(city, page, projectString, starterString);
        this.setState({
            ...this.state,
            isLoading: this.state.requestInProgressNb === 0,
            requestInProgressNb: this.state.requestInProgressNb - 1
        });
    }

    async onPageClick(selectedItem: { selected: number }) {
        this.setState({
            ...this.state,
            isLoading: true,
            currentPage: selectedItem.selected + 1
        });
        await this.getLog(this.state.selectedCity, selectedItem.selected + 1, this.state.projectString, this.state.starterString);
        this.setState({
            ...this.state,
            isLoading: false
        });
    }

    setErrorMessage(message: string) {
        if (this.context.changeAppStateByProperty) {
            this.context.changeAppStateByProperty('errorMessage', message, true);
        }
    }

    resetUser() {
        if (this.context.changeAppStateByProperty != null) {
            this.context.changeAppStateByProperty('user', userInitialState, false);
        }
    }

    async getLog(city: string, page: number, projectString: string, starterString: string) {
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            this.resetUser();
            return;
        }

        city = city.localeCompare('Any') === 0 ? '' : city;
        const res = await EpikinsApiService.getBuildLog(city, page, projectString, starterString, accessToken);
        if (res) {
            this.setState({
                ...this.state,
                buildLogInfo: res
            });
        } else {
            this.setErrorMessage('Cannot fetch data, please try to reload the page.');
        }
    }

    changeBuildLogStateByProperty(key: keyof IBuildLogState, value: any) {
        this.setState({
            ...this.state,
            [key]: value
        });
    }
}

export default BuildLog;
