package controllers

import (
	"fmt"
	"net/http"

	"github.com/noursaadallah/EHR/model"
)

// GetEHRhandler : controller to get ehr
func (app *Application) GetEHRhandler(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		EHR      model.EHR
		Success  bool
		Response bool
	}{
		EHR:      *new(model.EHR),
		Success:  false,
		Response: false,
	}
	if r.FormValue("submitted") == "true" {
		ehrID := r.FormValue("ehrID")
		ehr, err := app.Fabric.GetEHR(ehrID)
		if err != nil {
			fmt.Println("Error getting state : " + err.Error())
			http.Error(w, "Unable to invoke getEHR in the blockchain", 500)
		}
		data.EHR = *ehr
		data.Success = true
		data.Response = true
	}
	renderTemplate(w, r, "getEHR.html", data)
}
