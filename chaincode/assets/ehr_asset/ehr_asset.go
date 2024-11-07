package ehr_asset

import (
	"time"
)

type EHR_Asset struct {
	PatientID     string          `json:"patientID"`
	Prescriptions []Prescriptions `json:"prescriptions"`
	Appointments  []Appointments  `json:"appointments"`
}

type Prescriptions struct {
	ProfessionalID string    `json:"professionalID"`
	Date           time.Time `json:"date"`
	Description    string    `json:"description"`
	Medication     string    `json:"medication"`
}

type Appointments struct {
	ProfessionalID string    `json:"professionalID"`
	Date           time.Time `json:"date"`
	ClinicName     string    `json:"clinicName"`
}

func (p *Prescriptions) validatePrescription() bool {
	return true
}

func (a *Appointments) validateAppointment() bool {
	return true
}
