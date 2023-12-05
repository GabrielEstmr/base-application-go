package main_utils

import (
	"strings"
)

type PageableUtils struct {
}

func NewPageableUtils() *PageableUtils {
	return &PageableUtils{}
}

func (this *PageableUtils) BuildPageableSortToMap(sort []string) map[string]int {
	sortMetadata := make(map[string]int)
	for _, value := range sort {
		tmp := strings.Split(value, ",")
		if len(tmp) == 2 && len(strings.TrimSpace(tmp[0])) > 0 && len(strings.TrimSpace(tmp[1])) > 0 {
			tmp[0] = strings.TrimSpace(tmp[0])
			tmp[1] = strings.ToUpper(strings.TrimSpace(tmp[1]))
			if tmp[1] == "ASC" {
				sortMetadata[tmp[0]] = 1
			}
			if tmp[1] == "DESC" {
				sortMetadata[tmp[0]] = -1
			}
		}
	}
	return sortMetadata
}
