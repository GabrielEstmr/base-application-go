package main_gateways_ws_v1_request

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils "baseapplicationgo/main/utils"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"net/http"
)

const _MSG_EXTERNAL_PROVIDER_MALFORMED_REQUEST_BODY = "controllers.param.missing.or.malformed"

type ExternalProviderSessionRequest struct {
	SubjectIssuer []string `json:"subject_issuer,omitempty" validate:"required,min=1,max=1"`
	SubjectToken  []string `json:"subject_token,omitempty" validate:"required,min=1,max=1"`
}

func (this *ExternalProviderSessionRequest) QueryParamsToObject(
	w http.ResponseWriter,
	r *http.Request) (*ExternalProviderSessionRequest, main_domains_exceptions.ApplicationException) {
	filter := main_utils.NewQueryParams(this)
	object, err := main_utils.QueryParamsToObject(filter, w, r)
	if err != nil {
		return new(ExternalProviderSessionRequest), err
	}
	obj := object.GetObj()
	return obj.(*ExternalProviderSessionRequest), err
}

func (this *ExternalProviderSessionRequest) GetFirstSubjectIssuer() string {
	if this.SubjectIssuer == nil {
		return main_utils.STRING_UTILS_EMPTY_STRING
	}
	if len(this.SubjectIssuer) == 0 {
		return main_utils.STRING_UTILS_EMPTY_STRING
	}
	return this.SubjectIssuer[0]
}

func (this *ExternalProviderSessionRequest) GetFirstSubjectToken() string {
	if this.SubjectToken == nil {
		return main_utils.STRING_UTILS_EMPTY_STRING
	}
	if len(this.SubjectToken) == 0 {
		return main_utils.STRING_UTILS_EMPTY_STRING
	}
	return this.SubjectToken[0]
}

func (this *ExternalProviderSessionRequest) Validate() main_domains_exceptions.ApplicationException {
	structValidatorMessages := main_utils.NewStructValidatorMessages(this)
	if len(structValidatorMessages.GetMessages()) != 0 {
		return main_domains_exceptions.NewBadRequestException(structValidatorMessages.GetMessages())
	}
	return nil
}

func (this *ExternalProviderSessionRequest) ToDomain() (
	main_domains.ExternalProviderSessionArgs,
	main_domains_exceptions.ApplicationException) {

	provider := new(main_domains_enums.AuthProviderType).FromValue(this.GetFirstSubjectIssuer())
	if main_utils.NewStringUtils().IsEmpty(provider.Name()) {
		return *new(main_domains.ExternalProviderSessionArgs), main_domains_exceptions.NewBadRequestExceptionSglMsg(
			main_utils_messages.NewApplicationMessages().GetDefaultLocale(
				_MSG_EXTERNAL_PROVIDER_MALFORMED_REQUEST_BODY))
	}
	return *main_domains.NewExternalProviderSessionArgs(this.GetFirstSubjectToken(), provider), nil
}
