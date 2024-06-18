package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func SurveyHandler(w http.ResponseWriter, r *http.Request) {
	uadmin.RenderHTML(w, r, "templates/survey.html", nil)
}
