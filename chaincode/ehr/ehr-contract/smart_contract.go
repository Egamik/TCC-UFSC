package ehr_contract

import (
	"ehr_asset"
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

type EHRContract struct {
	contractapi.Contract
}

func (s *EHRContract) createRecord(ctx contractapi.TransactionContextInterface, ownerID string) error {
	// Only professionals should be able to create records
	exists := s.recordExists(ctx, ownerID)

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

	err = ctx.GetStub().PutState(ownerID, assetBytes)

	return err
}

func (s *EHRContract) addPrescription(ctx contractapi.TransactionContextInterface, ownerID string, prescriptionJSON string) error {
	// Check if record exists
	exists := s.recordExists(ctx, ownerID)

	if !exists {
		return nil
	}

	// Read record from world state
	recordBytes, err := ctx.GetStub().GetState(ownerID)

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

func (s *EHRContract) addAppointment(ctx contractapi.TransactionContextInterface, ownerID string, appointmentJSON string) error {
	return nil
}

func (s *EHRContract) addProcedure(ctx contractapi.TransactionContextInterface, ownerID string, procedureJSON string) error {
	return nil
}

// Read owner's complete record
func (s *EHRContract) readRecord(ctx contractapi.TransactionContextInterface, ownerID string) (ehr_asset.EHR_Asset, error) {
	var record ehr_asset.EHR_Asset
	return record, nil
}

// Read all prescriptions given by professional
func (s *EHRContract) readPrescriptions(ctx contractapi.TransactionContextInterface, ownerID string, professionalID string) ([]ehr_asset.Prescription, error) {
	return nil, nil
}

// Read all appointments by professional
func (s *EHRContract) readAppointments(ctx contractapi.TransactionContextInterface, ownerID string, professionalID string) ([]ehr_asset.Appointment, error) {
	return nil, nil
}

// Real all procedures by professional
func (s *EHRContract) readProcedures(ctx contractapi.TransactionContextInterface, ownerID string, professionalID string) ([]ehr_asset.Procedure, error) {
	return nil, nil
}

func (s *EHRContract) recordExists(ctx contractapi.TransactionContextInterface, ownerID string) bool {

	response, err := ctx.GetStub().GetState(ownerID)

	if err != nil {
		return false
	}

	return response != nil
}
