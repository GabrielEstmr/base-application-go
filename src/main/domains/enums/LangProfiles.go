package main_domains_enums

type LanguageProfiles string

const (
	LANGUAGE_PROFILE_DEFAULT LanguageProfiles = "DEFAULT"
	LANGUAGE_PROFILE_EN_US   LanguageProfiles = "EN_US"
	LANGUAGE_PROFILE_PT_BR   LanguageProfiles = "PT_BR"
)

var languageProfilesEnum = map[LanguageProfiles]LanguageProfiles{
	LANGUAGE_PROFILE_DEFAULT: LANGUAGE_PROFILE_DEFAULT,
	LANGUAGE_PROFILE_EN_US:   LANGUAGE_PROFILE_EN_US,
	LANGUAGE_PROFILE_PT_BR:   LANGUAGE_PROFILE_PT_BR,
}

func (this LanguageProfiles) ExistsLanguageProfiles(value LanguageProfiles) bool {
	_, exists := languageProfilesEnum[value]
	return exists
}

func (this LanguageProfiles) Name() string {
	switch this {
	case LANGUAGE_PROFILE_DEFAULT:
		return string(LANGUAGE_PROFILE_DEFAULT)
	case LANGUAGE_PROFILE_EN_US:
		return string(LANGUAGE_PROFILE_EN_US)
	case LANGUAGE_PROFILE_PT_BR:
		return string(LANGUAGE_PROFILE_PT_BR)
	}
	return "unknown"
}
