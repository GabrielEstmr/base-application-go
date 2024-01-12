package main_gateways_ws_commonsresources

import main_domains "baseapplicationgo/main/domains"

type PageResponse struct {
	Content       any   `json:"content"`
	Page          int64 `json:"page"`
	Size          int64 `json:"size"`
	TotalElements int64 `json:"totalElements"`
	TotalPages    int64 `json:"totalPages"`
}

func NewPageResponse(content any, page int64, size int64, totalElements int64, totalPages int64) *PageResponse {
	return &PageResponse{Content: content, Page: page, Size: size, TotalElements: totalElements, TotalPages: totalPages}
}

func NewPageResponseFromPage(page main_domains.Page) *PageResponse {
	return &PageResponse{
		Content:       page.GetContent(),
		Page:          page.GetPage(),
		Size:          page.GetSize(),
		TotalElements: page.GetTotalElements(),
		TotalPages:    page.GetTotalPages(),
	}
}
