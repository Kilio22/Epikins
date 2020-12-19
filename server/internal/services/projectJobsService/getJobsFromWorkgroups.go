package projectJobsService

import "epikins-api/pkg/libJenkins"

func getJobsFromWorkgroups(workgroups []libJenkins.Workgroup) []libJenkins.Job {
	var jobs []libJenkins.Job
	for _, groupJob := range workgroups {
		jobs = append(jobs, groupJob.Job)
	}
	return jobs
}
