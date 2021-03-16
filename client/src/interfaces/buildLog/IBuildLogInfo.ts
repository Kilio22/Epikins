export interface IBuildLog {
    module: string,
    project: string,
    starter: string,
    target: string,
    time: number
}

export interface IBuildLogInfo {
    buildLog: IBuildLog[],
    totalPage: number
}
