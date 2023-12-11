package main_configs_apm_logs_impl_request

type Stream struct {
	Level string `json:"level"`
	Job   string `json:"job"`
}

func NewStream(level string, job string) *Stream {
	return &Stream{Level: level, Job: job}
}
