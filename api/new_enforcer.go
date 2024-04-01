package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/mbdeguzman/safeguard/models"
	"github.com/uadmin/uadmin"
)

func NewEnforcer(w http.ResponseWriter, r *http.Request) {
	method := strings.ToUpper(r.Method)
	if method == "POST" {
		enforcer_type := models.EnforcerType{}
		enforcer_type.Designation = r.FormValue("designation")
		enforcer_type.Role = r.FormValue("role")

		active, err := strconv.ParseBool(r.FormValue("active"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			return
		}
		enforcer_type.Active = active
		enforcer_type.Save()
		http.Redirect(w, r, r.FormValue("redirect_url"), http.StatusSeeOther)
	}
}
