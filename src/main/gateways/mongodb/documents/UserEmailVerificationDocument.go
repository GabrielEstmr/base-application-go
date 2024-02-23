package main_gateways_mongodb_documents

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

type UserEmailVerificationDocument struct {
	Id               primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId           string              `json:"userId,omitempty" bson:"userId,omitempty"`
	VerificationCode string              `json:"verificationCode,omitempty" bson:"verificationCode,omitempty"`
	Scope            string              `json:"scope,omitempty" bson:"scope,omitempty"`
	EmailParams      EmailParamsDocument `json:"emailParams,omitempty" bson:"emailParams,omitempty"`
	Status           string              `json:"status,omitempty" bson:"status,omitempty"`
	CreatedDate      primitive.DateTime  `json:"createdDate,omitempty" bson:"createdDate"`
	LastModifiedDate primitive.DateTime  `json:"lastModifiedDate,omitempty" bson:"lastModifiedDate"`
}

func NewUserEmailVerificationDocument(
	emailVerification main_domains.UserEmailVerification,
) UserEmailVerificationDocument {
	oId, _ := primitive.ObjectIDFromHex(emailVerification.GetId())
	return UserEmailVerificationDocument{
		Id:               oId,
		UserId:           emailVerification.GetUserId(),
		VerificationCode: emailVerification.GetVerificationCode(),
		Scope:            emailVerification.GetScope().Name(),
		EmailParams:      *NewEmailParamsDocument(emailVerification.GetEmailParams()),
		Status:           emailVerification.GetStatus().Name(),
		CreatedDate:      primitive.NewDateTimeFromTime(emailVerification.GetCreatedDate()),
		LastModifiedDate: primitive.NewDateTimeFromTime(emailVerification.GetLastModifiedDate()),
	}
}

func (this UserEmailVerificationDocument) IsEmpty() bool {
	document := *new(UserEmailVerificationDocument)
	return reflect.DeepEqual(this, document)
}

func (this UserEmailVerificationDocument) ToDomain() main_domains.UserEmailVerification {
	if this.IsEmpty() {
		return *new(main_domains.UserEmailVerification)
	}
	return *main_domains.NewUserEmailVerificationAllArgs(
		this.Id.Hex(),
		this.UserId,
		this.VerificationCode,
		new(main_domains_enums.UserEmailVerificationScope).FromValue(this.Scope),
		this.EmailParams.ToDomain(),
		new(main_domains_enums.EmailVerificationStatus).FromValue(this.Status),
		this.CreatedDate.Time(),
		this.LastModifiedDate.Time(),
	)
}
