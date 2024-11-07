package al_asset

type AL_Asset struct {
	PatientID  string   `json:"patientID"`
	AllowedIDs []string `json:"allowedIDs"`
}
