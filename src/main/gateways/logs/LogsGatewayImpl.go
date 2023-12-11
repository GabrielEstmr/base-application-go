package main_gateways_logs

import (
	main_configs_apm_logs "baseapplicationgo/main/configs/apm/logs"
	main_configs_apm_logs_resources "baseapplicationgo/main/configs/apm/logs/resources"
	main_gateways_logs_request "baseapplicationgo/main/gateways/logs/request"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

type LogsGatewayImpl struct {
	logConfig main_configs_apm_logs_resources.LogProviderConfig
}

func NewLogsGatewayImpl() *LogsGatewayImpl {
	return &LogsGatewayImpl{
		logConfig: *main_configs_apm_logs.GetLogProviderBean(),
	}
}

func (this *LogsGatewayImpl) Error(
	traceId string,
	spanId string,
	msg string,
	args ...any,
) {

	client := http.Client{
		Timeout: 2000 * time.Millisecond,
	}
	baseUrl := "http://" + this.logConfig.GetHost() + "/loki/api/v1/push"

	stringJson, err := main_gateways_logs_request.NewLogMessage(traceId, spanId, "ERROR", msg, "", 1).TO_STRING_JSON()
	if err != nil {
		fmt.Println(err.Error())
	}

	request := main_gateways_logs_request.NewCreateLogRequest("ERROR", this.logConfig.GetAppName(), stringJson)

	b, err := json.Marshal(request)
	jsonRequest := string(b)
	fmt.Println(jsonRequest)

	body, _ := json.Marshal(request)
	payload := bytes.NewBuffer(body)

	resp, errPost := client.Post(baseUrl, "application/json", payload)
	if errPost != nil {
		fmt.Println("ERROR ==============", errPost.Error())
	}

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	//slog.Error(msg, args)
	//
	//stringJson, err := main_gateways_logs_request.NewLogMessage(traceId, spanId, "ERROR", msg, "", 1).TO_STRING_JSON()
	//if err != nil {
	//	return
	//}
	//request := main_gateways_logs_request.NewCreateLogRequest("ERROR", this.logConfig.GetAppName(), stringJson)
	//
	//body, _ := json.Marshal(request)
	//
	//payload := bytes.NewBuffer(body)
	//
	//req, err := http.NewRequest("POST", this.logConfig.GetHost()+"/loki/api/v1/push", bytes.NewBuffer(body))
	//
	//client := &http.Client{}
	//baseUrl := this.logConfig.GetHost() + "/loki/api/v1/push"
	//
	//req, err := http.NewRequest(http.MethodPost, baseUrl, nil)
	//_, err = client.Do(req)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//
	//
	//req, err := http.NewRequest(http.MethodPost, baseUrl, nil)
	//req.Header.Set("Content-Type", "application/json")
	//
	//client := &http.Client{}
	//
	//_, errP := http.Post(this.logConfig.GetHost()+"/loki/api/v1/push", "application/json", nil)
	//if errP != nil {
	//	return
	//}
	//
	//resp, err := client.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//
	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//bodyResp, _ := io.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(bodyResp))
	//
	//slog.Error(msg, args)

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	return rr
}
