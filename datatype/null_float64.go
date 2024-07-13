package datatype

import (
	"database/sql"
	"encoding/json"
	"strconv"
)

// NullFloat64 is a custom type that embeds sql.NullFloat64
type NullFloat64 struct {
	sql.NullFloat64
}

// MarshalJSON customizes the JSON marshalling for NullFloat64
func (nf NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nf.Float64)
}

// UnmarshalJSON customizes the JSON unmarshalling for NullFloat64
func (nf *NullFloat64) UnmarshalJSON(data []byte) error {
	var errString error
	var errFloat error

	// Unmarshal the data into a float
	var val *float64
	if err := json.Unmarshal(data, &val); err != nil {
		errFloat = err
	}

	// Unmarshal the data into a string
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		errString = err
	}

	if errString != nil && errFloat != nil {
		return errString
	} else if errFloat == nil && errString == nil && str == "" && val == nil {
		nf.Valid = false
		return nil
	} else if errFloat == nil && val == nil {
		nf.Valid = false
		return nil
	} else if errFloat == nil && val != nil {
		nf.Float64 = *val
		nf.Valid = true
		return nil
	}

	// If the string is empty, set the value to nil (invalid)
	if str == "" {
		nf.Valid = false
		return nil
	}

	// Parse the string to float64
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return err
	}

	// Set the value and mark it as valid
	nf.Float64 = value
	nf.Valid = true
	return nil
}
