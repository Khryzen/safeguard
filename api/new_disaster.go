package api

import (
	"net/http"
	"strings"

	"github.com/mbdeguzman/safeguard/models"
)

func NewDisaster(w http.ResponseWriter, r *http.Request) {
	method := strings.ToUpper(r.Method)
	if method == "POST" {
		disaster := models.Disasters{}

		disaster.Name = r.FormValue("disaster")
		disaster.Description = r.FormValue("description")

		disaster.Save()
		http.Redirect(w, r, r.FormValue("redirect_url"), http.StatusSeeOther)
	}
}
