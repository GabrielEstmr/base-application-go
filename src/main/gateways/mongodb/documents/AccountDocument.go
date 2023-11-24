package main_gateways_mongodb_documents

import (
	domains "baseapplicationgo/main/domains"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountDocument struct {
	Id               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserId           string             `json:"userId,omitempty"`
	CreatedDate      primitive.DateTime `json:"createdDate,omitempty"`
	LastModifiedDate primitive.DateTime `json:"lastModifiedDate,omitempty"`
}

func NewAccountDocument(account domains.Account) AccountDocument {
	// TODO: construtor com id e datas
	return AccountDocument{
		UserId: account.UserId,
	}
}

func (thisDoc *AccountDocument) ToDomain() domains.Account {
	return domains.Account{
		Id:               thisDoc.Id.Hex(),
		UserId:           thisDoc.UserId,
		CreatedDate:      thisDoc.CreatedDate.Time(),
		LastModifiedDate: thisDoc.LastModifiedDate.Time(),
	}
}
