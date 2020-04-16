package main

import (
	"fmt"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type recordInfo struct {
	identity string `json:"identity"`
	sickName string `json:"sickName"`
	drugName []string `json:"drugName"`
}
type resultData struct {
	recordInfos []recordInfo `json:"recordInfos"`
}

func(r *recordInfo) Init (stub shim.ChaincodeStubInterface) pb.Response{
	return shim.Success(nil)
}

func (r *recordInfo) Invoke(stub shim.ChaincodeStubInterface) pb.Response{
	funcName,args:=stub.GetFunctionAndParameters()
	if(funcName=="save"){
		return r.saveBasic(stub,args)
	}else if(funcName=="query"){
		return r.queryBasic(stub,args)
	}else{
		return shim.Error("no such function")
	}
	
}

func (r *recordInfo) saveBasic(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	if(len(args)!=2){
		return shim.Error("expect two args")
	}else{
		err := stub.PutState(args[0],[]byte(args[1]))
		if err!=nil {
			return shim.Error(err.Error())
		}
	}
	return shim.Success(nil)
}

func (r *recordInfo) queryBasic(stub shim.ChaincodeStubInterface, args []string)pb.Response {
	if len(args)!=1{
		return shim.Error("expect one arg")
	}else{
		value, err := stub.GetState(args[0])
		if err!=nil{
			return shim.Error("no data fund")
		}
		return shim.Success(value)
	}
}
func main(){
	err := shim.Start(new(recordInfo))
	if (err !=nil){
		fmt.Println("recordInfo chaincode start error")
	}
}
