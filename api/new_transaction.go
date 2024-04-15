package api

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/mbdeguzman/safeguard/models"
	"github.com/uadmin/uadmin"
)

func NewTransaction(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) == "POST" {
		transaction := models.Transaction{}
		line := models.InventoryLine{}

		transaction_type, err := strconv.ParseInt(r.FormValue("type"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			return
		}
		item_id, err := strconv.ParseUint(r.FormValue("item"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			return
		}
		qty, err := strconv.ParseInt(r.FormValue("quantity"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			return
		}
		enforcer, err := strconv.ParseUint(r.FormValue("enforcer"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": err.Error(),
			})
			return
		}

		uadmin.Get(&line, "item_id = ?", r.FormValue("item"))
		if int(transaction_type) == 1 {
			transaction.Code = generateCode()
			transaction.ItemID = uint(item_id)
			transaction.Quantity = int(qty)
			transaction.EnforcerID = uint(enforcer)
			transaction.TransactionType = models.TransactionType(1)
			uadmin.Save(&transaction)

			line.ItemID = uint(item_id)
			line.Remaining += int(qty)
			uadmin.Save(&line)
		}

		if int(transaction_type) == 2 {
			if line.Remaining > int(qty) {
				transaction.Code = generateCode()
				transaction.ItemID = uint(item_id)
				transaction.Quantity = int(qty)
				transaction.EnforcerID = uint(enforcer)
				transaction.TransactionType = models.TransactionType(1)
				uadmin.Save(&transaction)

				line.ItemID = uint(item_id)
				line.Remaining += int(qty)
				uadmin.Save(&line)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				uadmin.ReturnJSON(w, r, map[string]interface{}{
					"status":  "error",
					"err_msg": "Inventory out quantity should be less than the item remaining",
				})
				return
			}
		}
		http.Redirect(w, r, "/console/transactions/", http.StatusSeeOther)
	}
}

func generateCode() string {
	now := time.Now()
	formattedTime := now.Format("01022006:150405")
	return formattedTime
}
