package libJenkins

var JenkinsBaseURL = getEnvVariable("JENKINS_BASE_URL")

type JenkinsCredentials struct {
	Login  string `json:"login" validate:"required"`
	ApiKey string `json:"apiKey" validate:"required"`
}
