package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	uadmin.RenderHTML(w, r, "templates/index.html", map[string]interface{}{
		"Title": "MDRRMO System",
	})
}
