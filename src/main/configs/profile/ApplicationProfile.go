package mainConfigsProfile

import (
	"errors"
	"strings"
)

type ApplicationProfile string

const invalidApplicationProfile = "Invalid Indicator description."

const localApplicationProfileDescription = "LOCAL"
const hmgApplicationProfileDescription = "HMG"
const prdApplicationProfileDescription = "PRD"

const (
	LOCAL ApplicationProfile = localApplicationProfileDescription
	HMG   ApplicationProfile = hmgApplicationProfileDescription
	PRD   ApplicationProfile = prdApplicationProfileDescription
)

func (s ApplicationProfile) GetDescription() string {
	switch s {
	case LOCAL:
		return localApplicationProfileDescription
	case HMG:
		return hmgApplicationProfileDescription
	case PRD:
		return prdApplicationProfileDescription
	}
	return "unknown"
}

func (s ApplicationProfile) GetLowerCaseDescription() string {
	return strings.ToLower(s.GetDescription())
}

func FindApplicationProfileByDescription(description string) (ApplicationProfile, error) {
	switch description {
	case localApplicationProfileDescription:
		return LOCAL, nil
	case hmgApplicationProfileDescription:
		return HMG, nil
	case prdApplicationProfileDescription:
		return PRD, nil
	}
	return "", errors.New(invalidApplicationProfile)
}
