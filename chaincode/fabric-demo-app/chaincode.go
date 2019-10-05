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

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
	"github.com/fabric-demo-app/iota"
)

// Define the Smart Contract structure
type SmartContract struct {
}

/* Define Container structure, with 4 properties.  
Structure tags are used by encoding/json library
*/
type Container struct {
	Vessel string `json:"vessel"`
	Timestamp string `json:"timestamp"`
	Location  string `json:"location"`
	Holder  string `json:"holder"`
}

type Response struct {
	Name       		string `json:"name"`
	Model       	string `json:"model"`
	Manufacturer    string `json:"manufacturer"`
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
	return shim.Success(containerAsBytes)
}

/*
 * The initLedger method *
Will add test data (10 containers) to our network
 */
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	containers := []Container{
		Container{Vessel: "923F", Location: "67.0006, -70.5476", Timestamp: "1504054225", Holder: "Alex"},
		Container{Vessel: "M83T", Location: "91.2395, -49.4594", Timestamp: "1504057825", Holder: "Dave"},
		Container{Vessel: "T012", Location: "58.0148, 59.01391", Timestamp: "1493517025", Holder: "Igor"},
		Container{Vessel: "P490", Location: "-45.0945, 0.7949", Timestamp: "1496105425", Holder: "Amalea"},
		Container{Vessel: "S439", Location: "-107.6043, 19.5003", Timestamp: "1493512301", Holder: "Rafa"},
		Container{Vessel: "J205", Location: "-155.2304, -15.8723", Timestamp: "1494117101", Holder: "Shen"},
		Container{Vessel: "S22L", Location: "103.8842, 22.1277", Timestamp: "1496104301", Holder: "Leila"},
		Container{Vessel: "EI89", Location: "-132.3207, -34.0983", Timestamp: "1485066691", Holder: "Yuan"},
		Container{Vessel: "129R", Location: "153.0054, 12.6429", Timestamp: "1485153091", Holder: "Carlo"},
		Container{Vessel: "49W4", Location: "51.9435, 8.2735", Timestamp: "1487745091", Holder: "Bobby"},
	}

	i := 0
	for i < len(containers) {
		fmt.Println("i is ", i)
		containerAsBytes, _ := json.Marshal(containers[i])
		APIstub.PutState(strconv.Itoa(i+1), containerAsBytes)
		fmt.Println("Added", containers[i])
		i = i + 1
	}

	return shim.Success(nil)
}

/*
 * The recordContainer method *
Container owners like Sarah would use to record each of her containers. 
This method takes in five arguments (attributes to be saved in the ledger). 
 */
func (s *SmartContract) recordContainer(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var container = Container{ Vessel: args[1], Location: args[2], Timestamp: args[3], Holder: args[4] }

	containerAsBytes, _ := json.Marshal(container)
	err := APIstub.PutState(args[0], containerAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record container: %s", args[0]))
	}

	return shim.Success(nil)
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
	container.Holder = args[1]

	containerAsBytes, _ = json.Marshal(container)
	err := APIstub.PutState(args[0], containerAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change container holder: %s", args[0]))
	}

	// initiate IOTA transaction
	// iota.TransferTokens()
	var randomNumber = iota.Random()

	fmt.Println("randomNumber", randomNumber)
	
	rsp := &Response{}
	if err := iota.MakeRequest1("https://swapi.co/api/vehicles/42", rsp); err != nil {
		fmt.Println(666, err)
	}
	// b := []byte("My string " + strconv.Itoa(randomNumber))

	return shim.Success([]byte("changeContainerHolder success * "+ " | " + rsp.Name + " | " + strconv.Itoa(randomNumber)))
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
