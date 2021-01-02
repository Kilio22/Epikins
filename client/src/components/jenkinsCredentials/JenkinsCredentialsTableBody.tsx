import * as React from 'react';
import { IJenkinsCredentialsTableBodyProps } from '../../interfaces/jenkinsCredentials/IJenkinsCredentialsTableBody';

const JenkinsCredentialsTableBody: React.FunctionComponent<IJenkinsCredentialsTableBodyProps> = ({
                                                                                                     jenkinsCredentials,
                                                                                                     onFirstDeleteClick
                                                                                                 }) => {
    const sortedCredentials = jenkinsCredentials.sort(((a, b) => a.localeCompare(b)));
    return (
        <tbody>
        {
            sortedCredentials.map((credentials, credentialsIdx) => {
                return (
                    <tr key={credentialsIdx}>
                        <td>{credentials}</td>
                        {
                            <td>
                                <i className={'far fa-trash-alt trash'}
                                   onClick={() => onFirstDeleteClick(credentials)}/>
                            </td>
                        }
                    </tr>
                );
            })
        }
        </tbody>
    );
};

export default JenkinsCredentialsTableBody;
