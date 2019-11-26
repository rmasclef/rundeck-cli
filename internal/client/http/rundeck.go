package http

import (
	"fmt"
	"net/http"

	"github.com/gemalto/requester"

	rundeck_api "github.com/rmasclef/rundeck-cli/pkg/rundeck-api/http"
	"github.com/rmasclef/rundeck-cli/pkg/rundeck-api/http/request/body"
	"github.com/rmasclef/rundeck-cli/pkg/rundeck-api/http/request/query-params"
	"github.com/rmasclef/rundeck-cli/pkg/rundeck-api/http/response"
)

type RundeckClient struct {
	requester *requester.Requester
	//@TODO pass desired response format (XML or JSON)
}

// NewRundeckClient creates an http requester to call a rundeck api
func NewRundeckClient(apiUrl string, httpClient *http.Client) *RundeckClient {
	return &RundeckClient{
		requester: requester.MustNew(
			requester.URL(apiUrl),
			requester.WithDoer(httpClient),
			requester.ExpectCode(http.StatusOK),
			//@TODO set Marshaler in function of desired response format (XML or JSON)
			requester.WithMarshaler(&requester.JSONMarshaler{}),
		),
	}
}

// ListJobs will call list jobs endpoint
// @see https://docs.rundeck.com/docs/api/rundeck-api.html#listing-jobs
func (client *RundeckClient) ListJobs(projectName string, queryParams query_params.ListJobs) (*response.ListJobs, error) {
	listJobsResponse := new(response.ListJobs)
	requesterOptions := []requester.Option{
		requester.Get(fmt.Sprintf(rundeck_api.ListJobsUrl, projectName)),
		// add query params
		requester.QueryParams(queryParams),
	}
	// execute request
	_, _, err := client.requester.Receive(
		listJobsResponse,
		requesterOptions...,
	)
	return listJobsResponse, err
}

// RunJob will call run job endpoint
// @see https://docs.rundeck.com/docs/api/rundeck-api.html#running-a-job
func (client *RundeckClient) RunJob(jobID string, body *body.RunJob) (*response.RunJob, error) {
	runJobResponse := new(response.RunJob)
	requesterOptions := []requester.Option{
		requester.Post(fmt.Sprintf(rundeck_api.RunJobUrl, jobID)),
		// pass request body
		requester.Body(body),
	}
	// execute request
	_, _, err := client.requester.Receive(
		runJobResponse,
		requesterOptions...,
	)
	return runJobResponse, err
}
