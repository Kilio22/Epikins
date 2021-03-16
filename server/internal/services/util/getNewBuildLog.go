package util

import (
	"time"

	"epikins-api/internal"
)

func GetNewBuildLog(city string, module string, starter string, target string, project string) internal.BuildLogElem {
	return internal.BuildLogElem{
		City:    city,
		Module:  module,
		Project: project,
		Starter: starter,
		Target:  target,
		Time:    time.Now().Unix(),
	}
}
