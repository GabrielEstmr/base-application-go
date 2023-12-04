package main_gateways_ws_commons

import main_domains "baseapplicationgo/main/domains"

type PageResponse struct {
	Content       any   `json:"Content"`
	Page          int64 `json:"Page"`
	Size          int64 `json:"Size"`
	TotalElements int64 `json:"TotalElements"`
	TotalPages    int64 `json:"TotalPages"`
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
