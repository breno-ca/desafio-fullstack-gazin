package pagination

type Pagination struct {
	Page  int
	Limit int
	Total int
}

func NewPagination(Page, Limit int) Pagination {
	return Pagination{
		Page:  Page,
		Limit: Limit,
	}
}

func (pg *Pagination) Offset() int {
	return (pg.Page - 1) * pg.Limit
}

func (pg *Pagination) LastPage() int {
	return (pg.Total + pg.Limit - 1) / pg.Limit
}
