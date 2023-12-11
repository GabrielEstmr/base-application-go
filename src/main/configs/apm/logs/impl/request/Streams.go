package main_configs_apm_logs_impl_request

import (
	"strconv"
	"time"
)

type Streams struct {
	Stream Stream     `json:"stream"`
	Values [][]string `json:"values"`
}

func NewStreams(
	level string,
	job string,
	logMsg string) *Streams {
	value := make(map[string]string)
	value[strconv.FormatInt(time.Now().UnixNano(), 10)] = logMsg

	var valueList [][]string
	valueList = append(valueList, []string{strconv.FormatInt(time.Now().UnixNano(), 10), logMsg})

	return &Streams{
		Stream: *NewStream(level, job),
		Values: valueList}
}
