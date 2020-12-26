package config

import "epikins-api/internal"

const (
	PROJECTS    internal.Role = "projects"
	USERS       internal.Role = "users"
	CREDENTIALS internal.Role = "credentials"
	MODULE      internal.Role = "module"
)

var Roles = []internal.Role{
	PROJECTS,
	USERS,
	CREDENTIALS,
	MODULE,
}

const DefaultBuildNb int = 3
