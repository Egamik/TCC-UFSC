package ehr_contract

import (
	"ehr_asset"
	"encoding/json"
	"fmt"

	chaincodeErrors "chaincodeErrors"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

type EHRContract struct {
	contractapi.Contract
}

func (s *EHRContract) createRecord(ctx contractapi.TransactionContextInterface, ownerID string) error {
	// Only professionals should be able to create records
	funcName := "CreateRecord"
	exists, err := s.recordExists(ctx, ownerID)

	// For a professional to create a record they must be authorized previously
	issuerID, err := cid.New(ctx.GetStub())

	issuerCrtID, found, err := issuerID.GetAttributeValue("personID")

	if err != nil {
		return chaincodeErrors.NewGenericError(funcName, err)
	}

	if !found {
		return chaincodeErrors.NewGenericError(funcName, err)
	}

	alResponse := ctx.GetStub().InvokeChaincode("AccessList", ToChaincodeArgs("isIdentityApproved", ownerID, issuerCrtID), "access_chanel")

	if !isAllowed {
		return chaincodeErrors.NewForbiddenAccessError(funcName, issuerID, nil)
	}

	if err != nil {
		// return chaincodeErrors.New
		return nil
	}

	if exists {
		return chaincodeErrors.NewAssetNotFoundError(funcName, ownerID, err)
	}

	asset := ehr_asset.EHR_Asset{
		PatientID:     ownerID,
		Prescriptions: []ehr_asset.Prescription{},
		Appointments:  []ehr_asset.Appointment{},
		Procedures:    []ehr_asset.Procedure{},
	}

	assetBytes, err := json.Marshal(asset)

	if err != nil {
		return chaincodeErrors.NewMarshallingError(funcName, "EHRAsset", err)
	}

	err = ctx.GetStub().PutState(ownerID, assetBytes)

	if err != nil {
		return chaincodeErrors.NewUpdateWorldStateError(funcName, err)
	}

	return err
}

func (s *EHRContract) addPrescription(ctx contractapi.TransactionContextInterface, ownerID string, prescriptionJSON string) error {
	// Check if record exists
	funcName := "AddPrescription"
	exists, err := s.recordExists(ctx, ownerID)

	if err != nil {
		return err
	}

	if !exists {
		return nil
	}

	// Read record from world state
	recordBytes, err := ctx.GetStub().GetState(ownerID)

	if err != nil {
		return chaincodeErrors.NewReadWorldStateError(funcName, err)
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

	ehrRecord.Prescriptions = append(ehrRecord.Prescriptions, prescription)

	ehrRecordJSON, err := json.Marshal(ehrRecord)

	if err != nil {
		return chaincodeErrors.NewMarshallingError(funcName, "EHRRecord", err)
	}

	err = ctx.GetStub().PutState(ownerID, ehrRecordJSON)

	if err != nil {
		return chaincodeErrors.NewUpdateWorldStateError(funcName, err)
	}

	return nil
}

func (s *EHRContract) addAppointment(ctx contractapi.TransactionContextInterface, ownerID string, appointmentJSON string) error {
	// Check if record exists
	funcName := "AddAppointment"
	exists, err := s.recordExists(ctx, ownerID)

	if err != nil {
		return err
	}

	if !exists {
		return nil
	}

	var appointment ehr_asset.Appointment
	var record ehr_asset.EHR_Asset

	// Read record from world state
	recordBytes, err := ctx.GetStub().GetState(ownerID)

	if err != nil {
		return chaincodeErrors.NewReadWorldStateError(funcName, err)
	}

	err = json.Unmarshal(recordBytes, &record)

	if err != nil {
		return chaincodeErrors.NewMarshallingError(funcName, "EHRRecord", err)
	}

	err = json.Unmarshal(([]byte)(appointmentJSON), &appointment)

	if err != nil {
		return chaincodeErrors.NewMarshallingError(funcName, "Appointment", err)
	}

	record.Appointments = append(record.Appointments, appointment)

	resultJSON, err := json.Marshal(record)

	if err != nil {
		return chaincodeErrors.NewMarshallingError(funcName, "EHRResult", err)
	}

	err = ctx.GetStub().PutState(ownerID, resultJSON)

	if err != nil {
		return chaincodeErrors.NewUpdateWorldStateError(funcName, err)
	}

	return nil
}

func (s *EHRContract) addProcedure(ctx contractapi.TransactionContextInterface, ownerID string, procedureJSON string) error {
	funcName := "AddProcedure"
	exists, err := s.recordExists(ctx, ownerID)

	if err != nil {
		return err
	}

	if !exists {
		return nil
	}

	var procedure ehr_asset.Procedure
	var record ehr_asset.EHR_Asset

	// Read record from world state
	recordBytes, err := ctx.GetStub().GetState(ownerID)

	if err != nil {
		return chaincodeErrors.NewReadWorldStateError(funcName, err)
	}

	err = json.Unmarshal(recordBytes, &record)

	if err != nil {
		return chaincodeErrors.NewMarshallingError(funcName, "EHRRecord", err)
	}

	err = json.Unmarshal(([]byte)(procedureJSON), &procedure)

	if err != nil {
		return chaincodeErrors.NewMarshallingError(funcName, "Appointment", err)
	}

	record.Procedures = append(record.Procedures, procedure)

	resultJSON, err := json.Marshal(record)

	if err != nil {
		return chaincodeErrors.NewMarshallingError(funcName, "EHRResult", err)
	}

	err = ctx.GetStub().PutState(ownerID, resultJSON)

	if err != nil {
		return chaincodeErrors.NewUpdateWorldStateError(funcName, err)
	}

	return nil
}

// Read owner's complete record
func (s *EHRContract) readRecord(ctx contractapi.TransactionContextInterface, ownerID string) (*ehr_asset.EHR_Asset, error) {
	funcName := "ReadRecord"
	var record ehr_asset.EHR_Asset

	recordBytes, err := ctx.GetStub().GetState(ownerID)

	if err != nil {
		return nil, chaincodeErrors.NewReadWorldStateError(funcName, err)
	}

	err = json.Unmarshal(recordBytes, &record)

	if err != nil {
		return nil, chaincodeErrors.NewMarshallingError(funcName, "Record", err)
	}

	return &record, nil
}

// Read all prescriptions given by specified professional
func (s *EHRContract) readPrescriptions(ctx contractapi.TransactionContextInterface, ownerID string, professionalID string) ([]ehr_asset.Prescription, error) {
	// funcName := "ReadPrescription"

	record, err := s.readRecord(ctx, ownerID)

	if err != nil {
		return nil, err
	}

	var prescriptions []ehr_asset.Prescription
	for i := range record.Prescriptions {
		if record.Prescriptions[i].ProfessionalID == professionalID {
			prescriptions = append(prescriptions, record.Prescriptions[i])
		}
	}

	return prescriptions, nil
}

// Read all appointments by specified professional
func (s *EHRContract) readAppointments(ctx contractapi.TransactionContextInterface, ownerID string, professionalID string) ([]ehr_asset.Appointment, error) {
	// funcName := "ReadAppointments"

	record, err := s.readRecord(ctx, ownerID)

	if err != nil {
		return nil, err
	}

	var appointments []ehr_asset.Appointment
	for i := range record.Appointments {
		if record.Appointments[i].ProfessionalID == professionalID {
			appointments = append(appointments, record.Appointments[i])
		}
	}

	return appointments, nil
}

// Real all procedures by specified professional
func (s *EHRContract) readProcedures(ctx contractapi.TransactionContextInterface, ownerID string, professionalID string) ([]ehr_asset.Procedure, error) {
	funcName := "ReadProcedures"
	return nil, nil
}

func (s *EHRContract) recordExists(ctx contractapi.TransactionContextInterface, ownerID string) (bool, error) {
	funcName := "RecordExists"
	response, err := ctx.GetStub().GetState(ownerID)

	if err != nil {
		return false, chaincodeErrors.NewReadWorldStateError(funcName, err)
	}

	return response != nil, nil
}

func (s *EHRContract) isClientInAllowedList(ctx contractapi.TransactionContextInterface, ownerID string, proID string) (bool, error) {
	funcName := "IsClientInAllowedList"

	response := ctx.GetStub().InvokeChaincode("AccessList", ToChaincodeArgs("isIdentityApproved", ownerID, proID), "access_channel")

	if response.GetStatus() != 200 {
		return false, nil
	}

	payload := response.GetPayload()

	if payload == nil {
		return false, chaincodeErrors.NewGenericError(funcName, nil)
	}
	fmt.Print("IsclientInAllowedList payload: ", payload)
	return true, nil
}

func ToChaincodeArgs(args ...string) [][]byte {
	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}

	return bargs
}
