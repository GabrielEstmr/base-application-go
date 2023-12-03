package main_domains

type Pageable struct {
	page int
	size int
	sort string
}

func NewPageable(
	page int,
	size int,
	sort string) *Pageable {
	return &Pageable{
		page: page,
		size: size,
		sort: sort,
	}
}

func (p *Pageable) getPage() int {
	return p.page
}

func (p *Pageable) getSize() int {
	return p.size
}

func (p *Pageable) getSort() string {
	return p.sort
}
