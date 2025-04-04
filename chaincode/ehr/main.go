package main

import (
	"log"

	ehr "ehr/ehr-contract"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

func main() {
	ehrContract, err := contractapi.NewChaincode(&ehr.EHRContract{})
	if err != nil {
		log.Panicf("Error creating ehr chaincode: %v", err)
	}

	if err := ehrContract.Start(); err != nil {
		log.Panicf("Error starting ehr chaincode: %v", err)
	}
}
