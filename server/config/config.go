package config

import (
	"epikins-api/internal"
)

const AppIdKey = "APP_ID"
const TenantIdKey = "TENANT_ID"
const ServerPortKey = "SERVER_PORT"
const MongoHostKey = "MONGO_HOST"
const MongoPortKey = "MONGO_PORT"
const MongoUsernameKey = "MONGO_INITDB_ROOT_USERNAME"
const MongoPasswordKey = "MONGO_INITDB_ROOT_PASSWORD"
const MongoDbKey = "MONGO_INITDB_DATABASE"
const StudentJenkinsLoginKey = "STUDENT_JENKINS_LOGIN"
const IntraAutologinLinkKey = "INTRA_AUTOLOGIN_LINK"

const (
	CREDENTIALS internal.Role = "credentials"
	LOG         internal.Role = "log"
	MODULE      internal.Role = "module"
	PROJECTS    internal.Role = "projects"
	STUDENT     internal.Role = "student"
	USERS       internal.Role = "users"
)

var Roles = []internal.Role{
	CREDENTIALS,
	LOG,
	MODULE,
	PROJECTS,
	USERS,
}

const DefaultBuildNb int = 0
const LocalProjectListRefreshTime float64 = 1.0
const ProjectJobsRefreshTime float64 = 12.0
