package routes

import (
	"fmt"
	"strings"
)

const (
	Requests              = "/requests"
	RequestPathParam      = "/requests/:id"
	RequestFieldPathParam = "/requests/:id/:field"
)

func RequestPath(id int64) string {
	var b strings.Builder
	fmt.Fprintf(&b, "%s/%d", Requests, id)
	return b.String()
}

func RequestFieldPath(id int64, field string) string {
	var b strings.Builder
	fmt.Fprintf(&b, "%s/%d/%s", Requests, id, field)
	return b.String()
}
