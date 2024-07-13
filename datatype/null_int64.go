package datatype

import (
	"database/sql"
	"encoding/json"
	"strconv"
)

// NullInt64 is a custom type that embeds sql.NullInt64
type NullInt64 struct {
	sql.NullInt64
}

// MarshalJSON customizes the JSON marshalling for NullInt64
func (nf NullInt64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nf.Int64)
}

// UnmarshalJSON customizes the JSON unmarshalling for NullInt64
func (nf *NullInt64) UnmarshalJSON(data []byte) error {
	var errString error
	var errInt error

	// Unmarshal the data into a int64
	var val *int64
	if err := json.Unmarshal(data, &val); err != nil {
		errInt = err
	}

	// Unmarshal the data into a string
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		errString = err
	}

	if errString != nil && errInt != nil {
		return errString
	} else if errInt == nil && errString == nil && str == "" && val == nil {
		nf.Valid = false
		return nil
	} else if errInt == nil && val == nil {
		nf.Valid = false
		return nil
	} else if errInt == nil && val != nil {
		nf.Int64 = *val
		nf.Valid = true
		return nil
	}

	// If the string is empty, set the value to nil (invalid)
	if str == "" {
		nf.Valid = false
		return nil
	}

	// Parse the string to int64
	value, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}

	// Set the value and mark it as valid
	nf.Int64 = value
	nf.Valid = true
	return nil
}
