package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Check if there is an existing session
	session := uadmin.IsAuthenticated(r)
	if session != nil {
		http.Redirect(w, r, "/dashboard/", http.StatusSeeOther)
		return
	}
	// Check if request method is post
	// If method is post, process login and create session
	if r.Method == "POST" {
		// Get the values from the forms
		username := r.FormValue("username")
		password := r.FormValue("password")
		user := uadmin.User{}
		uadmin.Get(&user, "username = ?", username)
		session = user.Login(password, "")

		if session != nil {
			http.SetCookie(w, &http.Cookie{
				Path:  "/",
				Name:  "session",
				Value: session.Key,
			})
		}
		if session != nil {
			http.SetCookie(w, &http.Cookie{
				Path:  "/",
				Name:  "username",
				Value: user.Username,
			})
		}
		http.Redirect(w, r, "/dashboard/", http.StatusSeeOther)
	}

	uadmin.RenderHTML(w, r, "templates/login.html", map[string]interface{}{
		"Title": "Login - MDRRMO System",
	})
}
