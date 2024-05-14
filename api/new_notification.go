package api

import (
	"net/http"

	"github.com/mbdeguzman/safeguard/models"
	"github.com/uadmin/uadmin"
)

func NewNotification(w http.ResponseWriter, r *http.Request) {
	notif := models.Notification{}
	notif.Name = r.FormValue("name")
	notif.CellNumber = r.FormValue("cellnumber")
	notif.Emergency = r.FormValue("emergency")
	uadmin.Save(&notif)
}
