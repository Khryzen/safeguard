package api

import (
	"io"
	"net/http"
	"strings"

	"github.com/mbdeguzman/safeguard/models"
	"github.com/uadmin/uadmin"
)

func EditItem(w http.ResponseWriter, r *http.Request) {
	method := strings.ToUpper(r.Method)
	if method == "POST" {
		f, fh, err := r.FormFile("edit_photo")
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

		item := models.Item{}
		uadmin.Get(&item, "id = ?", r.FormValue("id"))
		item.Name = r.FormValue("edit_name")
		item.Image = "/media/img/" + fh.Filename
		item.Save()

		http.Redirect(w, r, r.FormValue("redirect_url"), http.StatusSeeOther)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
