import React from 'react';
import { ButtonGroup, Dropdown, DropdownButton } from 'react-bootstrap';
import { OnBuildClick } from '../../interfaces/Functions';

interface IBuildDropdownProps {
    title: string,
    disabled: boolean,
    onBuildClick: OnBuildClick
}

const BuildDropdown: React.FunctionComponent<IBuildDropdownProps> = ({title, disabled, onBuildClick}) => {
    return (
        <DropdownButton as={ButtonGroup}
                        title={<span><i className={'fas fa-play fa-color-play mr-2'}/>{title}</span>}
                        className={'mr-2'}
                        disabled={disabled}
                        size={'sm'}>
            <Dropdown.Item onClick={() => onBuildClick('Private')}>Private</Dropdown.Item>
            <Dropdown.Item onClick={(() => onBuildClick('Public'))}>Public</Dropdown.Item>
        </DropdownButton>
    );
};

export default BuildDropdown;
