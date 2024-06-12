package api

import (
	"net/http"
	"strings"

	"github.com/mbdeguzman/safeguard/models"
	"github.com/uadmin/uadmin"
)

func NewSMSUser(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) == "POST" {
		sms := models.SMSSubscriber{}
		sms.FirstName = r.FormValue("firstname")
		sms.LastName = r.FormValue("lastname")
		sms.Address = r.FormValue("address")
		sms.MobileNumber = r.FormValue("mobile")
		sms.EmailAddress = r.FormValue("email")

		uadmin.Save(&sms)

		http.Redirect(w, r, "/register/", http.StatusSeeOther)
	}
}
