package main_gateways_rabbitmq_resources

import main_domains "baseapplicationgo/main/domains"

type TransactionResource struct {
	Id              string `json:"id"`
	AccountId       string `json:"accountId"`
	OperationTypeId string `json:"operationTypeId"`
}

func NewTransactionResource(transaction main_domains.Transaction) *TransactionResource {
	return &TransactionResource{
		Id:              transaction.GetId(),
		AccountId:       transaction.GetAccountId(),
		OperationTypeId: transaction.GetOperationTypeId(),
	}
}
