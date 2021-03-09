package util

import (
	"time"

	"epikins-api/internal"
)

func GetNewBuildLog(module string, starter string, target string, project string) internal.BuildLogs {
	return internal.BuildLogs{
		Module:  module,
		Project: project,
		Starter: starter,
		Target:  target,
		Time:    time.Now().Unix(),
	}
}
