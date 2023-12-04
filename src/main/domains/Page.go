package main_domains

type Page struct {
	content       []Content
	page          int64
	size          int64
	totalElements int64
	totalPages    int64
}

func (this *Page) GetContent() []Content {
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

type Content struct {
	obj any
}

func (c Content) GetObj() any {
	return c.obj
}

func NewContent(obj any) *Content {
	return &Content{obj: obj}
}

func NewPage(
	content []Content,
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
