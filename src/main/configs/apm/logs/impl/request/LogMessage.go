package main_configs_apm_logs_impl_request

import (
	"encoding/json"
	"fmt"
	"strings"
)

type LogMessage struct {
	TraceId  string `json:"trace_id"`
	SpanId   string `json:"spanid"`
	Severity string `json:"severity"`
	Body     string `json:"body"`
	UserId   string `json:"user_id"`
	Flags    int    `json:"flags"`
}

func NewLogMessage(
	traceId string,
	spanId string,
	severity string,
	body string,
	userId string, flags int) *LogMessage {
	return &LogMessage{
		TraceId:  traceId,
		SpanId:   spanId,
		Severity: severity,
		Body:     body,
		UserId:   userId, Flags: flags}
}

func (this *LogMessage) TO_STRING_JSON() (string, error) {

	b, err := json.Marshal(this)
	strings.ReplaceAll(string(b), `"`, `\"`)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(b), nil
}
