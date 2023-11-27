package main_gateways_ws_v1

import (
	main_gateways_ws_v1_request "baseapplicationgo/main/gateways/ws/v1/request"
	main_gateways_ws_v1_response "baseapplicationgo/main/gateways/ws/v1/response"
	main_usecases "baseapplicationgo/main/usecases"
	main_utils "baseapplicationgo/main/utils"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type UserController struct {
	createNewUser main_usecases.CreateNewUser
}

func NewUserController(createNewUser main_usecases.CreateNewUser) *UserController {
	return &UserController{createNewUser}
}

func (this *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	requestBody, err := io.ReadAll(r.Body)
	if err != nil || len(requestBody) == 0 {
		main_utils.ERROR(w, http.StatusBadRequest, errors.New("A require param was missing or malformed"))
		return
	}

	var userRequest main_gateways_ws_v1_request.CreateUserRequest
	if err = json.Unmarshal(requestBody, &userRequest); err != nil {
		main_utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err2 := userRequest.Validate()
	if err2 != nil {
		main_utils.ERROR(w, http.StatusBadRequest, err2)
		return
	}
	user := userRequest.ToDomain()

	persistedUser, err := this.createNewUser.Execute(user)
	if err != nil {
		main_utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	main_utils.JSON(w, http.StatusCreated, main_gateways_ws_v1_response.NewUserResponse(persistedUser))
}
