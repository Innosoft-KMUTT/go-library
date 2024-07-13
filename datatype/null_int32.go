package datatype

import (
	"database/sql"
	"encoding/json"
	"strconv"
)

// NullInt32 is a custom type that embeds sql.NullInt32
type NullInt32 struct {
	sql.NullInt32
}

// MarshalJSON customizes the JSON marshalling for NullInt32
func (nf NullInt32) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nf.Int32)
}

// UnmarshalJSON customizes the JSON unmarshalling for NullInt32
func (nf *NullInt32) UnmarshalJSON(data []byte) error {
	var errString error
	var errInt error

	// Unmarshal the data into a int32
	var val *int32
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
		nf.Int32 = *val
		nf.Valid = true
		return nil
	}

	// If the string is empty, set the value to nil (invalid)
	if str == "" {
		nf.Valid = false
		return nil
	}

	// Parse the string to int32
	value, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return err
	}

	// Set the value and mark it as valid
	nf.Int32 = int32(value)
	nf.Valid = true
	return nil
}
