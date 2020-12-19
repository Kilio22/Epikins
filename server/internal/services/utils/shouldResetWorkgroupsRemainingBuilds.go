package utils

import (
	"time"

	"epikins-api/internal"
)

func ShouldResetWorkgroupsRemainingBuilds(projectData internal.MongoProjectData) bool {
	return time.Since(time.Unix(projectData.LastUpdate, 0)).Hours() >= float64(24*7)
}
