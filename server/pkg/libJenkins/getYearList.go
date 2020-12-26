package libJenkins

import (
	"errors"
	"log"
	"strconv"
)

func checkYearList(yearList []Job) error {
	for _, year := range yearList {
		if _, err := strconv.Atoi(year.Name); err != nil {
			log.Println(err)
			return errors.New("cannot convert string \"" + year.Name + "\" to int: " + err.Error())
		}
	}
	return nil
}

func getYearList(project Job, userLogs JenkinsCredentials) ([]Job, error) {
	yearList, err := GetJobsByURL(project.Url, userLogs)
	if err != nil {
		return []Job{}, errors.New("something went wrong when reaching year list: " + err.Error())
	}
	if err = checkYearList(yearList); err != nil {
		return []Job{}, errors.New("something went wrong when reaching year list: " + err.Error())
	}
	if len(yearList) == 0 {
		return []Job{}, nil
	}
	return yearList, nil
}
