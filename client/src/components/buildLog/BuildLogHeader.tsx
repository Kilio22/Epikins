import React from 'react';
import { NativeSelect, TextField } from '@material-ui/core';
import { Button } from 'react-bootstrap';
import { ChangeBuildLogStateByProperty } from '../../interfaces/buildLog/IBuildLogExportForm';

type UpdateBuildLog = (city: string, page: number, project: string, email: string) => void;

interface IBuildLogHeaderProps {
    changeBuildLogStateByProperty: ChangeBuildLogStateByProperty,
    cities: string[],
    currentPage: number,
    isLoading: boolean,
    projectString: string,
    updateBuildLog: UpdateBuildLog,
    selectedCity: string,
    starterString: string
}

const BuildLogHeader: React.FunctionComponent<IBuildLogHeaderProps> = ({
                                                                           changeBuildLogStateByProperty,
                                                                           cities,
                                                                           currentPage,
                                                                           isLoading,
                                                                           projectString,
                                                                           updateBuildLog,
                                                                           selectedCity,
                                                                           starterString
                                                                       }) => {
    return (
        <div className={'row ml-1'}>
            <TextField placeholder={'Starter email'}
                       variant={'standard'}
                       color={'primary'}
                       className={'col-md-2 mt-2 col-6'}
                       autoFocus={true}
                       onChange={(async (event) => {
                           changeBuildLogStateByProperty('projectString', event.target.value.trim());
                           updateBuildLog(selectedCity, currentPage, projectString, event.target.value.trim());
                       })}
            />
            <TextField placeholder={'Project name'}
                       variant={'standard'}
                       color={'primary'}
                       className={'col-md-2 col-6 mt-2 ml-2'}
                       onChange={(async (event) => {
                           changeBuildLogStateByProperty('starterString', event.target.value.trim());
                           updateBuildLog(selectedCity, currentPage, event.target.value.trim(), starterString);
                       })}
            />
            <NativeSelect variant={'standard'}
                          disabled={isLoading}
                          onChange={(async event => {
                              changeBuildLogStateByProperty('selectedCity', event.target.value);
                              updateBuildLog(event.target.value, currentPage, projectString, starterString);
                          })}
                          defaultValue={'Any'} className={'col-auto ml-2'}>
                <option>
                    {'Any'}
                </option>
                {
                    cities.map((value, index) => {
                        return (
                            <option key={index}>
                                {value}
                            </option>
                        );
                    })
                }
            </NativeSelect>
            <Button className={'col-auto ml-2'}
                    onClick={() => changeBuildLogStateByProperty('showExportForm', true)}>Extract</Button>
        </div>
    );
};

export default BuildLogHeader;
