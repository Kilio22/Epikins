package utils

import (
	"epikins-api/internal"
	"time"
)

func shouldResetWorkgroupsRemainingBuilds(projectData internal.MongoProjectData) bool {
	return time.Since(time.Unix(projectData.LastUpdate, 0)).Hours() >= float64(24*7)
}
