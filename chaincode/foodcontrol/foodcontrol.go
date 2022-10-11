package main

import (

	"encoding/json"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

 type SmartContract struct (
	contractapi.Contract
)

func main(){
	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil{
		fmt.Printf("Error create foodcontrol chaincode %s", err.Error())
		return 
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error create foodcontrol chaincode %s", err.Error())
	}

}