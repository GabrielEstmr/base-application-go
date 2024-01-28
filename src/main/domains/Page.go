package main_domains

import "math"

type Page struct {
	content       []any
	page          int64
	size          int64
	totalElements int64
	totalPages    int64
}

func (this *Page) GetContent() []any {
	return this.content
}

func (this *Page) GetPage() int64 {
	return this.page
}

func (this *Page) GetSize() int64 {
	return this.size
}

func (this *Page) GetTotalElements() int64 {
	return this.totalElements
}

func (this *Page) GetTotalPages() int64 {
	return this.totalPages
}

func NewPageAllArgs(
	content []any,
	page int64,
	size int64,
	totalElements int64,
	totalPages int64,
) *Page {
	return &Page{
		content: content,
		page:    page, size: size,
		totalElements: totalElements,
		totalPages:    totalPages,
	}
}

func NewPage(
	content []any,
	page int64,
	size int64,
	totalElements int64,
) *Page {
	return &Page{
		content: content,
		page:    page, size: size,
		totalElements: totalElements,
		totalPages:    buildTotalPages(size, totalElements),
	}
}

func NewPageFromContentAndPage(
	content []any,
	page Page,
) *Page {
	return NewPageAllArgs(
		content,
		page.GetPage(),
		page.GetSize(),
		page.GetTotalElements(),
		page.GetTotalPages(),
	)
}

func buildTotalPages(size int64,
	totalElements int64) int64 {
	var result float64
	result = float64(totalElements) / float64(size)
	return int64(math.Ceil(result))
}
