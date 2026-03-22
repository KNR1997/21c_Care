package utils

import (
	"net/http"
	"strconv"
)

type Pagination struct {
	Limit int
	Page  int
}

func ParsePagination(r *http.Request) Pagination {
	limit := 10
	page := 1

	q := r.URL.Query()

	if l := q.Get("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil && v > 0 {
			limit = v
		}
	}

	if p := q.Get("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}

	// Optional safety limits
	if limit > 100 {
		limit = 100
	}

	return Pagination{
		Limit: limit,
		Page:  page,
	}
}
