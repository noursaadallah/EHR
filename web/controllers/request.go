package controllers

import (
	"net/http"
)

// RequestHandler : controller to change hello state
func (app *Application) RequestHandler(w http.ResponseWriter, r *http.Request) {
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
		helloValue := r.FormValue("hello")
		txid, err := app.Fabric.InvokeHello(helloValue)
		if err != nil {
			http.Error(w, "Unable to invoke hello in the blockchain", 500)
		}
		data.TransactionID = txid
		data.Success = true
		data.Response = true
	}
	renderTemplate(w, r, "request.html", data)
}
