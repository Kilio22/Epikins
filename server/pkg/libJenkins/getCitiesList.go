package libJenkins

import (
	"errors"
)

func getCitiesList(url string, userLogs Logs) ([]Job, error) {
	citiesList, err := GetJobsByURL(url, userLogs)
	if err != nil {
		return []Job{}, errors.New("something went wrong when reaching cities list: " + err.Error())
	}
	if len(citiesList) == 0 {
		return []Job{}, nil
	}
	return citiesList, nil
}
