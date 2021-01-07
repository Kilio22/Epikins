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

export interface IMongoWorkgroupData {
    name: string,
    remainingBuilds: number,
    url: string
}

export interface IWorkgroupsData {
    jobInfos: IJobInfos,
    mongoWorkgroupData: IMongoWorkgroupData
}
