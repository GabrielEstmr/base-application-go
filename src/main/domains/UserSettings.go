package main_domains

import main_domains_enums "baseapplicationgo/main/domains/enums"

type UserSettings struct {
	language main_domains_enums.Language
	timeZone string
}
