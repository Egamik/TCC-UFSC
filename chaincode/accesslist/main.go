package main

import (
	"log"

	al "accesslist/al-contract"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

func main() {
	alContract, err := contractapi.NewChaincode(&al.AccessListContract{})
	if err != nil {
		log.Panicf("Error creating ehr chaincode: %v", err)
	}

	if err := alContract.Start(); err != nil {
		log.Panicf("Error starting ehr chaincode: %v", err)
	}
}
