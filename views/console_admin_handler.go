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

	enforcer_types := []models.EnforcerType{}
	uadmin.Filter(&enforcer_types, "active = ?", true)

	context["Enforcers"] = enforcers
	context["EnforcerTypes"] = enforcer_types
	context["Title"] = "Enforcers"
	return context
}

func DisasterHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	disaster := []models.Disasters{}
	uadmin.All(&disaster)

	context["Disaster"] = disaster
	context["Title"] = "Disaster"
	return context
}

func AlertLevelHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}
	alerts := []models.AlertLevel{}
	uadmin.All(&alerts)

	context["Alerts"] = alerts
	context["Title"] = "Alert Levels"
	return context
}

func ProtocolsHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	disasters := []models.Disasters{}
	uadmin.All(&disasters)
	alerts := []models.AlertLevel{}
	uadmin.All(&alerts)
	protocols := []models.Protocols{}
	uadmin.All(&protocols)

	for i := range protocols {
		uadmin.Preload(&protocols[i])
	}

	context["Disasters"] = disasters
	context["AlertLevels"] = alerts
	context["Protocols"] = protocols
	context["Title"] = "Protocols"
	return context
}

func IncidentReportHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	incidents := []models.IncidentReport{}
	uadmin.All(&incidents)
	for i := range incidents {
		uadmin.Preload(&incidents[i])
		uadmin.Preload(&incidents[i].Captain)
	}

	disaster := []models.Disasters{}
	uadmin.All(&disaster)

	alert := []models.AlertLevel{}
	uadmin.All(&alert)

	enforcer := []models.Enforcers{}
	uadmin.Filter(&enforcer, "active = ?", true)
	for i := range enforcer {
		uadmin.Preload(&enforcer[i])
	}

	context["Alerts"] = alert
	context["Disaster"] = disaster
	context["Enforcers"] = enforcer
	context["Incidents"] = incidents
	context["Title"] = "Incident Reports"
	return context
}

func EnforcerTypeHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	enforcer_type := []models.EnforcerType{}
	uadmin.All(&enforcer_type)

	context["EnforcerType"] = enforcer_type
	context["Title"] = "Enforcer Types"
	return context
}

func SettingsHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	context["Title"] = "Settings"
	return context
}
