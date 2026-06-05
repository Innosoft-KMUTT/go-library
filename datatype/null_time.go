package datatype

import (
	"database/sql"
	"encoding/json"
	"time"
)

// NullTime is a custom type that embeds sql.NullTime
type NullTime struct {
	sql.NullTime
}

// MarshalJSON customizes the JSON marshalling for NullTime
func (nf NullTime) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nf.Time)
}

// UnmarshalJSON customizes the JSON unmarshalling for NullTime
func (nf *NullTime) UnmarshalJSON(data []byte) error {

	// Unmarshal the data into a string
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	// If the string is empty, set the value to nil (invalid)
	if str == "" {
		nf.Valid = false
		return nil
	}

	// Parse the string to time.Time
	value, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}

	// Set the value and mark it as valid
	nf.Time = value
	nf.Valid = true
	return nil
}
