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

		// Inventory Models
		models.Item{},
		models.Transaction{},
		models.InventoryLine{},

		// SMS Model
		models.SMSSubscriber{},
	)

	category := uadmin.SettingCategory{}

	if uadmin.Count(&category, "name = ?", "SMS") == 0 {
		settings := uadmin.Setting{}
		category.Name = "SMS"
		uadmin.Save(&category)

		uadmin.Get(&category, "name = ?", "SMS")
		settings.Category = category
		settings.DataType = settings.DataType.String()
		settings.Name = "API Key"
		settings.Value = "71174e4de05abb5b79ac73a25d76fe9b"
		settings.Code = "SMSAPI"
		uadmin.Save(&settings)
	}

	http.HandleFunc("/login/", uadmin.Handler(views.LoginHandler))
	http.HandleFunc("/mobile/", uadmin.Handler(views.MobileHandler))

	http.HandleFunc("/", uadmin.Handler(views.IndexHandler))

	http.HandleFunc("/api/", uadmin.Handler(api.APIHandler))

	http.HandleFunc("/console/", uadmin.Handler(views.ConsoleHandler))
	http.HandleFunc("/console/login/", uadmin.Handler(views.ConsoleLoginHandler))
	http.HandleFunc("/register/", uadmin.Handler(views.SMSRegister))
	http.HandleFunc("/logout/", uadmin.Handler(views.LogoutHandler))

	uadmin.Port = 9880
	uadmin.RootURL = "/uadmin/"

	uadmin.StartServer()
}
