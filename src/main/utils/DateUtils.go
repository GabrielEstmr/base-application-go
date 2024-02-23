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

func (this *DateUtils) AddDays(date time.Time, numberOfDays int32) time.Time {
	return date.Add(time.Duration(numberOfDays*24) * time.Hour)
}

func (this *DateUtils) GetDateAtStartOfTheDay(date time.Time, location time.Location) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, &location)
}

func (this *DateUtils) GetDateUTCAtStartOfTheDay(date time.Time) time.Time {
	return this.GetDateAtStartOfTheDay(date, *time.UTC)
}

func (this *DateUtils) GetDateAtEndOfTheDay(date time.Time, location time.Location) time.Time {
	tomorrowStartOfTheDay := this.GetDateAtStartOfTheDay(this.AddDays(date, 1), location)
	return tomorrowStartOfTheDay.Add(-1 * time.Nanosecond)
}

func (this *DateUtils) GetDateUTCAtEndOfTheDay(date time.Time) time.Time {
	return this.GetDateAtEndOfTheDay(date, *time.UTC)
}
