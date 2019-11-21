// SPDX-License-Identifier: Apache-2.0

/*
  Sample Chaincode based on Demonstrated Scenario

 This code is based on code written by the Hyperledger Fabric community.
  Original code can be found here: https://github.com/hyperledger/fabric-samples/blob/release/chaincode/fabcar/fabcar.go
 */

package main

/* Imports  
* 4 utility libraries for handling bytes, reading and writing JSON, 
formatting, and string manipulation  
* 2 specific Hyperledger Fabric specific libraries for Smart Contracts  
*/ 
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
	"github.com/iota"
)

// Define the Smart Contract structure
type SmartContract struct {
}

/* Define Container structure, with 4 properties.  
Structure tags are used by encoding/json library
*/
type Container struct {
	Description string `json:"description"`
	Timestamp string `json:"timestamp"`
	Location  string `json:"location"`
	Holder  string `json:"holder"`
}

type IotaWallet struct {
	Seed        string `json:"seed"`
	Address     string `json:"address"`
	KeyIndex    uint64 `json:"keyIndex"`
}

type Participant struct {
	Role string `json:"role"`
	Description string `json:"description"`
	IotaWallet
}

type IotaPayload struct {
	Seed        string `json:"seed"`
	MamState    string `json:"mamState"`
	Root        string `json:"root"`
	Mode       	string `json:"mode"`
	SideKey     string `json:"sideKey"`
}

/*
 * The Init method *
 called when the Smart Contract "container-chaincode" is instantiated by the network
 * Best practice is to have any Ledger initialization in separate function 
 -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method *
 called when an application requests to run the Smart Contract "container-chaincode"
 The app also specifies the specific smart contract function to call with args
 */

 // https://hyperledger-fabric.readthedocs.io/en/release-1.4/chaincode4ade.html
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "queryContainer" {
		return s.queryContainer(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "recordContainer" {
		return s.recordContainer(APIstub, args)
	} else if function == "queryAllContainers" {
		return s.queryAllContainers(APIstub)
	} else if function == "changeContainerHolder" {
		return s.changeContainerHolder(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

/*
 * The initLedger method *
Will add test data (10 containers) to our network
 */
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	timestamp := strconv.FormatInt(time.Now().UnixNano() / 1000000, 10)
	containers := []Container{
		Container{Description: "Apples", Location: "67.0006, -70.5476", Timestamp: timestamp, Holder: "Producer"},
		Container{Description: "Oranges", Location: "91.2395, -49.4594", Timestamp: timestamp, Holder: "Freight Forwarder"},
		Container{Description: "Avocados", Location: "58.0148, 59.01391", Timestamp: timestamp, Holder: "Customs"},
		Container{Description: "Pineapples", Location: "-45.0945, 0.7949", Timestamp: timestamp, Holder: "Producer"},
		Container{Description: "Olives", Location: "-107.6043, 19.5003", Timestamp: timestamp, Holder: "Shipper"},
		Container{Description: "Mangos", Location: "-155.2304, -15.8723", Timestamp: timestamp, Holder: "Distributor"},
		Container{Description: "Grapefruits", Location: "103.8842, 22.1277", Timestamp: timestamp, Holder: "Customs"},
		Container{Description: "Watermelons", Location: "-132.3207, -34.0983", Timestamp: timestamp, Holder: "Freight Forwarder"},
		Container{Description: "Bananas", Location: "153.0054, 12.6429", Timestamp: timestamp, Holder: "Shipper"},
		Container{Description: "Clementines", Location: "51.9435, 8.2735", Timestamp: timestamp, Holder: "Retailer"},
	}

	i := 0
	for i < len(containers) {
		containerAsBytes, _ := json.Marshal(containers[i])
		APIstub.PutState(strconv.Itoa(i+1), containerAsBytes)

		// Define own values for IOTA MAM message mode and MAM message encryption key
		// If not set, default values from iota/config.go file will be used
		mode := iota.MamMode
		sideKey := iota.PadSideKey(iota.MamSideKey) // iota.PadSideKey(iota.GenerateRandomSeedString(50))
		
		mamState, root, seed := iota.PublishAndReturnState(string(containerAsBytes), false, "", "", mode, sideKey)
		iotaPayload := IotaPayload{Seed: seed, MamState: mamState, Root: root, Mode: mode, SideKey: sideKey}
		iotaPayloadAsBytes, _ := json.Marshal(iotaPayload)
		APIstub.PutState("IOTA_" + strconv.Itoa(i+1), iotaPayloadAsBytes)

		fmt.Println("New Asset", strconv.Itoa(i+1), containers[i], root, mode, sideKey)
		
		i = i + 1
	}

	participants := []Participant{
		Participant{Role: "Producer", Description: "Farmer / Goods producer"},
		Participant{Role: "Freight Forwarder", Description: "Logistics"},
		Participant{Role: "Customs", Description: ""},
		Participant{Role: "Shipper", Description: ""},
		Participant{Role: "Distributor", Description: "Fruits distributor"},
		Participant{Role: "Retailer", Description: "Large grocery store"},
	}

	for i := range participants {
		walletAddress, walletSeed := iota.CreateWallet()
		participants[i].Seed = walletSeed
		participants[i].Address = walletAddress
		participants[i].KeyIndex = 0
		participantAsBytes, _ := json.Marshal(participants[i])
		APIstub.PutState(participants[i].Role, participantAsBytes)
	}

	iotaWallet := IotaWallet{Seed: iota.DefaultWalletSeed, KeyIndex: iota.DefaultWalletKeyIndex, Address: ""}
	iotaWalletAsBytes, _ := json.Marshal(iotaWallet)
	APIstub.PutState("IOTA_WALLET", iotaWalletAsBytes)

	return shim.Success(nil)
}

/*
 * The recordContainer method *
Container owners like Sarah would use to record each of her containers. 
This method takes in five arguments (attributes to be saved in the ledger). 
 */
func (s *SmartContract) recordContainer(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano() / 1000000, 10)
	container := Container{ Description: args[1], Location: args[2], Timestamp: timestamp, Holder: args[3] }

	containerAsBytes, _ := json.Marshal(container)
	err := APIstub.PutState(args[0], containerAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record container: %s", args[0]))
	}

	// Define own values for IOTA MAM message mode and MAM message encryption key
	// If not set, default values from iota/config.go file will be used
	mode := iota.MamMode
	sideKey := iota.PadSideKey(iota.MamSideKey) // iota.PadSideKey(iota.GenerateRandomSeedString(50))
	
	mamState, root, seed := iota.PublishAndReturnState(string(containerAsBytes), false, "", "", mode, sideKey)	
	iotaPayload := IotaPayload{Seed: seed, MamState: mamState, Root: root, Mode: mode, SideKey: sideKey}
	iotaPayloadAsBytes, _ := json.Marshal(iotaPayload)
	APIstub.PutState("IOTA_" + args[0], iotaPayloadAsBytes)

	fmt.Println("New Asset", args[0], container, root, mode, sideKey)

	return shim.Success(nil)
}

/*
 * The queryContainer method *
Used to view the records of one particular container
It takes one argument -- the key for the container in question
 */
 func (s *SmartContract) queryContainer(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	containerAsBytes, _ := APIstub.GetState(args[0])
	if containerAsBytes == nil {
		return shim.Error("Could not locate container")
	}
	container := Container{}
	json.Unmarshal(containerAsBytes, &container)

	iotaPayloadAsBytes, _ := APIstub.GetState("IOTA_" + args[0])
	if iotaPayloadAsBytes == nil {
		return shim.Error("Could not locate IOTA state object")
	}
	iotaPayload := IotaPayload{}
	json.Unmarshal(iotaPayloadAsBytes, &iotaPayload)

	mamstate := map[string]interface{}{}
	mamstate["root"] = iotaPayload.Root
	mamstate["sideKey"] = iotaPayload.SideKey

	// IOTA MAM stream values
	messages := iota.Fetch(iotaPayload.Root, iotaPayload.Mode, iotaPayload.SideKey)

	participantAsBytes, _ := APIstub.GetState(container.Holder)
	if participantAsBytes == nil {
		return shim.Error("Could not locate participant")
	}
	participant := Participant{}
	json.Unmarshal(participantAsBytes, &participant)

	out := map[string]interface{}{}
	out["container"] = container
	out["mamstate"] = mamstate
	out["messages"] = strings.Join(messages, ", ")
	out["wallet"] = participant.Address
	
	result, _ := json.Marshal(out)

	return shim.Success(result)
}

/*
 * The queryAllContainers method *
allows for assessing all the records added to the ledger(all containers)
This method does not take any arguments. Returns JSON string containing results. 
 */
func (s *SmartContract) queryAllContainers(APIstub shim.ChaincodeStubInterface) sc.Response {
	startKey := "0"
	endKey := "999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		// Add comma before array members,suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		
		// iotaPayloadAsBytes, _ := APIstub.GetState("IOTA_" + queryResponse.Key)
		// if iotaPayloadAsBytes == nil {
		// 	return shim.Error("Could not locate IOTA state object")
		// }
		// iotaPayload := IotaPayload{}
		// json.Unmarshal(iotaPayloadAsBytes, &iotaPayload)

		// buffer.WriteString(", \"Root\":")
		// buffer.WriteString("\"" + string(iotaPayload.Root) + "\"")
		// buffer.WriteString(", \"SideKey\":")
		// buffer.WriteString("\"" + string(iotaPayload.SideKey) + "\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllContainers:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

/*
 * The changeContainerHolder method *
The data in the world state can be updated with who has possession. 
This function takes in 2 arguments, container ID and new holder name. 
 */
func (s *SmartContract) changeContainerHolder(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	containerAsBytes, _ := APIstub.GetState(args[0])
	if containerAsBytes == nil {
		return shim.Error("Could not locate container")
	}
	container := Container{}

	json.Unmarshal(containerAsBytes, &container)
	// Normally check that the specified argument is a valid holder of a container
	// we are skipping this check for this example
	previousContainerHolder := container.Holder
	container.Holder = args[1]

	timestamp := strconv.FormatInt(time.Now().UnixNano() / 1000000, 10)
	container.Timestamp = timestamp

	containerAsBytes, _ = json.Marshal(container)
	err := APIstub.PutState(args[0], containerAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change container holder: %s", args[0]))
	}

	iotaPayloadAsBytes, _ := APIstub.GetState("IOTA_" + args[0])
	if iotaPayloadAsBytes == nil {
		return shim.Error("Could not locate IOTA state object")
	}
	iotaPayload := IotaPayload{}
	json.Unmarshal(iotaPayloadAsBytes, &iotaPayload)

	mamState, _, _ := iota.PublishAndReturnState(string(containerAsBytes), true, iotaPayload.Seed, iotaPayload.MamState, iotaPayload.Mode, iotaPayload.SideKey)
	iotaPayloadNew := IotaPayload{Seed: iotaPayload.Seed, MamState: mamState, Root: iotaPayload.Root, Mode: iotaPayload.Mode, SideKey: iotaPayload.SideKey}
	iotaPayloadNewAsBytes, _ := json.Marshal(iotaPayloadNew)
	APIstub.PutState("IOTA_" + args[0], iotaPayloadNewAsBytes)

	// make payment to the participant
	participantAsBytes, _ := APIstub.GetState(previousContainerHolder)
	if participantAsBytes == nil {
		return shim.Error("Could not locate participant")
	}
	participant := Participant{}
	json.Unmarshal(participantAsBytes, &participant)

	iotaWalletAsBytes, _ := APIstub.GetState("IOTA_WALLET")
	if iotaWalletAsBytes == nil {
		return shim.Error("Could not locate wallet data")
	}
	iotaWallet := IotaWallet{}
	json.Unmarshal(iotaWalletAsBytes, &iotaWallet)

	newKeyIndex := iota.TransferTokens(iotaWallet.Seed, iotaWallet.KeyIndex, participant.Address)
	iotaWallet.KeyIndex = newKeyIndex
	iotaWalletAsBytes, _ = json.Marshal(iotaWallet)
	err = APIstub.PutState("IOTA_WALLET", iotaWalletAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to update wallet with index: %s", strconv.FormatUint(newKeyIndex, 10)))
	}

	return shim.Success([]byte("changeContainerHolder success"))
}

/*
 * main function *
calls the Start function 
The main function starts the chaincode in the container during instantiation.
 */
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
