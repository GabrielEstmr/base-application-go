package main_gateways_mongodb_documents

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

type EmailDocument struct {
	Id               primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	EventId          string              `json:"eventId,omitempty" bson:"eventId,omitempty"`
	EmailParams      EmailParamsDocument `json:"emailParams,omitempty" bson:"emailParams,omitempty"`
	Status           string              `json:"status,omitempty" bson:"status,omitempty"`
	ErrorMsg         string              `json:"errorMsg,omitempty" bson:"errorMsg,omitempty"`
	CreatedDate      primitive.DateTime  `json:"createdDate,omitempty" bson:"createdDate"`
	LastModifiedDate primitive.DateTime  `json:"lastModifiedDate,omitempty" bson:"lastModifiedDate"`
}

func NewEmailDocument(email main_domains.Email) EmailDocument {
	oId, _ := primitive.ObjectIDFromHex(email.GetId())
	return EmailDocument{
		Id:               oId,
		EventId:          email.GetEventId(),
		EmailParams:      *NewEmailParamsDocument(email.GetEmailParams()),
		Status:           main_domains_enums.GetEmailStatusDescription(email.GetStatus()),
		ErrorMsg:         email.GetErrorMsg(),
		CreatedDate:      primitive.NewDateTimeFromTime(email.GetCreatedDate()),
		LastModifiedDate: primitive.NewDateTimeFromTime(email.GetLastModifiedDate()),
	}
}

func (this EmailDocument) IsEmpty() bool {
	document := *new(EmailDocument)
	return reflect.DeepEqual(this, document)
}

func (this EmailDocument) ToDomain() main_domains.Email {
	if this.IsEmpty() {
		return *new(main_domains.Email)
	}
	return *main_domains.NewEmailAllArgs(
		this.Id.Hex(),
		this.EventId,
		this.EmailParams.ToDomain(),
		main_domains_enums.GetEmailStatusFromDescription(this.Status),
		this.ErrorMsg,
		this.CreatedDate.Time(),
		this.LastModifiedDate.Time(),
	)
}
