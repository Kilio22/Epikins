package buildService

func canBuild(jobBuildData GroupBuildData) bool {
	return jobBuildData.mongoGroupData.RemainingBuilds > 0
}
