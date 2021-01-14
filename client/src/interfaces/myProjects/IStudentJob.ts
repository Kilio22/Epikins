import { IMongoWorkgroupData } from '../IWorkgroupsData';

interface IStudentProject {
    module: string,
    name: string,
    buildLimit: string
}

export interface IStudentJob {
    city: string,
    mongoWorkgroupData: IMongoWorkgroupData,
    project: IStudentProject
}
