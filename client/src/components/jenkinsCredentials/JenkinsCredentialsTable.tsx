import * as React from 'react';
import { Table } from 'react-bootstrap';
import JenkinsCredentialsTableBody from './JenkinsCredentialsTableBody';
import JenkinsCredentialsTableHeader from './JenkinsCredentialsTableHeader';
import { IJenkinsCredentialsTableProps } from '../../interfaces/IJenkinsCredentialsTable/IJenkinsCredentialsTable';

const JenkinsCredentialsTable: React.FunctionComponent<IJenkinsCredentialsTableProps> = ({
                                                                                             jenkinsCredentials,
                                                                                             onFirstDeleteClick
                                                                                         }) => {
    return (
        <Table bordered responsive={'md'} striped className={'mt-3'}>
            <JenkinsCredentialsTableHeader/>
            <JenkinsCredentialsTableBody jenkinsCredentials={jenkinsCredentials}
                                         onFirstDeleteClick={onFirstDeleteClick}/>
        </Table>
    );
};

export default JenkinsCredentialsTable;
