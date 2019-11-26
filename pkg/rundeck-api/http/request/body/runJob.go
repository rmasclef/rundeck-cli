package body

// RunJob represent the request body of runJob endpoint
// @see https://docs.rundeck.com/docs/api/rundeck-api.html#running-a-job
type RunJob struct {
	ArgString string `json:"arg_string"`
}

// WithArgString to pass args to your job
func (b *RunJob) WithArgString(argString string) *RunJob {
	b.ArgString = argString
	return b
}
