package main_gateways_ws_v1

import (
	main_configs_logs "baseapplicationgo/main/configs/log"
	main_gateways_ws_v1_request "baseapplicationgo/main/gateways/ws/v1/request"
	main_gateways_ws_v1_response "baseapplicationgo/main/gateways/ws/v1/response"
	main_usecases "baseapplicationgo/main/usecases"
	main_utils "baseapplicationgo/main/utils"
	"encoding/json"
	"errors"
	"io"
	"log"
	"log/slog"
	"net"
	"net/http"
)

const _MSG_MALFORMED_REQUEST_BODY = "A require param was missing or malformed"

type UserController struct {
	createNewUser main_usecases.CreateNewUser
	apLog         *slog.Logger
}

func NewUserController(createNewUser main_usecases.CreateNewUser) *UserController {
	return &UserController{
		createNewUser,
		main_configs_logs.GetLogConfigBean(),
	}
}

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func (this *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	ipAddress, port, err := net.SplitHostPort(ReadUserIP(r))
	ip := net.ParseIP(ipAddress)
	log.Println(ip)
	log.Println(port)

	requestBody, err := io.ReadAll(r.Body)
	if err != nil || len(requestBody) == 0 {
		errLog := errors.New(_MSG_MALFORMED_REQUEST_BODY)
		this.apLog.Error(errLog.Error(), "Error", err)
		main_utils.ERROR(w, http.StatusBadRequest, errLog)
		return
	}

	var userRequest main_gateways_ws_v1_request.CreateUserRequest
	if err = json.Unmarshal(requestBody, &userRequest); err != nil {
		this.apLog.Error(err.Error(), "Error", err)
		main_utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	bodyErr := userRequest.Validate()
	if bodyErr != nil {
		this.apLog.Error(bodyErr.Error(), err)
		main_utils.ERROR_APP(w, bodyErr)
		return
	}
	user := userRequest.ToDomain()

	persistedUser, errApp := this.createNewUser.Execute(user)
	if errApp != nil {
		this.apLog.Error(errApp.Error(), "Error", errApp)
		main_utils.ERROR_APP(w, errApp)
		return
	}

	main_utils.JSON(w, http.StatusCreated, main_gateways_ws_v1_response.NewUserResponse(persistedUser))
}
