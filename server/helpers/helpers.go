package helpers

import (
	"net/http"
	"strconv"
)

//Pagination struct to load and convert pagination to limit offset
type Pagination struct {
	Page uint
	Size uint
}

//Load pagination params from request
func (p *Pagination) Load(r *http.Request) {
	size := r.FormValue("page-size")
	page := r.FormValue("page")
	if t, err := strconv.ParseUint(size, 10, 16); err == nil {
		p.Size = uint(t)
	}

	if t, err := strconv.ParseUint(page, 10, 16); err == nil && t >= 1 {
		p.Page = uint(t)
	}
}

//GetLimit returns limit
func (p *Pagination) GetLimit() uint {
	if p.Size == 0 {
		return 10
	}

	return p.Size
}

//GetOffset returns offset
func (p *Pagination) GetOffset() uint {
	if p.Page == 0 {
		return 0
	}

	return p.GetLimit() * (p.Page - 1)
}
