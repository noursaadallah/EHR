package blockchain

import (
	"fmt"
	"time"

	"github.com/hyperledger/fabric-sdk-go/api/apitxn/chclient"
)

// Invoke : generic Invoke to change state
func (setup *FabricSetup) Invoke(funcName string, args []string) (string, error) {

	var byteArgs [][]byte
	for i := 0; i < len(args); i++ {
		byteArgs = append(byteArgs, []byte(args[i]))
	}

	eventID := "eventInvoke"

	// Add data that will be visible in the proposal, like a description of the invoke request
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in invoke")

	// Register a notification handler on the client
	notifier := make(chan *chclient.CCEvent)
	rce, err := setup.client.RegisterChaincodeEvent(notifier, setup.ChainCodeID, eventID)
	if err != nil {
		return "", fmt.Errorf("failed to register chaincode event: %v", err)
	}

	// Create a request (proposal) and send it
	response, err := setup.client.Execute(chclient.Request{ChaincodeID: setup.ChainCodeID,
		Fcn:          funcName,
		Args:         byteArgs,
		TransientMap: transientDataMap})
	if err != nil {
		return "", fmt.Errorf("failed to move funds: %v", err)
	}

	// Wait for the result of the submission
	select {
	case ccEvent := <-notifier:
		fmt.Printf("Received CC event: %s\n", ccEvent)
	case <-time.After(time.Second * 30):
		return "", fmt.Errorf("did NOT receive CC event for eventId(%s)", eventID)
	}

	// Unregister the notification handler previously created on the client
	err = setup.client.UnregisterChaincodeEvent(rce)

	return response.TransactionID.ID, nil
}

// Query query the chaincode to get the state of key
func (setup *FabricSetup) Query(funcName string, args []string) ([]byte, error) {

	// Prepare arguments
	var byteArgs [][]byte
	for i := 0; i < len(args); i++ {
		byteArgs = append(byteArgs, []byte(args[i]))
	}

	response, err := setup.client.Query(chclient.Request{
		ChaincodeID: setup.ChainCodeID,
		Fcn:         funcName,
		Args:        byteArgs})
	if err != nil {
		return nil, fmt.Errorf("failed to query: %v", err)
	}

	return response.Payload, nil
}
