package views

import (
	"net/http"
	"time"

	"github.com/mbdeguzman/safeguard/models"
	"github.com/uadmin/uadmin"
)

func MobileDashboardHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}
	enforcers := models.Enforcers{}
	active_enforcer := uadmin.Count(&enforcers, "active = ?", true)

	incident := models.IncidentReport{}
	incident_count := uadmin.Count(&incident, "id > 0")

	currentTime := time.Now()
	oneDayAgo := currentTime.Add(-24 * time.Hour)
	disaster := models.IncidentReport{}
	disaster_count := uadmin.Count(&disaster, "incident_date > ? AND incident_date < ?", oneDayAgo, currentTime)

	context["Enforcers"] = active_enforcer
	context["IncidentReport"] = incident_count
	context["Disasters"] = disaster_count
	context["Title"] = "Dashboard"
	return context
}

func EnforcerMobileHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	enforcers := []models.Enforcers{}
	uadmin.Filter(&enforcers, "active = ?", true)

	for i := range enforcers {
		uadmin.Preload(&enforcers[i])
		uadmin.Preload(&enforcers[i].User)
	}

	context["Enforcers"] = enforcers
	return context
}

func HelpMobileHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	return context
}

func AboutMobileHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	return context
}
