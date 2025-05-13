package ehr_contract

import (
	"ehr_asset"
	"encoding/json"

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

	// For a professional to create a record they must be authorized previously
	allowed, err := isClientInAllowedList(ctx, ownerID)

	if err != nil {
		return err
	}

	if !allowed {
		return chaincodeErrors.NewForbiddenAccessError(funcName, ownerID, nil)
	}

	exists, err := s.recordExists(ctx, ownerID)

	if err != nil {
		return chaincodeErrors.NewGenericError(funcName, err)
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

	return nil
}

func (s *EHRContract) addPrescription(ctx contractapi.TransactionContextInterface, ownerID string, prescriptionJSON string) error {
	// Check if record exists
	funcName := "AddPrescription"

	allowed, err := isClientInAllowedList(ctx, ownerID)

	if err != nil {
		return err
	}

	if !allowed {
		return chaincodeErrors.NewForbiddenAccessError(funcName, ownerID, nil)
	}

	exists, err := s.recordExists(ctx, ownerID)

	if err != nil {
		return err
	}

	if !exists {
		return chaincodeErrors.NewReadWorldStateError(funcName, nil)
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

	allowed, err := isClientInAllowedList(ctx, ownerID)

	if err != nil {
		return err
	}

	if !allowed {
		return chaincodeErrors.NewForbiddenAccessError(funcName, ownerID, nil)
	}

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

	allowed, err := isClientInAllowedList(ctx, ownerID)

	if err != nil {
		return err
	}

	if !allowed {
		return chaincodeErrors.NewForbiddenAccessError(funcName, ownerID, nil)
	}

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

	allowed, err := isClientInAllowedList(ctx, ownerID)

	if err != nil {
		return nil, err
	}

	if !allowed {
		return nil, chaincodeErrors.NewForbiddenAccessError(funcName, ownerID, nil)
	}

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
	funcName := "ReadPrescription"

	asset, err := s.readRecord(ctx, ownerID)

	if err != nil {
		return nil, chaincodeErrors.NewGenericError(funcName, err)
	}

	return asset.Prescriptions, nil
}

// Read all appointments by specified professional
func (s *EHRContract) readAppointments(ctx contractapi.TransactionContextInterface, ownerID string, professionalID string) ([]ehr_asset.Appointment, error) {
	funcName := "ReadAppointments"

	asset, err := s.readRecord(ctx, ownerID)

	if err != nil {
		return nil, chaincodeErrors.NewGenericError(funcName, err)
	}

	return asset.Appointments, nil
}

// Real all procedures by specified professional
func (s *EHRContract) readProcedures(ctx contractapi.TransactionContextInterface, ownerID string) ([]ehr_asset.Procedure, error) {
	funcName := "ReadProcedures"

	asset, err := s.readRecord(ctx, ownerID)

	if err != nil {
		return nil, chaincodeErrors.NewGenericError(funcName, err)
	}

	return asset.Procedures, nil
}

func (s *EHRContract) recordExists(ctx contractapi.TransactionContextInterface, ownerID string) (bool, error) {
	funcName := "RecordExists"
	response, err := ctx.GetStub().GetState(ownerID)

	if err != nil {
		return false, chaincodeErrors.NewReadWorldStateError(funcName, err)
	}

	return response != nil, nil
}

func isClientInAllowedList(ctx contractapi.TransactionContextInterface, ownerID string) (bool, error) {
	funcName := "IsClientInAllowedList"

	issuerID, err := cid.New(ctx.GetStub())

	issuerCertID, found, err := issuerID.GetAttributeValue("personID")

	if err != nil {
		return false, chaincodeErrors.NewGenericError(funcName, err)
	}

	if !found {
		return false, chaincodeErrors.NewGenericError(funcName, err)
	}

	// Data owner always has access
	if ownerID == issuerCertID {
		return true, nil
	}

	alResponse := ctx.GetStub().InvokeChaincode("AccessList", ToChaincodeArgs("isIdentityApproved", ownerID, issuerCertID), "access_chanel")

	if alResponse.Status != 200 {
		return false, chaincodeErrors.NewGenericError(funcName, nil)
	}

	var isAllowed bool
	err = json.Unmarshal(alResponse.Payload, &isAllowed)

	if err != nil {
		return false, chaincodeErrors.NewMarshallingError(funcName, "IsAllowed", err)
	}

	return isAllowed, nil
}

func ToChaincodeArgs(args ...string) [][]byte {
	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}

	return bargs
}
