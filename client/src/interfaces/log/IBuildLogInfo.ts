interface IBuildLog {
    module: string,
    project: string,
    starter: string,
    target: string,
    time: number
}

export interface IBuildLogInfo {
    buildLogs: IBuildLog[],
    totalPage: number
}
