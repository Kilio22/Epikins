package libJenkins

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type BuildInfos struct {
	Building bool `json:"building"`
}

type Build struct {
	Number     int        `json:"number"`
	Url        string     `json:"url"`
	BuildInfos BuildInfos `json:"buildInfos"`
}

type JobInfos struct {
	InQueue   bool  `json:"inQueue"`
	LastBuild Build `json:"lastBuild"`
}

func getJobInfosByURL(url string, logs JenkinsCredentials) (JobInfos, error) {
	res, err := makeHttpRequest(http.MethodGet, url, JenkinsAPIPart, logs, "")
	if err != nil {
		return JobInfos{}, errors.New("cannot get job infos: " + err.Error())
	}

	var jobInfos JobInfos
	err = json.NewDecoder(res.Body).Decode(&jobInfos)
	_ = res.Body.Close()
	if err != nil {
		log.Println(err)
		return JobInfos{}, errors.New("cannot get job infos: " + err.Error())
	}

	if jobInfos.LastBuild.Number == 0 {
		return jobInfos, nil
	}
	res, err = makeHttpRequest(http.MethodGet, jobInfos.LastBuild.Url, JenkinsAPIPart, logs, "")
	if err != nil {
		return JobInfos{}, errors.New("cannot get job infos: " + err.Error())
	}

	err = json.NewDecoder(res.Body).Decode(&jobInfos.LastBuild.BuildInfos)
	if err != nil {
		return JobInfos{}, errors.New("cannot get job infos: " + err.Error())
	}
	return jobInfos, nil
}
