package main

import (
	"net/http"

	"github.com/mbdeguzman/safeguard/api"
	"github.com/mbdeguzman/safeguard/models"
	"github.com/mbdeguzman/safeguard/views"
	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.Register(
		models.AlertLevel{},
		models.Disasters{},
		models.EnforcerType{},
		models.Enforcers{},
		models.IncidentReport{},
		models.Protocols{},
	)

	http.HandleFunc("/login/", uadmin.Handler(views.LoginHandler))
	http.HandleFunc("/mobile/", uadmin.Handler(views.MobileHandler))

	http.HandleFunc("/", uadmin.Handler(views.IndexHandler))

	http.HandleFunc("/api/", uadmin.Handler(api.APIHandler))

	http.HandleFunc("/console/", uadmin.Handler(views.ConsoleHandler))
	http.HandleFunc("/console/login/", uadmin.Handler(views.ConsoleLoginHandler))
	http.HandleFunc("/logout/", uadmin.Handler(views.LogoutHandler))

	uadmin.Port = 9880
	uadmin.RootURL = "/uadmin/"

	uadmin.StartServer()
}
