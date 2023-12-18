package main_gateways_ws_v1_response

import (
	main_domains "baseapplicationgo/main/domains"
	"time"
)

type TransactionResponse struct {
	Id               string    `json:"id"`
	AccountId        string    `json:"accountId"`
	OperationTypeId  string    `json:"operationTypeId"`
	Amount           float64   `json:"amount"`
	CreatedDate      time.Time `json:"createdDate"`
	LastModifiedDate time.Time `json:"lastModifiedDate"`
}

func NewTransactionResponse(transaction main_domains.Transaction) *TransactionResponse {
	return &TransactionResponse{
		Id:               transaction.GetId(),
		AccountId:        transaction.GetAccountId(),
		OperationTypeId:  transaction.GetOperationTypeId(),
		Amount:           transaction.GetAmount(),
		CreatedDate:      transaction.GetCreatedDate(),
		LastModifiedDate: transaction.GetLastModifiedDate(),
	}
}
