package views

import (
	"net/http"

	"github.com/mbdeguzman/safeguard/models"
	"github.com/uadmin/uadmin"
)

func MobileDashboardHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

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
