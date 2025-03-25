package ehr_asset

import (
	"encoding/json"
	"time"
)

type EHR_Asset struct {
	PatientID     string         `json:"patientID"`
	Prescriptions []Prescription `json:"prescriptions"`
	Appointments  []Appointment  `json:"appointments"`
	Procedures    []Procedure    `json:"procedures"`
}

type Prescription struct {
	ProfessionalID string    `json:"professionalID"`
	Date           time.Time `json:"date"`
	Description    string    `json:"description"`
	Medication     string    `json:"medication"`
	Dosage         string    `json:"dosage"`
}

type Appointment struct {
	ProfessionalID string    `json:"professionalID"`
	Date           time.Time `json:"date"`
	ClinicName     string    `json:"clinicName"`
}

type Procedure struct {
	ProfessionalID       string    `json:"professionID"`
	Date                 time.Time `json:"date"`
	ProcedureID          string    `json:"procedureID"`
	ProcedurePlace       string    `json:"procedurePlace"`
	RelatedProfessionals []string  `json:"relatedProfessionals"`
}

/********************************************
 * 		Prescription functions				*
 *******************************************/
// Validates prescription JSON string and returns struct if valid
func (p *Prescription) validatePrescription(prescriptionJSON string) bool {

	if prescriptionJSON == "" {
		return false
	}

	var prescription Prescription
	err := json.Unmarshal([]byte(prescriptionJSON), prescription)

	if err != nil {
		return false
	}

	// Validar campos
	return true
}

/********************************************
 * 		Appointment functions				*
 *******************************************/
// Validates appointment JSON string and returns struct if valid
func validateAppointment(appointmentJSON string) bool {
	if appointmentJSON == "" {
		return false
	}

	var appointment Appointment

	err := json.Unmarshal([]byte(appointmentJSON), appointment)

	if err != nil {
		return false
	}

	return true
}

/********************************************
 * 		Procedure functions 				*
 *******************************************/
// Validates procedure JSON string and returns struct if valid
func validateProcedure(procedureJSON string) bool {

	if procedureJSON == "" {
		return false
	}

	var procedure Procedure
	err := json.Unmarshal([]byte(procedureJSON), procedure)

	if err != nil {
		return false
	}

	return true
}
