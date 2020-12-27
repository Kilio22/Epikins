package libJenkins

var JenkinsBaseURL = getEnvVariable("JENKINS_BASE_URL")

const JenkinsBuildURLPart string = "/build?delay=0"
const JenkinsAPIPart string = "/api/json"

type JenkinsCredentials struct {
	Login  string `json:"login" validate:"required"`
	ApiKey string `json:"apiKey" validate:"required"`
}
