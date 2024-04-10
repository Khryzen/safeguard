package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func SMSRegister(w http.ResponseWriter, r *http.Request) {
	uadmin.RenderHTML(w, r, "templates/register.html", nil)
}
