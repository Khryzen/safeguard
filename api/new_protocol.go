package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/mbdeguzman/safeguard/models"
	"github.com/uadmin/uadmin"
)

func NewProtocol(w http.ResponseWriter, r *http.Request) {
	method := strings.ToUpper(r.Method)
	if method == "POST" {
		protocols := models.Protocols{}
		protocols.Name = r.FormValue("name")
		protocols.Code = r.FormValue("code")
		disaster_id, err := strconv.ParseUint(r.FormValue("disaster"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			return
		}
		protocols.DisastersID = uint(disaster_id)
		alert_id, err := strconv.ParseUint(r.FormValue("alert_level"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			return
		}
		protocols.AlertLevelID = uint(alert_id)
		protocols.Description = r.FormValue("description")

		active, err := strconv.ParseBool(r.FormValue("active"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			return
		}
		protocols.Active = active
		uadmin.Save(&protocols)

		http.Redirect(w, r, r.FormValue("redirect_url"), http.StatusSeeOther)
	}
}
