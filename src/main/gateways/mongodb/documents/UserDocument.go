package main_gateways_mongodb_documents

import (
	domains "baseapplicationgo/main/domains"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDocument struct {
	Id               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name             string             `json:"name,omitempty"`
	DocumentNumber   string             `json:"documentNumber,omitempty"`
	Birthday         primitive.DateTime `json:"birthday,omitempty"`
	CreatedDate      primitive.DateTime `json:"createdDate,omitempty"`
	LastModifiedDate primitive.DateTime `json:"lastModifiedDate,omitempty"`
}

func NewUserDocument(user domains.User) UserDocument {
	return UserDocument{
		Name:           user.Name,
		DocumentNumber: user.DocumentNumber,
		Birthday:       primitive.NewDateTimeFromTime(user.Birthday),
	}
}

func (thisDoc *UserDocument) ToDomain() domains.User {
	return domains.User{
		Id:               thisDoc.Id.Hex(),
		Name:             thisDoc.Name,
		DocumentNumber:   thisDoc.DocumentNumber,
		Birthday:         thisDoc.Birthday.Time(),
		CreatedDate:      thisDoc.CreatedDate.Time(),
		LastModifiedDate: thisDoc.LastModifiedDate.Time(),
	}
}
