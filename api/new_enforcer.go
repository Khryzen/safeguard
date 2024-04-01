package api

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/mbdeguzman/safeguard/models"
	"github.com/uadmin/uadmin"
)

func NewEnforcer(w http.ResponseWriter, r *http.Request) {
	method := strings.ToUpper(r.Method)
	if method == "POST" {
		f, fh, err := r.FormFile("photo")
		if err != nil {
			uadmin.Trail(uadmin.ERROR, "Form file error. %s", err.Error())
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			return
		}

		if fh.Size <= 0 {
			w.WriteHeader(400)
			w.Write([]byte("Got File length <= 0"))
			return
		}

		outFile, err := getFile(fh.Filename)
		if err != nil {
			uadmin.Trail(uadmin.ERROR, "Get file error. %s", err.Error())
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			return
		}

		written, err := io.Copy(outFile, f)
		if err != nil {
			uadmin.Trail(uadmin.ERROR, "Copy error. %s", err.Error())
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			return
		}
		uadmin.Trail(uadmin.INFO, "Written. %#v", written)

		user := uadmin.User{}
		user.FirstName = r.FormValue("firstName")
		user.LastName = r.FormValue("lastName")
		user.Email = r.FormValue("emailAddress")
		user.Username = r.FormValue("emailAddress")
		user.Active = true
		user.RemoteAccess = true
		user.Password = "default_user"
		user.Photo = "/media/img/" + fh.Filename
		user.Save()

		enforcer := models.Enforcers{}
		enforcer.Position = r.FormValue("position")
		enforcer_type_id, _ := strconv.ParseUint(r.FormValue("enforcerType"), 10, 64)
		enforcer.EnforcerTypeID = uint(enforcer_type_id)
		uadmin.Get(&user, "username = ?", r.FormValue("emailAddress"))
		enforcer.User = user
		enforcer.Active = true
		uadmin.Save(&enforcer)

		http.Redirect(w, r, r.FormValue("redirect_url"), http.StatusSeeOther)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getFile(filename string) (*os.File, error) {
	// Open or create the file with the given filename
	// You can adjust file opening options as needed
	file, err := os.Create("media/img/" + filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}
