package main_configs_messages_resources

import (
	"errors"
	"strings"
)

type LanguageProfiles string

const _MSG_ERROR_INVALID_PROFILE = "Invalid language profile"

const defaultLangProfDescription = "default"
const usEnLangProfDescription = "en-US"
const ptBrLangProfDescription = "pt-BR"

const defaultLangProfName = "DEFAULT"
const usEnLangProfName = "EN_US"
const ptBrLangProfName = "PT_BR"

const (
	DEFAULT LanguageProfiles = defaultLangProfDescription
	EN_US   LanguageProfiles = usEnLangProfDescription
	PT_BR   LanguageProfiles = ptBrLangProfDescription
)

func (s LanguageProfiles) LanguageProfileName() string {
	switch s {
	case DEFAULT:
		return defaultLangProfName
	case EN_US:
		return usEnLangProfName
	case PT_BR:
		return ptBrLangProfName
	}
	return "unknown"
}

func (s LanguageProfiles) GetLanguageProfileDescription() string {
	switch s {
	case DEFAULT:
		return defaultLangProfDescription
	case EN_US:
		return usEnLangProfDescription
	case PT_BR:
		return ptBrLangProfDescription
	}
	return "unknown"
}

func (s LanguageProfiles) GetLowerCaseLanguageProfileName() string {
	return strings.ToLower(s.LanguageProfileName())
}

func GetLanguageProfileValues() []LanguageProfiles {
	return []LanguageProfiles{DEFAULT, EN_US, PT_BR}
}

func FindLanguageProfileByDescription(description string) (LanguageProfiles, error) {
	switch description {
	case defaultLangProfDescription:
		return DEFAULT, nil
	case usEnLangProfDescription:
		return EN_US, nil
	case ptBrLangProfDescription:
		return PT_BR, nil
	}
	return "", errors.New(_MSG_ERROR_INVALID_PROFILE)
}
