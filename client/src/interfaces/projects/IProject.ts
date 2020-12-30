import { IJob } from '../IJob';

export interface IProject {
    buildLimit: number,
    epikinsProjectURL: string,
    job: IJob,
    module: string,
}
