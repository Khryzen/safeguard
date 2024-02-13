package views

import (
	"fmt"
	"net/http"

	"github.com/uadmin/uadmin"
)

func ConsoleLoginHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Loginhandler")
	// Check if there is an existing session
	session := uadmin.IsAuthenticated(r)
	if session != nil {
		http.Redirect(w, r, "/console/dashboard/", http.StatusSeeOther)
		return
	}
	// Check if request method is post
	// If method is post, process login and create session
	if r.Method == "POST" {
		// Get the values from the forms
		username := r.FormValue("username")
		password := r.FormValue("password")
		uadmin.Trail(uadmin.DEBUG, "Username: %s | Password: %s", username, password)
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
		http.Redirect(w, r, "/console/dashboard/", http.StatusSeeOther)
	}

	uadmin.RenderHTML(w, r, "templates/console/login.html", map[string]interface{}{
		"Title": "Login - MDRRMO System Admin",
	})
}
