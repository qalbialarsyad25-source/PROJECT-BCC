package model

type Pagination struct {
	Lembar int `json:"lembar"`
	Limit  int `json:"limit"`
}

type PaginationResponse[T any] struct {
	Data       []T        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

func (p *Pagination) Check() {
	if p.Lembar < 1 {
		p.Lembar = 10
	}

	if p.Limit < 1 || p.Limit > 50 {
		p.Limit = 10
	}
}

func (p *Pagination) Offset() int {
	return (p.Lembar - 1) * p.Limit
}
