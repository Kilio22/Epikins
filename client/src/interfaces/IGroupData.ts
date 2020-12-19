import {IJob} from "./IJob";

export interface IBuildInfos {
    building: boolean
}

export interface IBuild {
    number: number,
    url: string,
    buildInfos: IBuildInfos
}

export interface IJobInfos {
    inQueue: boolean,
    lastBuild: IBuild
}

export interface IGroupJob {
    job: IJob,
    jobInfos: IJobInfos
}

export interface IMongoGroupData {
    name: string,
    remainingBuilds: number
}

export interface IGroupData {
    groupJob: IGroupJob,
    mongoGroupData: IMongoGroupData
}