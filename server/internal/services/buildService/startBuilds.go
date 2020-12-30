package buildService

import (
	"epikins-api/internal/services/utils"
	"errors"

	"epikins-api/pkg/libJenkins"

	"go.mongodb.org/mongo-driver/mongo"
)

func buildLoop(buildParams BuildParams, groupsBuildData []GroupBuildData, userLogs libJenkins.JenkinsCredentials) error {
	for _, jobName := range buildParams.JobsToBuild {
		for idx := range groupsBuildData {
			if groupsBuildData[idx].groupJob.Name != jobName {
				continue
			}
			shouldBreak, err := startBuild(&groupsBuildData[idx], buildParams, userLogs)
			if err != nil {
				return err
			}
			if shouldBreak {
				break
			}
		}
	}
	return nil
}

func startBuilds(buildParams BuildParams, jobs []libJenkins.Job, collection *mongo.Collection, userLogs libJenkins.JenkinsCredentials) error {
	projectData, err := utils.FetchProjectData(buildParams.Project, jobs, collection)
	if err != nil {
		return errors.New("cannot start builds: " + err.Error())
	}

	groupsBuildData, err := getGroupsBuildData(jobs, projectData, buildParams.Project, collection)
	if err != nil {
		return errors.New("cannot start builds: " + err.Error())
	}
	err = buildLoop(buildParams, groupsBuildData, userLogs)
	if err != nil {
		return err
	}
	return updateMongoWorkgroupsData(buildParams.Project, groupsBuildData, collection)
}
