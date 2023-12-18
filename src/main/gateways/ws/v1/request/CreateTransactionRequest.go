package main_gateways_ws_v1_request

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils "baseapplicationgo/main/utils"
)

type CreateTransactionRequest struct {
	AccountId       string  `json:"accountId" validate:"required,min=4,max=15"`
	OperationTypeId string  `json:"operationTypeId" validate:"required,min=4,max=15"`
	Amount          float64 `json:"amount" validate:"required"`
}

func (this *CreateTransactionRequest) ToDomain() main_domains.Transaction {
	return *main_domains.NewTransaction(
		this.AccountId,
		this.OperationTypeId,
		this.Amount)
}

func (this *CreateTransactionRequest) Validate() main_domains_exceptions.ApplicationException {
	structValidatorMessages := main_utils.NewStructValidatorMessages(this)
	if len(structValidatorMessages.GetMessages()) != 0 {
		return main_domains_exceptions.NewBadRequestException(structValidatorMessages.GetMessages())
	}
	return nil
}
