package main_configs_apm_logs_impl_request

type T struct {
	Streams []struct {
		Stream struct {
			Level string `json:"level"`
			Job   string `json:"job"`
		} `json:"stream"`
		Values [][]string `json:"values"`
	} `json:"streams"`
}
