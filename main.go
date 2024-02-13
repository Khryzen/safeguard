package main

import (
	"net/http"

	"github.com/mbdeguzman/safeguard/models"
	"github.com/mbdeguzman/safeguard/views"
	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.Register(
		models.AlertLevel{},
		models.Disasters{},
		models.Enforcers{},
		models.Protocols{},
	)

	http.HandleFunc("/", views.MobileHandler)
	http.HandleFunc("/login/", views.LoginHandler)

	http.HandleFunc("/console/login/", views.ConsoleLoginHandler)
	http.HandleFunc("/console/", views.ConsoleHandler)

	uadmin.Port = 8080
	uadmin.RootURL = "/uadmin/"

	uadmin.StartServer()
}
