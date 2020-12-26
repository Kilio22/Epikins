package libJenkins

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Visibility string

const (
	PRIVATE Visibility = "Private"
	PUBLIC  Visibility = "Public"
)

var VisibilityMap = map[string]Visibility{
	"private": PRIVATE,
	"public":  PUBLIC,
}

const JenkinsBuildURLPart string = "/build?delay=0"

func getFullUrl(postUrl string) string {
	fullUrl := postUrl

	if postUrl[len(postUrl)-1] == '/' {
		if postUrl[len(postUrl)-1:] == "/" {
			fullUrl = postUrl[:len(postUrl)-1] + JenkinsBuildURLPart
		} else {
			fullUrl = postUrl + JenkinsBuildURLPart
		}
	}
	return fullUrl
}

func getForm(visibility Visibility) (form url.Values) {
	form = url.Values{}
	form.Add("json", string(
		"{\"parameter\":"+
			" [{\"name\": \"VISIBILITY\", \"value\": \""+visibility+"\"},"+
			" {\"name\": \"DELIVERY\", \"value\": \"Git\"},"+
			" {\"name\": \"FORCE\", \"value\": false}]"+
			"}"))
	return
}

func BuildJob(postUrl string, visibility Visibility, logs JenkinsCredentials) error {
	form := getForm(visibility)
	fullUrl := getFullUrl(postUrl)
	req, err := http.NewRequest(http.MethodPost, fullUrl, strings.NewReader(form.Encode()))
	if err != nil {
		log.Println(err)
		return errors.New("cannot build job: " + err.Error())
	}

	req.SetBasicAuth(logs.Login, logs.ApiKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return errors.New("cannot build job: " + err.Error())
	}

	if res.StatusCode < http.StatusOK || res.StatusCode > http.StatusIMUsed {
		return errors.New("cannot build job: bad response code")
	}
	return nil
}
