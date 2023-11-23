package main_gateways_ws_v1

import (
	utils "baseapplicationgo/main/utils"
	"net/http"
)

const IDX_TRACING_FIND_INDICATOR_CONTROLLER = "find-indicator-controller"

//const ACCOUNT_PATH_PREFIX = "/accounts/"

type AccountController struct {
	//apiIndicatorTrigger                 main_usecases.ApiIndicatorTrigger
	//findIndicatorByIdWithCacheAndLocked main_usecases.FindIndicatorByIdWithCacheAndLocked
}

func NewIndicatorController(
// useCaseApiIndicatorTrigger *main_usecases.ApiIndicatorTrigger,
// findIndicatorByIdWithCacheAndLocked *main_usecases.FindIndicatorByIdWithCacheAndLocked
) *AccountController {
	return &AccountController{
		//apiIndicatorTrigger:                 *useCaseApiIndicatorTrigger,
		//findIndicatorByIdWithCacheAndLocked: *findIndicatorByIdWithCacheAndLocked,
	}
}

func (thisController *AccountController) FindAccount(w http.ResponseWriter, r *http.Request) {

	//ctx := baggage.ContextWithoutBaggage(r.Context())
	//tr := otel.GetTracerProvider().Tracer(main_domains.APP_INDICATOR_TYPE_FIND_INDICATOR.GetDescription())
	//ctx, controller := tr.Start(ctx, IDX_TRACING_FIND_INDICATOR_CONTROLLER)
	//defer controller.End()

	//main_gateways_rabbitmq_producers.Produce(&ctx)

	//id := strings.TrimPrefix(r.URL.Path, ACCOUNT_PATH_PREFIX)
	indicator := "thisController.findIndicatorByIdWithCacheAndLocked.Execute(&ctx, id)"
	//log.Println(id)

	utils.JSON(w, http.StatusOK, indicator)

}
