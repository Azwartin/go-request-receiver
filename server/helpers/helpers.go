package helpers

import (
	"net/http"
	"strconv"
)

//Pagination struct to load and convert pagination to limit offset
type Pagination struct {
	Page int
	Size int
}

//Load pagination params from request
func (p *Pagination) Load(r *http.Request) {
	size := r.FormValue("page-size")
	page := r.FormValue("page")
	if t, err := strconv.Atoi(size); err == nil {
		p.Size = t
	}

	if t, err := strconv.Atoi(page); err == nil && t >= 1 {
		p.Page = t
	}
}

//GetLimit returns limit
func (p *Pagination) GetLimit() int {
	if p.Size == 0 {
		return 10
	}

	return p.Size
}

//GetOffset returns offset
func (p *Pagination) GetOffset() int {
	if p.Page == 0 {
		return 0
	}

	return p.GetLimit() * (p.Page - 1)
}
