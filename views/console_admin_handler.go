package views

import (
	"net/http"

	"github.com/mbdeguzman/safeguard/models"
	"github.com/uadmin/uadmin"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}
	enforcers := models.Enforcers{}
	active_enforcer := uadmin.Count(&enforcers, "active = ?", true)

	sms := models.SMSSubscriber{}
	sms_subscriber := uadmin.Count(&sms, "id > 0")

	incident := models.IncidentReport{}
	incident_count := uadmin.Count(&incident, "id > 0")

	items := models.Item{}
	item_count := uadmin.Count(&items, "id > 0")

	transactions := models.Transaction{}
	transaction_count := uadmin.Count(&transactions, "id > 0")

	stocks := models.InventoryLine{}
	stocks_count := uadmin.Count(&stocks, "remaining = ?", 0)

	context["Enforcers"] = active_enforcer
	context["SMS"] = sms_subscriber
	context["IncidentReport"] = incident_count
	context["Items"] = item_count
	context["Transactions"] = transaction_count
	context["Stocks"] = stocks_count
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

	type Enforcer struct {
		Enforcer  models.Enforcers
		LastLogin string
	}

	e := []Enforcer{}
	enforcers := []models.Enforcers{}
	uadmin.Filter(&enforcers, "active = ?", true)
	for i := range enforcers {
		uadmin.Preload(&enforcers[i])
		uadmin.Trail(uadmin.DEBUG, "%#v", enforcers[i].User.LastLogin)
		if enforcers[i].User.LastLogin == nil {
			e = append(e, Enforcer{
				Enforcer:  enforcers[i],
				LastLogin: "No login records yet.",
			})
		} else {
			e = append(e, Enforcer{
				Enforcer:  enforcers[i],
				LastLogin: enforcers[i].User.LastLogin.Format("January 2, 2006 3:04 PM"),
			})
		}
	}

	context["Enforcers"] = e
	context["Title"] = "Settings"
	return context
}

func ItemsHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	line := []models.InventoryLine{}
	uadmin.All(&line)
	for i := range line {
		uadmin.Preload(&line[i])
	}

	context["Items"] = line
	context["Title"] = "Items - Inventory"
	return context
}

func TransactionHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	transactions := []models.Transaction{}
	uadmin.All(&transactions)
	for i := range transactions {
		uadmin.Preload(&transactions[i])
		uadmin.Preload(&transactions[i].Enforcer)
	}
	items := []models.Item{}
	uadmin.All(&items)

	session := uadmin.IsAuthenticated(r)
	enforcer := models.Enforcers{}
	uadmin.Get(&enforcer, "user_id = ?", session.UserID)
	uadmin.Preload(&enforcer)
	uadmin.Preload(&enforcer.User)

	// context["User"] = session.User
	context["Enforcer"] = enforcer
	context["Items"] = items
	context["Transactions"] = transactions
	context["Title"] = "Items - Inventory"
	return context
}

func NotificationsHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	context["Title"] = "Notifications"
	return context
}
