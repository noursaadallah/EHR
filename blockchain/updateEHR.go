package blockchain

// UpdateEHR : Add appointment info to ehr
func (setup *FabricSetup) UpdateEHR(ehrID string, drID string, comment string) (string, error) {

	// Prepare arguments
	funcName := "updateEHR"
	args := []string{ehrID, drID, comment}

	return setup.Invoke(funcName, args)
}
