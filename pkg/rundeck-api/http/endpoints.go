package http

/*
	This file contains each available rundeck api endpoints
*/

const (
	// @TODO see if we can use named params such as {projectId}
	ListJobsUrl = "/project/%s/jobs"
	RunJobUrl = "/job/%s/run"
)
