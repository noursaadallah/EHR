package blockchain

// CreateEHR : create EHR
func (setup *FabricSetup) CreateEHR(firstName string, lastName string,
	socialSecNbr string, birthday string) (string, error) {

	// Prepare arguments
	funcName := "createEHR"
	args := []string{firstName, lastName, socialSecNbr, birthday}

	return setup.Invoke(funcName, args)
}
