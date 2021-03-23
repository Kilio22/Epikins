import React from 'react';
import { TextField } from '@material-ui/core';
import Select, { ValueType } from 'react-select';
import { ISelectOption } from '../interfaces/projects/ISelectOption';
import { IProjectsToolbox } from '../interfaces/IProjectsToolbox';
import { Button, OverlayTrigger, Tooltip } from 'react-bootstrap';

const ProjectsToolbox: React.FunctionComponent<IProjectsToolbox> = ({
                                                                        allSelected,
                                                                        handleString,
                                                                        onForceUpdateClick,
                                                                        onSelectAllClick,
                                                                        onSelectChange,
                                                                        onSelectSearchChange,
                                                                        projects,
                                                                        selectOptions,
                                                                        selectSearch
                                                                    }) => {
    return (
        <div className={'d-flex d-flex-row align-items-center'}>
            <TextField placeholder={'Project name'} variant={'standard'}
                       color={'primary'}
                       onChange={(event => handleString(event.target.value.trim()))}
                       className={'ml-1'}
                       autoFocus={true}/>
            <Select
                isMulti
                inputValue={selectSearch}
                onInputChange={onSelectSearchChange}
                name={'modules'}
                options={selectOptions}
                onChange={(selectedOption: ValueType<ISelectOption, true>) => onSelectChange((selectedOption as ISelectOption[]))}
                className={'basic-multi-select ml-2 w-50'}
                classNamePrefix={'select'}
                closeMenuOnSelect={selectSearch === ''}/>
            {
                onSelectAllClick &&
                <div className={'d-flex align-items-center'}>
                    <input
                        className={'jobs-checkbox ml-2'}
                        type={'checkbox'}
                        checked={allSelected}
                        onChange={(event) => onSelectAllClick(event.target.checked, projects)}/>
                    <span className={'ml-1'}>Select all</span>
                </div>
            }
            <OverlayTrigger placement={'bottom'}
                            overlay={<Tooltip id={`tooltip-force-update`}>Force project list update</Tooltip>}>
                <Button onClick={onForceUpdateClick} className={'ml-2'}><i
                    className="fas fa-sync"/></Button>
            </OverlayTrigger>
        </div>
    );
};

export default ProjectsToolbox;
