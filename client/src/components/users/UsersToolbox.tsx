import * as React from 'react';
import { Button, ButtonGroup } from 'react-bootstrap';
import { IUsersToolboxProps } from '../../interfaces/users/IUsersToolbox';

const UsersToolbox: React.FunctionComponent<IUsersToolboxProps> = ({
                                                                       isEditing,
                                                                       isSaving,
                                                                       onSaveClick,
                                                                       onCancelClick,
                                                                       onEditClick,
                                                                       onAddClick
                                                                   }) => {
    return (
        isEditing ?
            <ButtonGroup>
                <Button disabled={isSaving}
                        onClick={onSaveClick}>
                    <span><i className={'far fa-check-circle'}/> Save</span>
                </Button>
                <Button className={'ml-2'}
                        disabled={isSaving}
                        onClick={onCancelClick}>
                    <span><i className={'far fa-times-circle'}/> Cancel</span>
                </Button>
            </ButtonGroup>
            :
            <ButtonGroup>
                <Button className={'ml-2'}
                        onClick={onEditClick}>
                    <span><i className={'far fa-edit'}/> Edit</span>
                </Button>
                <Button className={'ml-2'}
                        onClick={onAddClick}>
                    <span><i className={'far fa-plus-square'}/> Add</span>
                </Button>
            </ButtonGroup>
    );
};

export default UsersToolbox;
