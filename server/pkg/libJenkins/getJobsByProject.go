package libJenkins

import (
	"errors"
)

func GetJobsByProject(project Job, city string, userLogs JenkinsCredentials) ([]Job, error) {
	yearList, err := getYearList(project, userLogs)
	if err != nil || len(yearList) == 0 {
		return []Job{}, err
	}

	sortYearList(yearList)
	citiesList, err := getCities(yearList[0].Url, userLogs)
	if err != nil || len(citiesList) == 0 {
		return []Job{}, err
	}

	jobsUrl := getDesiredCityJobsUrl(citiesList, city)
	if jobsUrl == "" {
		return []Job{}, errors.New("cannot get jobs for given project: no city containing string \"" + city + "\" in its name found")
	}
	return GetJobsByURL(jobsUrl, userLogs)
}
