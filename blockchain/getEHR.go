package blockchain

import (
	"encoding/json"
	"fmt"

	"github.com/noursaadallah/EHR/model"
)

// GetEHR query the chaincode to get the state of ehrID
func (setup *FabricSetup) GetEHR(ehrID string) (*model.EHR, error) {

	// Prepare arguments
	funcName := "getEHR"
	args := []string{ehrID}

	// invoke the chaincode and return
	payload, err := setup.Query(funcName, args)
	if err != nil {
		fmt.Println("error reading state of EHR : " + err.Error())
		return nil, err
	}
	var ehr model.EHR
	err = json.Unmarshal(payload, &ehr)
	if err != nil {
		fmt.Println("Error unmarshalling JSON : " + err.Error())
		return nil, err
	}
	return &ehr, nil
}
