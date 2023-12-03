package main_domains

type Page struct {
	content       []any
	page          int
	size          int
	totalElements int
	totalPages    int
}

func NewPage(
	content []any,
	page int, size int,
	totalElements int,
	totalPages int,
) *Page {
	return &Page{
		content: content,
		page:    page, size: size,
		totalElements: totalElements,
		totalPages:    totalPages,
	}
}

func (p Page) getContent() []any {
	return p.content
}

func (p Page) getPage() int {
	return p.page
}

func (p Page) getSize() int {
	return p.size
}

func (p Page) getTotalElements() int {
	return p.totalElements
}

func (p Page) getTotalPages() int {
	return p.totalPages
}
