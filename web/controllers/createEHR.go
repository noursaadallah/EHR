package controllers

import (
	"net/http"
)

// CreateEHRhandler : controller to createEHR
func (app *Application) CreateEHRhandler(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		TransactionId string
		Success       bool
		Response      bool
	}{
		TransactionId: "",
		Success:       false,
		Response:      false,
	}
	if r.FormValue("submitted") == "true" {
		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")
		socialSecNbr := r.FormValue("socialSecNbr")
		birthday := r.FormValue("birthday")

		txid, err := app.Fabric.CreateEHR(firstName, lastName, socialSecNbr, birthday)
		if err != nil {
			http.Error(w, "Unable to invoke createEHR in the blockchain : "+err.Error(), 500)
		}
		data.TransactionId = txid
		data.Success = true
		data.Response = true
	}
	renderTemplate(w, r, "createEHR.html", data)
}
