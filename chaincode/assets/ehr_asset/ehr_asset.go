package ehr_asset

import (
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

// Validates prescription JSON string and returns struct if valid
func validatePrescription(prescriptionJSON string) (*Prescription, error) {

	if prescriptionJSON == "" {
		return nil, nil
	}

	prescription := Prescription{
		ProfessionalID: "",
		Date:           time.Now(),
		Description:    "",
		Medication:     "",
	}

	return &prescription, nil
}

// Validates appointment JSON string and returns struct if valid
func validateAppointment(appointmentJSON string) (*Appointment, error) {
	if appointmentJSON == "" {
		return nil, nil
	}

	appointment := Appointment{
		ProfessionalID: "",
		Date:           time.Now(),
		ClinicName:     "",
	}

	return &appointment, nil
}

// Validates procedure JSON string and returns struct if valid
func validateProcedure(procedureJSON string) (*Procedure, error) {

	if procedureJSON == "" {
		return nil, nil
	}

	procedure := Procedure{
		ProfessionalID:       "",
		Date:                 time.Now(),
		ProcedureID:          "",
		RelatedProfessionals: []string{},
	}

	return &procedure, nil
}
