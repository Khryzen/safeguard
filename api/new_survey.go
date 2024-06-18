package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/mbdeguzman/safeguard/models"
	"github.com/uadmin/uadmin"
)

func NewSurvey(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) == "POST" {
		survey := models.Survey{}
		survey.Usage = StringToInt(r.FormValue("usage"))
		survey.Overall = StringToInt(r.FormValue("overall"))
		survey.Navigate = StringToInt(r.FormValue("navigate"))
		survey.RTMonitoring = StringToInt(r.FormValue("rtMonitoring"))
		survey.SMS = StringToInt(r.FormValue("sms"))
		survey.IncidentReporting = StringToInt(r.FormValue("reporting"))
		survey.Chatbot = StringToInt(r.FormValue("chatbot"))
		survey.TrainingModules = StringToInt(r.FormValue("trainingModules"))
		survey.Comments = r.FormValue("comments")
		uadmin.Save(&survey)
		http.Redirect(w, r, "/mobile/", http.StatusSeeOther)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"status": "error",
			"msg":    "method not allowed",
		})
	}
}

func StringToInt(str string) int {
	x, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		uadmin.Trail(uadmin.ERROR, "Unable to parse %s to int. %v", str, err.Error())
		return 0
	}
	return int(x)
}
