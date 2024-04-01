package api

import (
	"net/http"
	"strings"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = strings.TrimPrefix(path, "/api/")
	path = strings.TrimSuffix(path, "/")

	switch path {
	case "new_enforcer_type":
		NewEnforcer(w, r)
	default:
		InvalidAPI(w, r)
	}
}
