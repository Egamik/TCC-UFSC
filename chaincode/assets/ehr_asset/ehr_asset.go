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
func (p *Prescription) validatePrescription(prescriptionJSON string) (string, error) {

	if prescriptionJSON == "" {
		return "", nil
	}

	var prescription Prescription
	err := json.Unmarshal([]byte(prescriptionJSON), prescription)

	if err != nil {
		return "", err
	}

	// Validar campos
	return "", nil
}

/********************************************
 * 		Appointment functions				*
 *******************************************/
// Validates appointment JSON string and returns struct if valid
func (a *Appointment) validateAppointment(appointmentJSON string) (string, error) {
	if appointmentJSON == "" {
		return "", nil
	}

	var appointment Appointment

	err := json.Unmarshal([]byte(appointmentJSON), appointment)

	if err != nil {
		return "", nil
	}

	return "", nil
}

/********************************************
 * 		Procedure functions 				*
 *******************************************/
// Validates procedure JSON string and returns struct if valid
func (p *Procedure) validateProcedure(procedureJSON string) (string, error) {

	if procedureJSON == "" {
		return "", nil
	}

	var procedure Procedure
	err := json.Unmarshal([]byte(procedureJSON), procedure)

	if err != nil {
		return "", nil
	}

	return "", nil
}
