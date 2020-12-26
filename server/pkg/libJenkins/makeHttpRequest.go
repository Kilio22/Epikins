package libJenkins

import (
	"errors"
	"log"
	"net/http"
)

const JenkinsAPIPart string = "/api/json"

func makeHttpRequest(method string, url string, userLogs JenkinsCredentials) (*http.Response, error) {
	var fullUrl string
	if url[len(url)-1:] == "/" {
		fullUrl = url[:len(url)-1] + JenkinsAPIPart
	} else {
		fullUrl = url + JenkinsAPIPart
	}

	req, err := http.NewRequest(method, fullUrl, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.SetBasicAuth(userLogs.Login, userLogs.ApiKey)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if res.StatusCode < http.StatusOK || res.StatusCode > http.StatusIMUsed {
		return nil, errors.New("bad response code when making request to \"" + url + "\"")
	}
	return res, nil
}
