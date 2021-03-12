import React from 'react';
import { IRouteProps } from '../../interfaces/IRoute';
import Loading from '../Loading';
import BuildLogCard from './BuildLogCard';
import BuildLogFooter from './BuildLogFooter';
import { userInitialState } from '../../interfaces/IUser';
import { authServiceObj } from '../../services/AuthService';
import EpikinsApiService from '../../services/EpikinsApiService';
import { TextField } from '@material-ui/core';
import { IBuildLogState, LogInitialState } from '../../interfaces/buildLog/IBuildLogState';

class BuildLog extends React.Component<IRouteProps, IBuildLogState> {
    constructor(props: IRouteProps) {
        super(props);

        this.onPageClick = this.onPageClick.bind(this);
        this.updateLog = this.updateLog.bind(this);

        this.state = LogInitialState;
    }

    async componentDidMount() {
        await this.updateLog(this.state.currentPage, this.state.projectString, this.state.starterString);
    }

    render() {
        return (
            <div>
                <TextField placeholder={'Starter email'}
                           variant={'standard'}
                           color={'primary'}
                           className={'ml-1'}
                           autoFocus={true}
                           onChange={(async (event) => {
                               this.setState({
                                   ...this.state,
                                   projectString: event.target.value.trim()
                               });
                               await this.updateLog(this.state.currentPage, this.state.projectString, event.target.value.trim());
                           })}
                />
                <TextField placeholder={'Project name'}
                           variant={'standard'}
                           color={'primary'}
                           className={'ml-1'}
                           onChange={(async (event) => {
                               this.setState({
                                   ...this.state,
                                   starterString: event.target.value.trim()
                               });
                               await this.updateLog(this.state.currentPage, event.target.value.trim(), this.state.starterString);
                           })}
                />
                {
                    this.state.isLoading ?
                        <Loading/>
                        :
                        <div>
                            {
                                this.state.buildLogInfo && this.state.buildLogInfo.totalPage !== 0 ?
                                    <div className={'d-flex flex-column'}>
                                        {
                                            this.state.buildLogInfo.buildLogs.map(((value, index) => {
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

    async updateLog(page: number, projectString: string, starterString: string) {
        this.setState({
            ...this.state,
            isLoading: true,
            requestInProgressNb: this.state.requestInProgressNb + 1
        });
        await this.getLog(page, projectString, starterString);
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
        await this.getLog(selectedItem.selected + 1, this.state.projectString, this.state.starterString);
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

    async getLog(page: number, projectString: string, starterString: string) {
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            this.resetUser();
            return;
        }

        const res = await EpikinsApiService.getLog(page, projectString, starterString, accessToken);
        if (res) {
            this.setState({
                ...this.state,
                buildLogInfo: res
            });
        } else {
            this.setErrorMessage('Cannot fetch data, please try to reload the page.');
        }
    }
}

export default BuildLog;
