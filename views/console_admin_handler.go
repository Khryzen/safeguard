package views

import (
	"net/http"

	"github.com/mbdeguzman/safeguard/models"
	"github.com/uadmin/uadmin"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	context["Title"] = "Dashboard"
	return context
}

func EnforcerHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	enforcers := []models.Enforcers{}
	uadmin.Filter(&enforcers, "active = ?", true)
	for i := range enforcers {
		uadmin.Preload(&enforcers[i])
	}

	context["Enforcers"] = enforcers
	context["Title"] = "Enforcers"
	return context
}

func DisasterHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	context["Title"] = "Disaster"
	return context
}

func AlertLevelHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	context["Title"] = "Alert Levels"
	return context
}

func ProtocolsHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	context["Title"] = "Protocols"
	return context
}

func IncidentReportHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	context["Title"] = "Incident Reports"
	return context
}

func EnforcerTypeHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	context["Title"] = "Enforcer Types"
	return context
}

func SettingsHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	context["Title"] = "Settings"
	return context
}
