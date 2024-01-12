package main_utils

const _STRING_UTILS_EMPTY_STRING = ""

type StringUtils struct {
}

func NewStringUtils() *StringUtils {
	return &StringUtils{}
}

func (this *StringUtils) IsEmpty(value string) bool {
	return value == _STRING_UTILS_EMPTY_STRING
}
