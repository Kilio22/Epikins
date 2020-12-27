package libJenkins

import "strings"

func getDesiredCityJobsUrl(citiesList []Job, wantedCity string) string {
	for _, cityJobs := range citiesList {
		if strings.Contains(cityJobs.Name, wantedCity) && !strings.Contains(cityJobs.Name, "jobs") {
			return cityJobs.Url
		}
	}
	return ""
}
