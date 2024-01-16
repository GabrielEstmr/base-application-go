package main_gateways_mongodb_documents

import (
	main_domains "baseapplicationgo/main/domains"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionDocument struct {
	Id               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AccountId        string             `json:"accountId,omitempty" bson:"accountId,omitempty"`
	OperationTypeId  string             `json:"operationTypeId,omitempty" bson:"operationTypeId,omitempty"`
	Amount           float64            `json:"amount,omitempty" bson:"amount,omitempty"`
	CreatedDate      primitive.DateTime `json:"createdDate,omitempty" bson:"createdDate"`
	LastModifiedDate primitive.DateTime `json:"lastModifiedDate,omitempty" bson:"lastModifiedDate"`
}

func NewTransactionDocument(transaction main_domains.Transaction) TransactionDocument {
	// TODO: create id from user.id
	return TransactionDocument{
		AccountId:        transaction.GetAccountId(),
		OperationTypeId:  transaction.GetOperationTypeId(),
		Amount:           transaction.GetAmount(),
		CreatedDate:      primitive.NewDateTimeFromTime(transaction.GetCreatedDate()),
		LastModifiedDate: primitive.NewDateTimeFromTime(transaction.GetLastModifiedDate()),
	}
}

func (this *TransactionDocument) IsEmpty() bool {
	return *this == TransactionDocument{}
}

func (this *TransactionDocument) ToDomain() main_domains.Transaction {
	if (*this == TransactionDocument{}) {
		return *new(main_domains.Transaction)
	}
	return *main_domains.NewTransactionAllArgs(
		this.Id.Hex(),
		this.AccountId,
		this.OperationTypeId,
		this.Amount,
		this.CreatedDate.Time(),
		this.LastModifiedDate.Time(),
	)
}
