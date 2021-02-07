import { IJob } from '../IJob';

export interface IProject {
    buildLimit: number,
    checked: boolean,
    cities: string[],
    epikinsProjectURL: string,
    job: IJob,
    module: string,
}
