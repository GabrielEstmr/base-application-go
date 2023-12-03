package main_gateways_ws_v1

import (
	main_configs_logs "baseapplicationgo/main/configs/log"
	main_configs_messages "baseapplicationgo/main/configs/messages"
	main_gateways_ws_v1_request "baseapplicationgo/main/gateways/ws/v1/request"
	main_gateways_ws_v1_response "baseapplicationgo/main/gateways/ws/v1/response"
	main_usecases "baseapplicationgo/main/usecases"
	main_utils "baseapplicationgo/main/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net"
	"net/http"
	"strings"
)

const _USER_CONTROLLER_MSG_MALFORMED_REQUEST_BODY = "controllers.param.missing.or.malformed"

const _USER_CONTROLLER_PATH_PREFIX = "/users/"

type UserController struct {
	createNewUser main_usecases.CreateNewUser
	findUserById  main_usecases.FindUserById
	apLog         *slog.Logger
}

func NewUserController(createNewUser main_usecases.CreateNewUser, findUserById main_usecases.FindUserById) *UserController {
	return &UserController{
		createNewUser,
		findUserById,
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

	// TODO get locale from ip
	ipAddress, port, err := net.SplitHostPort(ReadUserIP(r))
	ip := net.ParseIP(ipAddress)
	log.Println(ip)
	log.Println(port)

	requestBody, err := io.ReadAll(r.Body)
	if err != nil || len(requestBody) == 0 {
		errLog := errors.New(main_configs_messages.GetMessagesConfigBean().GetDefaultLocale(_USER_CONTROLLER_MSG_MALFORMED_REQUEST_BODY))
		main_utils.ERROR(w, http.StatusBadRequest, errLog)
		return
	}

	var userRequest main_gateways_ws_v1_request.CreateUserRequest
	if err = json.Unmarshal(requestBody, &userRequest); err != nil {
		main_utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	bodyErr := userRequest.Validate()
	if bodyErr != nil {
		main_utils.ERROR_APP(w, bodyErr)
		return
	}
	user := userRequest.ToDomain()

	persistedUser, errApp := this.createNewUser.Execute(user)
	if errApp != nil {
		main_utils.ERROR_APP(w, errApp)
		return
	}

	main_utils.JSON(w, http.StatusCreated, main_gateways_ws_v1_response.NewUserResponse(persistedUser))
}

func (this *UserController) FindUserById(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, _USER_CONTROLLER_PATH_PREFIX)
	this.apLog.Info(fmt.Sprintf("Finding User by id: %s", id))
	persistedUser, errApp := this.findUserById.Execute(id)
	if errApp != nil {
		main_utils.ERROR_APP(w, errApp)
		return
	}

	main_utils.JSON(w, http.StatusOK, main_gateways_ws_v1_response.NewUserResponse(persistedUser))
}

func (this *UserController) FindUser(w http.ResponseWriter, r *http.Request) {
	this.apLog.Info("Finding User by query")
	filter, err := new(main_gateways_ws_v1_request.FindUserFilterRequest).FindUserFilterToObject(w, r)
	if err != nil {
		main_utils.ERROR(w, err)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, _USER_CONTROLLER_PATH_PREFIX)
	log.Println(id)

	persistedUser, errApp := this.findUserById.Execute(id)
	if errApp != nil {
		main_utils.ERROR_APP(w, errApp)
		return
	}

	main_utils.JSON(w, http.StatusCreated, main_gateways_ws_v1_response.NewUserResponse(persistedUser))
}
