import * as React from 'react';
import { Button } from 'react-bootstrap';
import { IJenkinsCredentialsToolboxProps } from '../../interfaces/IJenkinsCredentialsTable/IJenkinsCredentialsToolbox';

const JenkinsCredentialsToolbox: React.FunctionComponent<IJenkinsCredentialsToolboxProps> = ({
                                                                                                 onAddClick
                                                                                             }) => {
    return (
        <Button className={'ml-2'}
                onClick={onAddClick}>
            <span><i className={'far fa-plus-square'}/> Add</span>
        </Button>
    );
};

export default JenkinsCredentialsToolbox;
