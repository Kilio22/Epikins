package libJenkins

import (
	"errors"
	"strings"
)

func getWantedGlobalJobUrl(citiesList []Job, wantedGlobalJob string) string {
	for _, job := range citiesList {
		if strings.Contains(job.Name, wantedGlobalJob) && strings.Contains(job.Name, "jobs") {
			return job.Url
		}
	}
	return ""
}

func GetGlobalJobUrlByProject(project Job, userLogs JenkinsCredentials) (string, error) {
	yearList, err := getYearList(project, userLogs)
	if err != nil || len(yearList) == 0 {
		return "", err
	}

	sortYearList(yearList)
	citiesList, err := getCitiesList(yearList[0].Url, userLogs)
	if err != nil || len(citiesList) == 0 {
		return "", err
	}

	globalJobUrl := getWantedGlobalJobUrl(citiesList, "REN")
	if globalJobUrl == "" {
		return "", errors.New("cannot get global job URL: no city containing strings \"REN\" and \"jobs\" in its names found")
	}
	return globalJobUrl, nil
}
