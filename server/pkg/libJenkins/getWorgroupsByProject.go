package libJenkins

import (
	"errors"
)

type Workgroup struct {
	Job      Job      `json:"job"`
	JobInfos JobInfos `json:"jobInfos"`
}

func getWorkgroups(jobsList []Job, userLogs JenkinsCredentials) ([]Workgroup, error) {
	var groupsOfJobs []Workgroup
	for _, job := range jobsList {
		jobInfos, err := getJobInfosByURL(job.Url, userLogs)
		if err != nil {
			return []Workgroup{}, errors.New("cannot get groups of jobs for given project: something went wrong when reaching infos for a job: " + err.Error())
		}
		groupsOfJobs = append(groupsOfJobs, Workgroup{Job: job, JobInfos: jobInfos})
	}
	return groupsOfJobs, nil
}

func GetWorkgroupsByProject(project Job, city string, userLogs JenkinsCredentials) ([]Workgroup, error) {
	yearList, err := getYearList(project, userLogs)
	if err != nil || len(yearList) == 0 {
		return []Workgroup{}, err
	}

	sortYearList(yearList)
	citiesList, err := getCities(yearList[0].Url, userLogs)
	if err != nil || len(citiesList) == 0 {
		return []Workgroup{}, err
	}

	jobsUrl := getDesiredCityJobsUrl(citiesList, city)
	if jobsUrl == "" {
		return []Workgroup{}, errors.New("cannot get groups of jobs for given project: no city containing string \"" + city + "\" in its name found")
	}

	jobsList, err := GetJobsByURL(jobsUrl, userLogs)
	if err != nil {
		return []Workgroup{}, errors.New("cannot get groups of jobs for given project: something went wrong when reaching jobs list: " + err.Error())
	}
	return getWorkgroups(jobsList, userLogs)
}
