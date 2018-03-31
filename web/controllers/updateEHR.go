package controllers

import (
	"net/http"
)

// UpdateEHRhandler : controller to Add appointment to EHR
func (app *Application) UpdateEHRhandler(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		TransactionID string
		Success       bool
		Response      bool
	}{
		TransactionID: "",
		Success:       false,
		Response:      false,
	}
	if r.FormValue("submitted") == "true" {
		ehrID := r.FormValue("ehrID")
		drID := r.FormValue("drID")
		comment := r.FormValue("comment")

		txid, err := app.Fabric.UpdateEHR(ehrID, drID, comment)
		if err != nil {
			http.Error(w, "Unable to invoke updateEHR in the blockchain : "+err.Error(), 500)
		}
		data.TransactionID = txid
		data.Success = true
		data.Response = true
	}
	renderTemplate(w, r, "updateEHR.html", data)
}
