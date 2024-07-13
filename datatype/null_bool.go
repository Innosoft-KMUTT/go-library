package datatype

import (
	"database/sql"
	"encoding/json"
	"strconv"
	"strings"
)

// NullBool is a custom type that embeds sql.NullBool
type NullBool struct {
	sql.NullBool
}

// MarshalJSON customizes the JSON marshalling for NullBool
func (nf NullBool) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nf.Bool)
}

// UnmarshalJSON customizes the JSON unmarshalling for NullBool
func (nf *NullBool) UnmarshalJSON(data []byte) error {
	var errString error
	var errBool error

	// Unmarshal the data into a bool
	var val *bool
	if err := json.Unmarshal(data, &val); err != nil {
		errBool = err
	}

	// Unmarshal the data into a string
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		errString = err
	}

	// fmt.Printf("errBool: %v\n", errBool)
	// fmt.Printf("errString: %v\n", errString)

	// fmt.Printf("val: %v\n", val)
	// fmt.Printf("str: %v\n", str)

	if errString != nil && errBool != nil {
		if errString != nil {
			return errString
		}

		if errBool != nil {
			return errBool
		}
	} else if errBool == nil && errString == nil && str == "" && val == nil {
		nf.Valid = false
		return nil
	} else if errBool == nil && val == nil {
		nf.Valid = false
		return nil
	} else {
		nf.Bool = *val
		nf.Valid = true
		return nil
	}

	// If the string is empty, set the value to nil (invalid)
	if str == "" {
		nf.Valid = false
		return nil
	}

	// Parse the string to int64
	value, err := strconv.ParseBool(strings.ToLower(str))
	if err != nil {
		return err
	}

	// Set the value and mark it as valid
	nf.Bool = value
	nf.Valid = true
	return nil
}
