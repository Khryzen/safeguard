package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

func MobileHandler(w http.ResponseWriter, r *http.Request) {
	session := uadmin.IsAuthenticated(r)
	if session == nil {
		http.Redirect(w, r, "/login/", http.StatusSeeOther)
		return
	}

	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/mobile/")
	page := strings.TrimSuffix(r.URL.Path, "/")

	// Define the context that will receive the page context from the handler
	context := map[string]interface{}{}

	switch page {
	case "dashboard":
		context = MobileDashboardHandler(w, r)
	case "enforcers":
		context = EnforcerMobileHandler(w, r)
	default:
		page = "dashboard"
		context = MobileDashboardHandler(w, r)
	}
	MobileRender(w, r, page, context)
}

func MobileRender(w http.ResponseWriter, r *http.Request, tpl string, context map[string]interface{}) {
	templateList := []string{}
	templateList = append(templateList, "./templates/mobile/base.html")

	path := "./templates/mobile/" + tpl + ".html"
	templateList = append(templateList, path)

	uadmin.RenderMultiHTML(w, r, templateList, context)
}
