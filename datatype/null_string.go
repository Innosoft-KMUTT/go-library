package datatype

import (
	"database/sql"
	"encoding/json"
)

// NullString is a custom type that embeds sql.NullString
type NullString struct {
	sql.NullString
}

// MarshalJSON customizes the JSON marshalling for NullString
func (nf NullString) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nf.String)
}

// UnmarshalJSON customizes the JSON unmarshalling for NullString
func (nf *NullString) UnmarshalJSON(data []byte) error {

	// Unmarshal the data into a string
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	// Set the value and mark it as valid
	nf.String = str
	nf.Valid = true
	return nil
}
