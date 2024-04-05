package api

import (
	"net/http"
	"strings"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = strings.TrimPrefix(path, "/api/")
	path = strings.TrimSuffix(path, "/")

	switch path {
	case "new_enforcer":
		NewEnforcer(w, r)
	case "new_enforcer_type":
		NewEnforcerType(w, r)
	case "new_disaster":
		NewDisaster(w, r)
	case "new_alert_level":
		NewAlertLevel(w, r)
	case "new_protocol":
		NewProtocol(w, r)
	case "new_incident_report":
		NewIncidentReport(w, r)
	default:
		InvalidAPI(w, r)
	}
}
