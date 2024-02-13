package views

import (
	"net/http"
	"strings"

	"github.com/mbdeguzman/safeguard/models"
	"github.com/uadmin/uadmin"
)

func ConsoleHandler(w http.ResponseWriter, r *http.Request) {
	session := uadmin.IsAuthenticated(r)
	if session == nil {
		http.Redirect(w, r, "/console/login/", http.StatusSeeOther)
		return
	}

	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/console/")
	page := strings.TrimSuffix(r.URL.Path, "/")

	// Define the context that will receive the page context from the handler
	context := map[string]interface{}{}

	switch page {
	case "dashboard":
		context = DashboardHandler(w, r)
	default:
		page = "dashboard"
		context = DashboardHandler(w, r)
	}

	enforcer := models.Enforcers{}
	uadmin.Get(&enforcer, "user_id = ?", session.UserID)
	position := ""

	context["User"] = session.User
	context["Position"] = position
	ConsoleRender(w, r, page, context)
}

func ConsoleRender(w http.ResponseWriter, r *http.Request, tpl string, context map[string]interface{}) {
	templateList := []string{}
	templateList = append(templateList, "./templates/console/base.html")

	path := "./templates/console/" + tpl + ".html"
	templateList = append(templateList, path)

	uadmin.RenderMultiHTML(w, r, templateList, context)
}
