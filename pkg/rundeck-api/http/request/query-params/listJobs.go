package query_params

import "github.com/rmasclef/rundeck-cli/pkg/rundeck-api/request"

// ListJobs contains each query params related to ListingJobs endpoint
// @see https://docs.rundeck.com/docs/api/rundeck-api.html#listing-jobs
type ListJobs request.Params

// WithJobExactFilter allows you to list jobs having a specific name
func (b ListJobs) WithJobExactFilter(jobName string) ListJobs {
	// add request "jobExactFilter" query param
	b.Set("jobExactFilter", jobName)
	return b
}
