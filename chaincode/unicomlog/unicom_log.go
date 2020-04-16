package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type LogChaincode struct {
}
func (t *LogChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Log Chaincode Init")
	
	return shim.Success(nil)
}

func (t *LogChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Log Chaincode Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "save" {
		// Make payment of X units from A to B
		return t.save(stub, args)
	} else if function == "query" {
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	}
	
	return shim.Error("Invalid invoke function name. Expecting \"save\" \"query\"")
}

// Transaction makes save log
func (t *LogChaincode) save(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	fmt.Println("Execute operation of saving log")
	
	log_key := args[0]
	log_value := args[1]
	
	err := stub.PutState(log_key, []byte(log_value))
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("Execute operation of saving log suyccessfully")
	return shim.Success(nil)
}

// query callback representing the query of a chaincode
func (t *LogChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var key string // Entities
	var err error
	
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting time of the log to query")
	}
	
	key = args[0]
	
	// Get the state from the ledger
	Avalbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get log for " + key + "\"}"
		return shim.Error(jsonResp)
	}
	
	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Empty log for " + key + "\"}"
		return shim.Error(jsonResp)
	}
	
	jsonResp := "{\"Time\":\"" + key + "\",\"Log\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}

func main() {
	err := shim.Start(new(LogChaincode))
	if err != nil {
		fmt.Printf("Error starting log chaincode: %s", err)
	}
}

