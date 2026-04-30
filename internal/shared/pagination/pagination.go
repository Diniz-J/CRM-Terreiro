package pagination

const (
	DefaultPage     = 1
	DefaultPageSize = 20
	MaxPageSize     = 100
)

type Params struct {
	Page     int
	PageSize int
}

type Result[T any] struct {
	Data       []T        `json:"data"`
	Pagination ResultMeta `json:"pagination"`
}

type ResultMeta struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

func Normalize(page, pageSize int) Params {
	if page < 1 {
		page = DefaultPage
	}
	if pageSize < 1 {
		pageSize = DefaultPageSize
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}
	return Params{Page: page, PageSize: pageSize}
}

func NewResult[T any](data []T, p Params, total int) Result[T] {
	totalPages := total / p.PageSize
	if total%p.PageSize != 0 {
		totalPages++
	}
	if data == nil {
		data = []T{}
	}
	return Result[T]{
		Data: data,
		Pagination: ResultMeta{
			Page:       p.Page,
			PageSize:   p.PageSize,
			Total:      total,
			TotalPages: totalPages,
		},
	}
}

func Offset(p Params) int {
	return (p.Page - 1) * p.PageSize
}
