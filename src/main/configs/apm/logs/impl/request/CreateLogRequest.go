package main_configs_apm_logs_impl_request

type CreateLogRequest struct {
	Streams []Streams `json:"streams"`
}

func NewCreateLogRequest(
	level string,
	job string,
	logMsg string) *CreateLogRequest {
	streamsList := []Streams{*NewStreams(level, job, logMsg)}
	return &CreateLogRequest{Streams: streamsList}
}
