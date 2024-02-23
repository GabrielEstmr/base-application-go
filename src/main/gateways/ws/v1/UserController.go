package main_gateways_ws_v1

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_ws_commonsresources "baseapplicationgo/main/gateways/ws/commonsresources"
	main_gateways_ws_v1_request "baseapplicationgo/main/gateways/ws/v1/request"
	main_gateways_ws_v1_response "baseapplicationgo/main/gateways/ws/v1/response"
	main_usecases "baseapplicationgo/main/usecases"
	main_usecases_lockers "baseapplicationgo/main/usecases/lockers"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const _USER_CONTROLLER_MSG_KEY_MALFORMED_REQUEST_BODY = "controllers.param.missing.or.malformed"
const _USER_CONTROLLER_MSG_KEY_ARCH_ISSUE = "exceptions.architecture.application.issue"

const _USER_CONTROLLER_PATH_PREFIX = "/api/v1/users/"
const _USER_CONTROLLER_PATH_SUFFIX_ENABLE_EXTERNAL = "/authentication-providers/external-provider/enable"
const _USER_CONTROLLER_PATH_SUFFIX_ENABLE_INTERNAL = "/authentication-providers/internal-provider/enable"
const _USER_CONTROLLER_PATH_SUFFIX_REQ_CHANGE_PASS_INTERNAL = "/authentication-providers/internal-provider/enable/change-password"
const _USER_CONTROLLER_PATH_SUFFIX_ENABLE_CHANGE_PASS_INTERNAL = "/authentication-providers/internal-provider/enable/validate-change-password"

type UserController struct {
	createNewUser              *main_usecases_lockers.AtomicLockedCreateNewUser
	findUserById               *main_usecases.FindUserById
	findUsersByFilter          *main_usecases.FindUsersByFilter
	enableInternalProviderUser *main_usecases_lockers.AtomicLockedEnableInternalProviderUser
	enableExternalProviderUser *main_usecases_lockers.AtomicLockedEnableExternalProviderUser
	requestPasswordChange      *main_usecases_lockers.AtomicLockedCreateInternalAuthUserPasswordChangeRequest
	changePassword             *main_usecases_lockers.AtomicLockedChangeInternalProviderUserPassword
	messageUtils               main_utils_messages.ApplicationMessages
	logsMonitoringGateway      main_gateways.LogsMonitoringGateway
	spanGateway                main_gateways.SpanGateway
}

func NewUserController(
	createNewUser *main_usecases_lockers.AtomicLockedCreateNewUser,
	findUserById *main_usecases.FindUserById,
	findUsersByFilter *main_usecases.FindUsersByFilter,
	enableInternalProviderUser *main_usecases_lockers.AtomicLockedEnableInternalProviderUser,
	enableExternalProviderUser *main_usecases_lockers.AtomicLockedEnableExternalProviderUser,
	requestPasswordChange *main_usecases_lockers.AtomicLockedCreateInternalAuthUserPasswordChangeRequest,
	changePassword *main_usecases_lockers.AtomicLockedChangeInternalProviderUserPassword,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
) *UserController {
	return &UserController{
		createNewUser,
		findUserById,
		findUsersByFilter,
		enableInternalProviderUser,
		enableExternalProviderUser,
		requestPasswordChange,
		changePassword,
		*main_utils_messages.NewApplicationMessages(),
		logsMonitoringGateway,
		spanGateway,
	}
}

func (this *UserController) CreateUser(_ http.ResponseWriter, r *http.Request) (
	main_gateways_ws_commonsresources.ControllerResponse,
	main_domains_exceptions.ApplicationException,
) {

	span := this.spanGateway.Get(r.Context(), "UserController-CreateUser")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "Creating a new user")

	requestBody, err := io.ReadAll(r.Body)
	if err != nil || len(requestBody) == 0 {
		return *new(main_gateways_ws_commonsresources.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_USER_CONTROLLER_MSG_KEY_MALFORMED_REQUEST_BODY))
	}

	var userRequest main_gateways_ws_v1_request.CreateUserRequest
	if err = json.Unmarshal(requestBody, &userRequest); err != nil {
		return *new(main_gateways_ws_commonsresources.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_USER_CONTROLLER_MSG_KEY_MALFORMED_REQUEST_BODY))
	}

	bodyErr := userRequest.Validate()
	if bodyErr != nil {
		this.logsMonitoringGateway.ERROR(span, bodyErr.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				bodyErr.Error())
	}
	user := userRequest.ToDomain()

	persistedUser, errApp := this.createNewUser.Execute(span.GetCtx(), user)
	if errApp != nil {
		this.logsMonitoringGateway.ERROR(span, errApp.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), errApp
	}

	return *main_gateways_ws_commonsresources.NewControllerResponse(
		http.StatusCreated,
		main_gateways_ws_v1_response.NewUserResponse(persistedUser)), nil
}

func (this *UserController) EnableInternalAuthUser(_ http.ResponseWriter, r *http.Request) (
	main_gateways_ws_commonsresources.ControllerResponse,
	main_domains_exceptions.ApplicationException,
) {

	span := this.spanGateway.Get(r.Context(), "UserController-EnableInternalAuthUser")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "Enabling a new user")

	id := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, _USER_CONTROLLER_PATH_PREFIX),
		_USER_CONTROLLER_PATH_SUFFIX_ENABLE_INTERNAL)

	requestBody, err := io.ReadAll(r.Body)
	if err != nil || len(requestBody) == 0 {
		return *new(main_gateways_ws_commonsresources.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_USER_CONTROLLER_MSG_KEY_MALFORMED_REQUEST_BODY))
	}

	var userRequest main_gateways_ws_v1_request.EnableInternalUserRequest
	if err = json.Unmarshal(requestBody, &userRequest); err != nil {
		return *new(main_gateways_ws_commonsresources.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_USER_CONTROLLER_MSG_KEY_MALFORMED_REQUEST_BODY))
	}

	bodyErr := userRequest.Validate()
	if bodyErr != nil {
		this.logsMonitoringGateway.ERROR(span, bodyErr.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				bodyErr.Error())
	}

	enabledUser, errApp := this.enableInternalProviderUser.Execute(span.GetCtx(), id, userRequest.Email, userRequest.VerificationCode)
	if errApp != nil {
		this.logsMonitoringGateway.ERROR(span, errApp.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), errApp
	}

	return *main_gateways_ws_commonsresources.NewControllerResponse(
		http.StatusCreated,
		main_gateways_ws_v1_response.NewUserResponse(enabledUser)), nil
}

func (this *UserController) EnableExternalAuthUser(_ http.ResponseWriter, r *http.Request) (
	main_gateways_ws_commonsresources.ControllerResponse,
	main_domains_exceptions.ApplicationException,
) {

	span := this.spanGateway.Get(r.Context(), "UserController-EnableExternalAuthUser")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "Enabling a new user")

	id := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, _USER_CONTROLLER_PATH_PREFIX),
		_USER_CONTROLLER_PATH_SUFFIX_ENABLE_EXTERNAL)

	requestBody, err := io.ReadAll(r.Body)
	if err != nil || len(requestBody) == 0 {
		return *new(main_gateways_ws_commonsresources.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_USER_CONTROLLER_MSG_KEY_MALFORMED_REQUEST_BODY))
	}

	var userRequest main_gateways_ws_v1_request.EnableExternalUserRequest
	if err = json.Unmarshal(requestBody, &userRequest); err != nil {
		return *new(main_gateways_ws_commonsresources.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_USER_CONTROLLER_MSG_KEY_MALFORMED_REQUEST_BODY))
	}

	bodyErr := userRequest.Validate()
	if bodyErr != nil {
		this.logsMonitoringGateway.ERROR(span, bodyErr.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				bodyErr.Error())
	}

	enabledUser, errApp := this.enableExternalProviderUser.Execute(span.GetCtx(), id, userRequest.ToDomain())
	if errApp != nil {
		this.logsMonitoringGateway.ERROR(span, errApp.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), errApp
	}

	return *main_gateways_ws_commonsresources.NewControllerResponse(
		http.StatusCreated,
		main_gateways_ws_v1_response.NewUserResponse(enabledUser)), nil
}

func (this *UserController) RequestChangePasswordInternalProvider(_ http.ResponseWriter, r *http.Request) (
	main_gateways_ws_commonsresources.ControllerResponse,
	main_domains_exceptions.ApplicationException,
) {

	span := this.spanGateway.Get(r.Context(), "UserController-RequestChangePasswordInternalProvider")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "Changing password for internal provider")

	id := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, _USER_CONTROLLER_PATH_PREFIX),
		_USER_CONTROLLER_PATH_SUFFIX_REQ_CHANGE_PASS_INTERNAL)

	requestBody, err := io.ReadAll(r.Body)
	if err != nil || len(requestBody) == 0 {
		return *new(main_gateways_ws_commonsresources.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_USER_CONTROLLER_MSG_KEY_MALFORMED_REQUEST_BODY))
	}

	errApp := this.requestPasswordChange.Execute(span.GetCtx(), id)
	if errApp != nil {
		this.logsMonitoringGateway.ERROR(span, errApp.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), errApp
	}

	return *main_gateways_ws_commonsresources.NewControllerResponse(
		http.StatusAccepted,
		nil), nil
}

func (this *UserController) ValidateChangePasswordInternalProvider(_ http.ResponseWriter, r *http.Request) (
	main_gateways_ws_commonsresources.ControllerResponse,
	main_domains_exceptions.ApplicationException,
) {

	span := this.spanGateway.Get(r.Context(), "UserController-ValidateChangePasswordInternalProvider")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "Validating Changing password")

	id := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, _USER_CONTROLLER_PATH_PREFIX),
		_USER_CONTROLLER_PATH_SUFFIX_ENABLE_CHANGE_PASS_INTERNAL)

	requestBody, err := io.ReadAll(r.Body)
	if err != nil || len(requestBody) == 0 {
		return *new(main_gateways_ws_commonsresources.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_USER_CONTROLLER_MSG_KEY_MALFORMED_REQUEST_BODY))
	}

	var changePasswordRequest main_gateways_ws_v1_request.ChangeUserPasswordRequest
	if err = json.Unmarshal(requestBody, &changePasswordRequest); err != nil {
		return *new(main_gateways_ws_commonsresources.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_USER_CONTROLLER_MSG_KEY_MALFORMED_REQUEST_BODY))
	}

	bodyErr := changePasswordRequest.Validate()
	if bodyErr != nil {
		this.logsMonitoringGateway.ERROR(span, bodyErr.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				bodyErr.Error())
	}

	passwordChange, errApp := this.changePassword.Execute(
		span.GetCtx(),
		id,
		changePasswordRequest.Password,
		changePasswordRequest.VerificationCode)
	if errApp != nil {
		this.logsMonitoringGateway.ERROR(span, errApp.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), errApp
	}

	return *main_gateways_ws_commonsresources.NewControllerResponse(
		http.StatusCreated,
		main_gateways_ws_v1_response.NewUserResponse(passwordChange)), nil
}

func (this *UserController) FindUserById(
	_ http.ResponseWriter,
	r *http.Request,
) (
	main_gateways_ws_commonsresources.ControllerResponse,
	main_domains_exceptions.ApplicationException,
) {

	span := this.spanGateway.Get(r.Context(), "UserController-FindUserById")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "Finding User by id")

	id := strings.TrimPrefix(r.URL.Path, _USER_CONTROLLER_PATH_PREFIX)
	this.logsMonitoringGateway.INFO(span, fmt.Sprintf("Finding User by id: %s", id))
	persistedUser, errApp := this.findUserById.Execute(span.GetCtx(), id)
	if errApp != nil {
		this.logsMonitoringGateway.ERROR(span, errApp.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), errApp
	}

	return *main_gateways_ws_commonsresources.NewControllerResponse(
		http.StatusOK, main_gateways_ws_v1_response.NewUserResponse(persistedUser)), nil
}

func (this *UserController) FindUser(
	w http.ResponseWriter,
	r *http.Request,
) (main_gateways_ws_commonsresources.ControllerResponse,
	main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(r.Context(), "UserController-FindUser")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "Finding User by query")

	filter, err1 := new(main_gateways_ws_v1_request.FindUserFilterRequest).QueryParamsToObject(w, r)
	if err1 != nil {
		this.logsMonitoringGateway.ERROR(span, err1.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), err1
	}

	pageableReq, err2 := new(main_gateways_ws_commonsresources.PageableRequest).QueryParamsToObject(w, r)
	if err2 != nil {
		this.logsMonitoringGateway.ERROR(span, err2.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), err2
	}
	pageable, errT := pageableReq.ToDomain()
	if errT != nil {
		this.logsMonitoringGateway.ERROR(span, errT.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
			this.messageUtils.GetDefaultLocale(
				_USER_CONTROLLER_MSG_KEY_ARCH_ISSUE))
	}
	page, err := this.findUsersByFilter.Execute(span.GetCtx(), filter.ToDomain(), pageable)
	if err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), err
	}

	content := page.GetContent()
	respContent := make([]main_gateways_ws_v1_response.UserResponse, 0)
	for _, value := range content {
		user := value.(main_domains.User)
		respContent = append(respContent,
			main_gateways_ws_v1_response.NewUserResponse(user))
	}
	return *main_gateways_ws_commonsresources.NewControllerResponse(http.StatusOK,
		main_gateways_ws_commonsresources.NewPageResponse(
			respContent,
			page.GetPage(),
			page.GetSize(),
			page.GetTotalElements(),
			page.GetTotalPages())), nil
}
