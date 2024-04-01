package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/mbdeguzman/safeguard/models"
	"github.com/uadmin/uadmin"
)

func NewAlertLevel(w http.ResponseWriter, r *http.Request) {
	method := strings.ToUpper(r.Method)
	if method == "POST" {
		alert := models.AlertLevel{}
		alert.Name = r.FormValue("name")
		level, err := strconv.ParseInt(r.FormValue("level"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			return
		}
		alert.Level = int(level)
		alert.Save()
		http.Redirect(w, r, r.FormValue("redirect_url"), http.StatusSeeOther)
	}
}
