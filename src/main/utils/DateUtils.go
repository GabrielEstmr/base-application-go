package main_utils

import (
	"time"
)

type DateUtils struct {
}

func NewDateUtils() *DateUtils {
	return &DateUtils{}
}

func (this *DateUtils) GetMinValue(arr []time.Time) time.Time {
	minValue := arr[0]
	for _, v := range arr {
		if v.Before(minValue) {
			minValue = v
		}
	}
	return minValue
}

func (this *DateUtils) GetMaxValue(arr []time.Time) time.Time {
	minValue := arr[0]
	for _, v := range arr {
		if v.After(minValue) {
			minValue = v
		}
	}
	return minValue
}
