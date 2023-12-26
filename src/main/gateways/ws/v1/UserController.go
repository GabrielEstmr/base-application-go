package main_gateways_ws_v1

import (
	main_domains "baseapplicationgo/main/domains"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs_resources "baseapplicationgo/main/gateways/logs/resources"
	main_gateways_ws_commons "baseapplicationgo/main/gateways/ws/commons"
	main_gateways_ws_v1_request "baseapplicationgo/main/gateways/ws/v1/request"
	main_gateways_ws_v1_response "baseapplicationgo/main/gateways/ws/v1/response"
	main_usecases "baseapplicationgo/main/usecases"
	main_utils "baseapplicationgo/main/utils"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const _USER_CONTROLLER_MSG_MALFORMED_REQUEST_BODY = "controllers.param.missing.or.malformed"

const _USER_CONTROLLER_PATH_PREFIX = "/api/v1/users/"

type UserController struct {
	createNewUser         *main_usecases.CreateNewUser
	findUserById          *main_usecases.FindUserById
	findUsersByFilter     *main_usecases.FindUsersByFilter
	messageUtils          main_utils_messages.ApplicationMessages
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
}

func NewUserController(
	createNewUser *main_usecases.CreateNewUser,
	findUserById *main_usecases.FindUserById,
	findUsersByFilter *main_usecases.FindUsersByFilter,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
) *UserController {
	return &UserController{
		createNewUser,
		findUserById,
		findUsersByFilter,
		*main_utils_messages.NewApplicationMessages(),
		logsMonitoringGateway,
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

	span := *main_gateways_logs_resources.NewSpanLogInfo(r.Context())
	defer span.End()

	this.logsMonitoringGateway.INFO(span, "Creating a new user")

	//roll := 1 + rand.Intn(6)
	//rollCnt, err := meter.Int64Counter("dice.rolls")
	//rollValueAttr := attribute.Int("roll.value", roll)
	//if err != nil {
	//	panic(err)
	//}
	//span.SetAttributes(rollValueAttr)
	//rollCnt.Add(r.Context(), 1)

	//// TODO get locale from ip
	//ipAddress, port, err := net.SplitHostPort(ReadUserIP(r))
	//ip := net.ParseIP(ipAddress)
	//log.Println(ip)
	//log.Println(port)

	requestBody, err := io.ReadAll(r.Body)
	if err != nil || len(requestBody) == 0 {
		errLog := errors.New(
			this.messageUtils.GetDefaultLocale(
				_USER_CONTROLLER_MSG_MALFORMED_REQUEST_BODY))
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
		this.logsMonitoringGateway.ERROR(span, bodyErr.Error())
		main_utils.ERROR_APP(w, bodyErr)
		return
	}
	user := userRequest.ToDomain()

	persistedUser, errApp := this.createNewUser.Execute(r.Context(), user)
	if errApp != nil {
		this.logsMonitoringGateway.ERROR(span, errApp.Error())
		main_utils.ERROR_APP(w, errApp)
		return
	}

	main_utils.JSON(w, http.StatusCreated, main_gateways_ws_v1_response.NewUserResponse(persistedUser))
}

func (this *UserController) FindUserById(w http.ResponseWriter, r *http.Request) {

	span := *main_gateways_logs_resources.NewSpanLogInfo(r.Context())
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "Finding User by id")

	id := strings.TrimPrefix(r.URL.Path, _USER_CONTROLLER_PATH_PREFIX)
	this.logsMonitoringGateway.INFO(span, fmt.Sprintf("Finding User by id: %s", id))
	persistedUser, errApp := this.findUserById.Execute(id)
	if errApp != nil {
		this.logsMonitoringGateway.ERROR(span, errApp.Error())
		main_utils.ERROR_APP(w, errApp)
		return
	}

	main_utils.JSON(w, http.StatusOK, main_gateways_ws_v1_response.NewUserResponse(persistedUser))
}

func (this *UserController) FindUser(w http.ResponseWriter, r *http.Request) {

	span := *main_gateways_logs_resources.NewSpanLogInfo(r.Context())
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "Finding User by query")

	filter, err1 := new(main_gateways_ws_v1_request.FindUserFilterRequest).QueryParamsToObject(w, r)
	if err1 != nil {
		this.logsMonitoringGateway.ERROR(span, err1.Error())
		main_utils.ERROR_APP(w, err1)
		return
	}

	pageableReq, err2 := new(main_gateways_ws_commons.PageableRequest).QueryParamsToObject(w, r)
	if err2 != nil {
		this.logsMonitoringGateway.ERROR(span, err2.Error())
		main_utils.ERROR_APP(w, err2)
		return
	}
	pageable, errT := pageableReq.ToDomain()
	if errT != nil {
		this.logsMonitoringGateway.ERROR(span, errT.Error())
		main_utils.ERROR(w, http.StatusInternalServerError, errT)
		return
	}
	page, err := this.findUsersByFilter.Execute(filter.ToDomain(), pageable)
	if err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		main_utils.ERROR_APP(w, err)
		return
	}

	content := page.GetContent()
	respContent := make([]main_gateways_ws_v1_response.UserResponse, 0)
	for _, value := range content {
		test := value.GetObj()
		user := test.(main_domains.User)
		respContent = append(respContent,
			main_gateways_ws_v1_response.NewUserResponse(user))
	}
	main_utils.JSON(w, http.StatusOK,
		main_gateways_ws_commons.NewPageResponse(
			respContent,
			page.GetPage(),
			page.GetSize(),
			page.GetTotalElements(),
			page.GetTotalPages()))

}
