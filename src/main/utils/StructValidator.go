package main_utils

import (
	"fmt"
	"reflect"
)

type StructValidator struct {
	data interface{}
	val  reflect.Value
}

func NewValidatorTest(data interface{}) *StructValidator {
	return &StructValidator{
		data: data,
		val:  reflect.ValueOf(data).Elem(),
	}
}

func (this *StructValidator) getFieldName(fieldIdx int) string {
	return this.val.Type().Field(fieldIdx).Name
}

func (this *StructValidator) StringRequired(fieldIdx int, fieldValue string, messages []string) []string {
	name := this.getFieldName(fieldIdx)
	if len(fieldValue) == 0 {
		messages = append(messages, fmt.Sprintf("%s: Field must not be empty", name))
	}
	return messages
}

func (this *StructValidator) StringLen(fieldIdx int, fieldValue string, messages []string, min int, max int) []string {
	name := this.getFieldName(fieldIdx)
	if len(fieldValue) < min || len(fieldValue) > max {
		messages = append(messages, fmt.Sprintf("%s: Field over lenght", name))
	}
	return messages
}

func (this *StructValidator) IntBetween(fieldIdx int, fieldValue int, messages []string, min int, max int) []string {
	name := this.getFieldName(fieldIdx)
	if fieldValue < min || fieldValue > max {
		messages = append(messages, fmt.Sprintf("%s: Field over lenght", name))
	}
	return messages
}
