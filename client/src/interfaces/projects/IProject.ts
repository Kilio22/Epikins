import { IJob } from '../IJob';

export interface IProject {
    buildLimit: number,
    cities: string[],
    epikinsProjectURL: string,
    job: IJob,
    module: string,
}
