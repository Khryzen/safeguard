package api

import (
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/mbdeguzman/safeguard/models"
	"github.com/uadmin/uadmin"
)

func EditEnforcer(w http.ResponseWriter, r *http.Request) {
	method := strings.ToUpper(r.Method)
	if method == "POST" {
		// f, fh, err := r.FormFile("ephoto")
		// if err != nil {
		// 	uadmin.Trail(uadmin.ERROR, "Form file error. %s", err.Error())
		// 	uadmin.ReturnJSON(w, r, map[string]interface{}{
		// 		"status":  "error",
		// 		"err_msg": err.Error(),
		// 	})
		// }
		f, fh, err := r.FormFile("ephoto")
		if err != nil {
			// Check if the error is ErrMissingFile
			if err != http.ErrMissingFile {
				// Some other error occurred
				uadmin.Trail(uadmin.ERROR, "Form file error. %s", err.Error())
				uadmin.ReturnJSON(w, r, map[string]interface{}{
					"status":  "error",
					"err_msg": err.Error(),
				})
				return // Stop execution or handle the error accordingly
			}
			// No file uploaded, just proceed
			uadmin.Trail(uadmin.INFO, "No file uploaded")
			// Proceed with your application logic
			uadmin.Trail(uadmin.DEBUG, "%#v", fh)
			if fh != nil {
				if fh.Size <= 0 {
					w.WriteHeader(400)
					w.Write([]byte("Got File length <= 0"))
				}
				outFile, err := getFile(fh.Filename)
				if err != nil {
					uadmin.Trail(uadmin.ERROR, "Get file error. %s", err.Error())
					uadmin.ReturnJSON(w, r, map[string]interface{}{
						"status":  "error",
						"err_msg": err.Error(),
					})
				}
				written, err := io.Copy(outFile, f)
				if err != nil {
					uadmin.Trail(uadmin.ERROR, "Copy error. %s", err.Error())
					uadmin.ReturnJSON(w, r, map[string]interface{}{
						"status":  "error",
						"err_msg": err.Error(),
					})
				}
				uadmin.Trail(uadmin.INFO, "Written. %#v", written)
			}

			enforcer := models.Enforcers{}
			uadmin.Get(&enforcer, "id = ?", r.FormValue("id"))
			user := uadmin.User{}
			uadmin.Get(&user, "id = ?", enforcer.UserID)
			user.FirstName = r.FormValue("efirstName")
			user.LastName = r.FormValue("elastName")
			user.Email = r.FormValue("eemailAddress")
			enforcer.Position = r.FormValue("eposition")
			enforcerTypeID, _ := strconv.ParseUint(r.FormValue("eenforcerType"), 10, 64)
			enforcer.EnforcerTypeID = uint(enforcerTypeID)
			user.Save()
			uadmin.Save(&enforcer)
			http.Redirect(w, r, r.FormValue("redirect_url"), http.StatusSeeOther)
		} else {
			if fh != nil {
				if fh.Size <= 0 {
					w.WriteHeader(400)
					w.Write([]byte("Got File length <= 0"))
				}
				outFile, err := getFile(fh.Filename)
				if err != nil {
					uadmin.Trail(uadmin.ERROR, "Get file error. %s", err.Error())
					uadmin.ReturnJSON(w, r, map[string]interface{}{
						"status":  "error",
						"err_msg": err.Error(),
					})
				}
				written, err := io.Copy(outFile, f)
				if err != nil {
					uadmin.Trail(uadmin.ERROR, "Copy error. %s", err.Error())
					uadmin.ReturnJSON(w, r, map[string]interface{}{
						"status":  "error",
						"err_msg": err.Error(),
					})
				}
				uadmin.Trail(uadmin.INFO, "Written. %#v", written)
			}

			enforcer := models.Enforcers{}
			uadmin.Get(&enforcer, "id = ?", r.FormValue("id"))
			user := uadmin.User{}
			uadmin.Get(&user, "id = ?", enforcer.UserID)
			user.FirstName = r.FormValue("efirstName")
			user.LastName = r.FormValue("elastName")
			user.Email = r.FormValue("eemailAddress")
			enforcer.Position = r.FormValue("eposition")
			enforcerTypeID, _ := strconv.ParseUint(r.FormValue("eenforcerType"), 10, 64)
			enforcer.EnforcerTypeID = uint(enforcerTypeID)
			user.Photo = "/media/img/" + fh.Filename
			user.Save()
			uadmin.Save(&enforcer)
			http.Redirect(w, r, r.FormValue("redirect_url"), http.StatusSeeOther)
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
