package main_gateways_rabbitmq_resources

import main_domains "baseapplicationgo/main/domains"

type TransactionResource struct {
	AccountId       string  `json:"accountId"`
	OperationTypeId string  `json:"operationTypeId"`
	Amount          float64 `json:"amount"`
}

func NewTransactionResource(transaction main_domains.Transaction) *TransactionResource {
	return &TransactionResource{
		AccountId:       transaction.GetAccountId(),
		OperationTypeId: transaction.GetOperationTypeId(),
		Amount:          transaction.GetAmount(),
	}
}

func NewTransactionResourceFromProps(props map[string]interface{}) *TransactionResource {
	return &TransactionResource{
		AccountId:       props["accountId"].(string),
		OperationTypeId: props["operationTypeId"].(string),
		Amount:          props["amount"].(float64),
	}
}

func (this *TransactionResource) ToDomain() main_domains.Transaction {
	if (*this == TransactionResource{}) {
		return *new(main_domains.Transaction)
	}
	return *main_domains.NewTransaction(
		this.AccountId,
		this.OperationTypeId,
		this.Amount,
	)
}
