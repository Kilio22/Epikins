import { HandleType } from '../Functions';

export interface IJenkinsCredentialsTableProps {
    jenkinsCredentials: string[],
    onFirstDeleteClick: HandleType<string>
}
