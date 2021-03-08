package config

import (
	"epikins-api/internal"
)

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
