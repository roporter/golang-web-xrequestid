package xrequestid2

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	id := uuid.NewV4().String()
	r.Header.Set("X-Request-Id", id)
	rw.Header().Set("X-Request-Id", id)
	next(rw, r)
}
