package blockchain

// QueryHello query the chaincode to get the state of hello
func (setup *FabricSetup) QueryHello() (string, error) {

	// Prepare arguments
	funcName := "invoke"
	var args []string
	args = append(args, "query")
	args = append(args, "hello")

	payload, err := setup.Query(funcName, args)
	return string(payload), err
}
