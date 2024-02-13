package views

import (
	"net/http"
	"strings"

	"github.com/mbdeguzman/safeguard/models"
	"github.com/uadmin/uadmin"
)

func MobileHandler(w http.ResponseWriter, r *http.Request) {
	session := uadmin.IsAuthenticated(r)
	if session == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	page := strings.TrimSuffix(r.URL.Path, "/")

	// Define the context that will receive the page context from the handler
	context := map[string]interface{}{}

	switch page {
	case "dashboard":
		context = MobileDashboardHandler(w, r)
	default:
		page = "dashboard"
		context = MobileDashboardHandler(w, r)
	}

	enforcer := models.Enforcers{}
	uadmin.Get(&enforcer, "user_id = ?", session.UserID)
	position := ""

	context["User"] = session.User
	context["Position"] = position
	MobileRender(w, r, page, context)
}

func MobileRender(w http.ResponseWriter, r *http.Request, tpl string, context map[string]interface{}) {
	templateList := []string{}
	templateList = append(templateList, "./templates/mobile/base.html")

	path := "./templates/mobile/" + tpl + ".html"
	templateList = append(templateList, path)

	uadmin.RenderMultiHTML(w, r, templateList, context)
}
