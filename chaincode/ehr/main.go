package main

import (
	"ehr_asset"
	"encoding/json"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	peer "github.com/hyperledger/fabric-protos-go/peer"
)

type EHealth struct {
	contractapi.Contract
}

func (t *EHealth) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *EHealth) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// Function name and arguments
	fn, args := stub.GetFunctionAndParameters()

	switch fn {
	case "CreateRecord":
		if len(args) != 1 {
			return peer.Response{
				Status:  500,
				Message: "Wrong number of arguments",
				Payload: nil,
			}
		}
		return peer.Response{
			Status:  200,
			Payload: nil,
		}

	case "AddPrescription":
		if len(args) != 1 {
			return peer.Response{
				Status:  500,
				Message: "Wrong number of arguments",
				Payload: nil,
			}
		}
		return peer.Response{
			Status:  200,
			Payload: nil,
		}

	case "AddAppointment":
		if len(args) != 1 {
			return peer.Response{
				Status:  500,
				Message: "Wrong number of arguments",
				Payload: nil,
			}
		}
		return peer.Response{
			Status:  200,
			Payload: nil,
		}

	case "AddProcedure":
		if len(args) != 1 {
			return peer.Response{
				Status:  500,
				Message: "Wrong number of arguments",
				Payload: nil,
			}
		}
		return peer.Response{
			Status:  200,
			Payload: nil,
		}

	case "ReadPrescription":
		if len(args) != 1 {
			return peer.Response{
				Status:  500,
				Message: "Wrong number of arguments",
				Payload: nil,
			}
		}
		return peer.Response{
			Status:  200,
			Payload: nil,
		}

	case "ReadAppointment":
		if len(args) != 1 {
			return peer.Response{
				Status:  500,
				Message: "Wrong number of arguments",
				Payload: nil,
			}
		}
		return peer.Response{
			Status:  200,
			Payload: nil,
		}

	case "ReadProcedure":
		if len(args) != 1 {
			return peer.Response{
				Status:  500,
				Message: "Wrong number of arguments",
				Payload: nil,
			}
		}
		return peer.Response{
			Status:  200,
			Payload: nil,
		}

	// Read all patint's records
	case "ReadAllRecords":
		if len(args) != 1 {
			return peer.Response{
				Status:  500,
				Message: "Wrong number of arguments",
				Payload: nil,
			}
		}
		return peer.Response{
			Status:  200,
			Payload: nil,
		}

	default:
		return peer.Response{
			Status:  500,
			Message: "Invalid function name",
			Payload: nil,
		}
	}
}

func createRecord(stub shim.ChaincodeStubInterface, ownerID string) error {
	exists := recordExists(stub, ownerID)

	if exists {
		return nil
	}

	asset := ehr_asset.EHR_Asset{
		PatientID:     ownerID,
		Prescriptions: []ehr_asset.Prescription{},
		Appointments:  []ehr_asset.Appointment{},
		Procedures:    []ehr_asset.Procedure{},
	}

	assetBytes, err := json.Marshal(asset)

	if err != nil {
		return nil
	}

	err = stub.PutState(ownerID, assetBytes)

	if err != nil {
		return nil
	}

	return nil
}

func addPrescription(stub shim.ChaincodeStubInterface, ownerID string, prescriptionJSON string) error {
	// Check if record exists
	exists := recordExists(stub, ownerID)

	if !exists {
		return nil
	}

	// Read record from world state
	recordBytes, err := stub.GetState(ownerID)

	if err != nil {
		return nil
	}

	var ehrRecord ehr_asset.EHR_Asset

	err = json.Unmarshal(recordBytes, &ehrRecord)

	if err != nil {
		return nil
	}

	// Unmarshal prescription data
	var prescription ehr_asset.Prescription
	err = json.Unmarshal(([]byte)(prescriptionJSON), prescription)

	if err != nil {
		return nil
	}

	// ehr_asset.validatePrescription()

	return nil
}

func addAppointment(stub shim.ChaincodeStubInterface, ownerID string, appointmentJSON string) error {
	return nil
}

func addProcedure(stub shim.ChaincodeStubInterface, ownerID string, procedureJSON string) error {
	return nil
}

// Read owner's complete record
func readRecord(stub shim.ChaincodeStubInterface, ownerID string) (ehr_asset.EHR_Asset, error) {
	var record ehr_asset.EHR_Asset
	return record, nil
}

// Read all prescriptions given by professional
func readPrescriptions(stub shim.ChaincodeStubInterface, ownerID string, professionalID string) ([]ehr_asset.Prescription, error) {
	return nil, nil
}

// Read all appointments by professional
func readAppointments(stub shim.ChaincodeStubInterface, ownerID string, professionalID string) ([]ehr_asset.Appointment, error) {
	return nil, nil
}

// Real all procedures by professional
func readProcedures(stub shim.ChaincodeStubInterface, ownerID string, professionalID string) ([]ehr_asset.Procedure, error) {
	return nil, nil
}

func recordExists(stub shim.ChaincodeStubInterface, ownerID string) bool {

	response, err := stub.GetState(ownerID)

	if err != nil {
		return false
	}

	return response != nil
}
