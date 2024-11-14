package al_asset

type AccessList struct {
	OwnerID    string   `json:"ownerID"`
	AllowedIDs []string `json:"allowedIDs"`
}
