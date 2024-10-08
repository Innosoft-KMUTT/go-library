package datatype

import (
	"database/sql"
	"encoding/json"
	"errors"
	"strconv"
)

// NullUInt is a custom type that embeds sql.NullUInt
type NullUInt struct {
	sql.NullInt64
}

// MarshalJSON customizes the JSON marshalling for NullUInt
func (nf NullUInt) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nf.Int64)
}

// UnmarshalJSON customizes the JSON unmarshalling for NullUInt
func (nf *NullUInt) UnmarshalJSON(data []byte) error {
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

	// fmt.Printf("errBool: %v\n", errInt)
	// fmt.Printf("errString: %v\n", errString)

	// fmt.Printf("val: %v\n", val)
	// fmt.Printf("str: %v\n", str)

	if errString != nil && errInt != nil {
		return errString
	} else if errInt == nil && errString == nil && str == "" && val == nil {
		nf.Valid = false
		return nil
	} else if errInt == nil && val == nil {
		nf.Valid = false
		return nil
	} else if errInt == nil && val != nil {
		if *val < 0 {
			return errors.New("value must greater than zero")
		}
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

	if value < 0 {
		return errors.New("value must greater than zero")
	}

	// Set the value and mark it as valid
	nf.Int64 = value
	nf.Valid = true
	return nil
}
