package libJenkins

import (
	"errors"
	"net/http"
	"net/url"
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

func BuildJob(postUrl string, visibility Visibility, credentials JenkinsCredentials) error {
	form := getForm(visibility)
	_, err := makeHttpRequest(http.MethodPost, postUrl, JenkinsBuildURLPart, credentials, form.Encode())
	if err != nil {
		return errors.New("cannot build job: " + err.Error())
	}
	return nil
}
