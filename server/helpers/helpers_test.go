package helpers

import (
	"net/http/httptest"
	"testing"
)

type PaginationTest struct {
	p      *Pagination
	limit  uint
	offset uint
}

func TestPaginationToLimitOffset(t *testing.T) {
	tests := []PaginationTest{
		PaginationTest{
			p: &Pagination{
				Page: 1,
				Size: 1,
			},
			limit:  1,
			offset: 0,
		},
		PaginationTest{
			p: &Pagination{
				Page: 3,
				Size: 7,
			},
			limit:  7,
			offset: 14,
		},
		PaginationTest{
			p: &Pagination{
				Page: 0,
				Size: 0,
			},
			limit:  10,
			offset: 0,
		},
	}

	for _, test := range tests {
		if test.p.GetLimit() != test.limit {
			t.Errorf("Expect limit %v, got %v", test.limit, test.p.GetLimit())
		}

		if test.p.GetOffset() != test.offset {
			t.Errorf("Expect offset %v, got %v", test.offset, test.p.GetOffset())
		}
	}
}

func TestPaginationLoad(t *testing.T) {
	p := Pagination{}
	r := httptest.NewRequest("GET", "/tasks?page=3&page-size=14", nil)
	p.Load(r)
	if p.Page != 3 {
		t.Errorf("Expect page %v, but got %v", 3, p.Page)
	}

	if p.Size != 14 {
		t.Errorf("Expect page %v, but got %v", 14, p.Size)
	}
}
