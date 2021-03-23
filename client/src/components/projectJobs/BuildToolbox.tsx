import { Button, ButtonGroup, OverlayTrigger, Spinner, Tooltip } from 'react-bootstrap';
import BuildDropdown from './BuildDropdown';
import React from 'react';
import { OnBuildClick, OnButtonClick } from '../../interfaces/Functions';
import { NativeSelect } from '@material-ui/core';

interface IBuildToolboxProps {
    availableCities: string[],
    isBuilding: boolean,
    onBuildClick: OnBuildClick,
    onCitySelected: OnBuildClick,
    onForceUpdateClick: OnButtonClick,
    onGlobalBuildClick: OnBuildClick,
    selectedCity: string,
    selectedJobs: string[],
}

const BuildToolbox: React.FunctionComponent<IBuildToolboxProps> = ({
                                                                       availableCities,
                                                                       isBuilding,
                                                                       onBuildClick,
                                                                       onCitySelected,
                                                                       onForceUpdateClick,
                                                                       onGlobalBuildClick,
                                                                       selectedCity,
                                                                       selectedJobs
                                                                   }) => {
    availableCities = availableCities.sort((a, b) => {
        return a.localeCompare(b);
    });
    return (
        <div className={'build-toolbox'}>
            <ButtonGroup>
                <BuildDropdown title={'Build'}
                               disabled={selectedJobs.length === 0 || isBuilding}
                               onBuildClick={onBuildClick}/>
                <BuildDropdown title={'Global build'} disabled={isBuilding}
                               onBuildClick={onGlobalBuildClick}/>
            </ButtonGroup>
            <NativeSelect variant={'standard'}
                          disabled={isBuilding}
                          onChange={(event => {
                              onCitySelected(event.target.value);
                          })}
                          defaultValue={selectedCity}>
                {
                    availableCities.map((value, index) => {
                        return (
                            <option key={index}>
                                {value}
                            </option>
                        );
                    })
                }
            </NativeSelect>
            <OverlayTrigger placement={'bottom'}
                            overlay={<Tooltip id={`tooltip-force-update`}>Force workgroup list update</Tooltip>}>
                <Button onClick={onForceUpdateClick} className={'ml-2'}><i
                    className="fas fa-sync"/></Button>
            </OverlayTrigger>
            {
                isBuilding &&
                <span>
                    <Spinner animation={'border'} variant={'primary'} style={{width: '20px', height: '20px'}}/> starting builds...
                </span>
            }
        </div>
    );
};

export default BuildToolbox;
