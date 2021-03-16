import React from 'react';
import { Button, FormControl, InputGroup, Modal, Spinner } from 'react-bootstrap';
import DatePicker from 'react-datepicker';
import { saveAs } from 'file-saver';

import 'react-datepicker/dist/react-datepicker.css';

import { authServiceObj } from '../../services/AuthService';
import EpikinsApiService from '../../services/EpikinsApiService';
import { userInitialState } from '../../interfaces/IUser';
import {
    buildLogExportFormInitialState,
    IBuildLogExportFormProps,
    IBuildLogExportFormState
} from '../../interfaces/buildLog/IBuildLogExportForm';
import { NativeSelect } from '@material-ui/core';


class BuildLogExportForm extends React.Component<IBuildLogExportFormProps, IBuildLogExportFormState> {
    constructor(props: IBuildLogExportFormProps) {
        super(props);

        this.onExportClick = this.onExportClick.bind(this);
        this.resetUser = this.resetUser.bind(this);
        this.setErrorMessage = this.setErrorMessage.bind(this);

        this.state = {
            ...buildLogExportFormInitialState,
            startDate: new Date(),
            endDate: new Date()
        };
    }

    render() {
        return (
            <Modal show
                   onHide={() => this.props.changeBuildLogStateByProperty('showExportForm', false)}
                   size={'lg'}
                   centered>
                <Modal.Header closeButton>
                    <Modal.Title>
                        Export build log
                    </Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <h6>Start date</h6>
                    <DatePicker selected={this.state.startDate}
                                disabled={this.state.isLoading}
                                showTimeSelect
                                dateFormat={'MMMM d, yyyy, HH:mm OOOO'}
                                timeFormat={'HH:mm'}
                                onChange={(date) => {
                                    if (date instanceof Date) {
                                        this.setState({
                                            ...this.state,
                                            startDate: date
                                        });
                                    }
                                }}/>
                    <h6 className={'mt-2'}>End date</h6>
                    <DatePicker selected={this.state.endDate}
                                disabled={this.state.isLoading}
                                showTimeSelect
                                dateFormat={'MMMM d, yyyy, HH:mm OOOO'}
                                timeFormat={'HH:mm'}
                                onChange={(date) => {
                                    if (date instanceof Date) {
                                        this.setState({
                                            ...this.state,
                                            endDate: date
                                        });
                                    }
                                }}/>
                    <h6 className={'mt-2'}>City</h6>
                    <NativeSelect variant={'standard'}
                                  disabled={this.state.isLoading}
                                  onChange={(event => this.setState({...this.state, selectedCity: event.target.value}))}
                                  defaultValue={'Any'}>
                        <option>
                            {'Any'}
                        </option>
                        {
                            this.props.cities.map((value, index) => {
                                return (
                                    <option key={index}>
                                        {value}
                                    </option>
                                );
                            })
                        }
                    </NativeSelect>
                    <h6 className={'mt-2'}>Project name</h6>
                    <InputGroup>
                        <FormControl disabled={this.state.isLoading} onChange={((event) => this.setState({
                            ...this.state,
                            projectValue: event.target.value.trim()
                        }))}
                                     placeholder={'Toto'}
                                     aria-label={'Toto'}
                        />
                    </InputGroup>
                    <div className={'d-flex justify-content-center mt-3'}>
                        <Button variant={'primary'} disabled={this.state.isLoading} onClick={this.onExportClick}>
                            {
                                this.state.isLoading ?
                                    <Spinner animation={'border'}/>
                                    :
                                    'Export'
                            }
                        </Button>
                    </div>
                </Modal.Body>
            </Modal>
        );
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

    async onExportClick() {
        const accessToken: string = await authServiceObj.getToken();
        if (accessToken === '') {
            this.resetUser();
            return;
        }

        this.setState({
            ...this.state,
            isLoading: true
        });
        const startTimestamp = parseInt((this.state.startDate.getTime() / 1000).toFixed(0));
        const endTimestamp = parseInt((this.state.endDate.getTime() / 1000).toFixed(0));
        const selectedCity = this.state.selectedCity.localeCompare('Any') === 0 ? '' : this.state.selectedCity;
        const res = await EpikinsApiService.exportBuildLog(selectedCity, startTimestamp, endTimestamp, this.state.projectValue, accessToken);
        if (res) {
            saveAs(res, 'epikins_export_' + startTimestamp + '_' + endTimestamp + '.csv');
        } else {
            this.setErrorMessage('Cannot fetch data, please try to reload the page.');
        }
        this.props.changeBuildLogStateByProperty('showExportForm', false);
    }
}

export default BuildLogExportForm;
