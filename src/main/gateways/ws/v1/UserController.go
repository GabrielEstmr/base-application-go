package main_gateways_ws_v1

import (
	main_usecases "baseapplicationgo/main/usecases"
	utils "baseapplicationgo/main/utils"
	"net/http"
	"time"
)

const IDX_TRACING_FIND_USER_CONTROLLER = "find-indicator-controller"

//const ACCOUNT_PATH_PREFIX = "/accounts/"

type UserController struct {
	createNewUser main_usecases.CreateNewUser
}

func NewUserController(createNewUser *main_usecases.CreateNewUser) *UserController {
	return &UserController{*createNewUser}
}

func (this *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	//ctx := baggage.ContextWithoutBaggage(r.Context())
	//tr := otel.GetTracerProvider().Tracer(main_domains.APP_INDICATOR_TYPE_FIND_INDICATOR.GetDescription())
	//ctx, controller := tr.Start(ctx, IDX_TRACING_FIND_INDICATOR_CONTROLLER)
	//defer controller.End()

	//main_gateways_rabbitmq_producers.Produce(&ctx)

	//id := strings.TrimPrefix(r.URL.Path, ACCOUNT_PATH_PREFIX)
	user, err := this.createNewUser.Execute("name", "documentNumber", time.Now())
	if err != nil {
		return
	}
	//log.Println(id)

	utils.JSON(w, http.StatusOK, user)

}
