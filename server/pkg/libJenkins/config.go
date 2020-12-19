package libJenkins

var JenkinsBaseURL = getEnvVariable("JENKINS_BASE_URL")

type Logs struct {
	Login       string
	Password    string
	AccountType AccountType
}

type AccountType int

const (
	TEK2 AccountType = 0
	TEK3 AccountType = 1
)

var JenkinsLogs = map[AccountType]Logs{
	TEK2: {Login: getEnvVariable("ASSIST_REN_USER"), Password: getEnvVariable("ASSIST_REN_API_KEY"), AccountType: TEK2},
	TEK3: {Login: getEnvVariable("ASSIST3_REN_USER"), Password: getEnvVariable("ASSIST3_REN_API_KEY"), AccountType: TEK3},
}
