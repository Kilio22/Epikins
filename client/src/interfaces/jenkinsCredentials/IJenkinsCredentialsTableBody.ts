import { HandleType } from '../Functions';

export interface IJenkinsCredentialsTableBodyProps {
    jenkinsCredentials: string[],
    onFirstDeleteClick: HandleType<string>
}
