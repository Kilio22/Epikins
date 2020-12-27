package libJenkins

import (
	"errors"
	"log"
	"net/http"
	"strings"
)

func initRequest(method string, fullUrl string, credentials JenkinsCredentials, body string) (*http.Request, error) {
	var req *http.Request
	var err error

	if body != "" {
		req, err = http.NewRequest(method, fullUrl, strings.NewReader(body))
	} else {
		req, err = http.NewRequest(method, fullUrl, nil)
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.SetBasicAuth(credentials.Login, credentials.ApiKey)
	if body != "" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	return req, nil
}

func getFullUrl(url string, endUrlPart string) string {
	var fullUrl string
	if url[len(url)-1:] == "/" {
		fullUrl = url[:len(url)-1] + endUrlPart
	} else {
		fullUrl = url + endUrlPart
	}
	return fullUrl
}

func makeHttpRequest(method string, url string, endUrlPart string, credentials JenkinsCredentials, body string) (*http.Response, error) {
	fullUrl := getFullUrl(url, endUrlPart)
	req, err := initRequest(method, fullUrl, credentials, body)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if res.StatusCode < http.StatusOK || res.StatusCode > http.StatusIMUsed {
		return nil, errors.New("bad response code when making request to \"" + fullUrl + "\"")
	}
	return res, nil
}
