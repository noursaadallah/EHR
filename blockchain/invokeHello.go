package blockchain

// InvokeHello : change hello value
func (setup *FabricSetup) InvokeHello(value string) (string, error) {

	// Prepare arguments
	funcName := "invoke"
	var args []string
	args = append(args, "invoke")
	args = append(args, "hello")
	args = append(args, value)

	return setup.Invoke(funcName, args)
}
