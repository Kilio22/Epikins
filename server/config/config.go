package config

import (
	"epikins-api/internal"
)

const (
	PROJECTS    internal.Role = "projects"
	USERS       internal.Role = "users"
	CREDENTIALS internal.Role = "credentials"
	MODULE      internal.Role = "module"
	STUDENT     internal.Role = "student"
)

var Roles = []internal.Role{
	PROJECTS,
	USERS,
	CREDENTIALS,
	MODULE,
}

const DefaultBuildNb int = 3
const HighestPrivilegeJenkinsLogin = "assist3_ren"
