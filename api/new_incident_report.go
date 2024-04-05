package api

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/mbdeguzman/safeguard/models"
	"github.com/uadmin/uadmin"
)

func NewIncidentReport(w http.ResponseWriter, r *http.Request) {
	r.Method = strings.ToUpper(r.Method)

	if r.Method == "POST" {
		incident_date, err := time.Parse("2006-01-02T15:04", r.FormValue("incident_date"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			return
		}

		incident := models.IncidentReport{}
		incident.IncidentDate = incident_date

		disaster, err := strconv.ParseUint(r.FormValue("disaster"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			return
		}
		incident.DisastersID = uint(disaster)

		alert, err := strconv.ParseUint(r.FormValue("alert_level"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			return
		}
		incident.AlertLevelID = uint(alert)

		captain, err := strconv.ParseUint(r.FormValue("captain"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			return
		}
		incident.CaptainID = uint(captain)
		incident.Remarks = r.FormValue("remarks")
		incident.Location = r.FormValue("location")

		casualties, err := strconv.ParseInt(r.FormValue("casualties"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			return
		}
		incident.Casualties = int(casualties)

		responded, err := strconv.ParseBool(r.FormValue("responded"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			return
		}
		incident.Responded = responded
		incident.Save()

		http.Redirect(w, r, "/console/incidents/", http.StatusSeeOther)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Method not allowed",
		})
		return
	}
}
