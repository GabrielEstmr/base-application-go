package main_gateways_mongodb_documents

import (
	main_domains "baseapplicationgo/main/domains"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDocument struct {
	Id               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name             string             `json:"name,omitempty" bson:"name,omitempty"`
	DocumentNumber   string             `json:"documentNumber,omitempty" bson:"documentNumber,omitempty"`
	Birthday         primitive.DateTime `json:"birthday,omitempty" bson:"birthday,omitempty"`
	CreatedDate      primitive.DateTime `json:"createdDate,omitempty" bson:"createdDate"`
	LastModifiedDate primitive.DateTime `json:"lastModifiedDate,omitempty" bson:"lastModifiedDate"`
}

func NewUserDocument(user main_domains.User) UserDocument {
	return UserDocument{
		Name:           user.Name,
		DocumentNumber: user.DocumentNumber,
		Birthday:       primitive.NewDateTimeFromTime(user.Birthday),
	}
}

func (this *UserDocument) IsEmpty() bool {
	return *this == UserDocument{}
}

func (this *UserDocument) ToDomain() main_domains.User {
	if (*this == UserDocument{}) {
		return main_domains.User{}
	}
	return main_domains.User{
		Id:               this.Id.Hex(),
		Name:             this.Name,
		DocumentNumber:   this.DocumentNumber,
		Birthday:         this.Birthday.Time(),
		CreatedDate:      this.CreatedDate.Time(),
		LastModifiedDate: this.LastModifiedDate.Time(),
	}
}
