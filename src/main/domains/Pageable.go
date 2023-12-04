package main_domains

type Pageable struct {
	page int64
	size int64
	sort map[string]int
}

func NewPageable(
	page int64,
	size int64,
	sort map[string]int) *Pageable {
	return &Pageable{
		page: page,
		size: size,
		sort: sort,
	}
}

func (this *Pageable) GetPage() int64 {
	return this.page
}

func (this *Pageable) GetSize() int64 {
	return this.size
}

func (this *Pageable) GetSort() map[string]int {
	return this.sort
}

func (this *Pageable) SetSort(sort map[string]int) {
	this.sort = sort
}

func (this *Pageable) HasEmptySort() bool {
	return len(this.sort) == 0
}
