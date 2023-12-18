package main_domains

import "time"

type Transaction struct {
	id               string
	accountId        string
	operationTypeId  string
	amount           float64
	createdDate      time.Time
	lastModifiedDate time.Time
}

func NewTransactionAllArgs(
	id string,
	accountId string,
	operationTypeId string,
	amount float64,
	createdDate time.Time,
	lastModifiedDate time.Time,
) *Transaction {
	return &Transaction{
		id:               id,
		accountId:        accountId,
		operationTypeId:  operationTypeId,
		amount:           amount,
		createdDate:      createdDate,
		lastModifiedDate: lastModifiedDate,
	}
}

func NewTransaction(
	accountId string,
	operationTypeId string,
	amount float64,
) *Transaction {
	return &Transaction{
		accountId:       accountId,
		operationTypeId: operationTypeId,
		amount:          amount,
	}
}

func (this *Transaction) GetId() string {
	return this.id
}

func (this *Transaction) GetAccountId() string {
	return this.accountId
}

func (this *Transaction) GetOperationTypeId() string {
	return this.operationTypeId
}

func (this *Transaction) GetAmount() float64 {
	return this.amount
}

func (this *Transaction) GetCreatedDate() time.Time {
	return this.createdDate
}

func (this *Transaction) GetLastModifiedDate() time.Time {
	return this.lastModifiedDate
}

func (this *Transaction) IsEmpty() bool {
	return *this == Transaction{}
}
