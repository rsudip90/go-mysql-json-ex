package dbm

import "encoding/json"

const (
	// ErrInvalidJSON error constant string
	ErrInvalidJSON = "JSON data is invalid"
)

// Applicant struct
type Applicant struct {
	AID       int64
	Name      string
	Email     string
	CellPhone string
	Address   string
}

// JSONDoc struct
type JSONDoc struct {
	DocID int64
	Data  json.RawMessage
}

// IsValidJSON checks raw message is valid or not
func IsValidJSON(raw json.RawMessage) bool {
	_, err := json.Marshal(&raw)
	return err == nil
}
