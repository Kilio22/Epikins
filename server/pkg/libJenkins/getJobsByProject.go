package libJenkins

import (
	"errors"
)

func GetJobsByProject(project Job, userLogs Logs) ([]Job, error) {
	yearList, err := getYearList(project, userLogs)
	if err != nil || len(yearList) == 0 {
		return []Job{}, err
	}

	sortYearList(yearList)
	citiesList, err := getCitiesList(yearList[0].Url, userLogs)
	if err != nil || len(citiesList) == 0 {
		return []Job{}, err
	}

	jobsUrl := getWantedCityJobUrl(citiesList, "REN")
	if jobsUrl == "" {
		return []Job{}, errors.New("cannot get jobs for given project: no city containing string \"REN\" in its name found")
	}
	return GetJobsByURL(jobsUrl, userLogs)
}
