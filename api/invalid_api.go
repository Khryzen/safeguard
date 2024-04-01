package api

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func InvalidAPI(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	uadmin.ReturnJSON(w, r, map[string]interface{}{
		"msg": "Not found",
	})
}
