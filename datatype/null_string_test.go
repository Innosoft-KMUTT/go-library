package datatype

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestNullStringType struct {
	Value NullString `json:"value"`
}

func TestNullString(t *testing.T) {
	fmt.Println("TestNullString")

	var data TestNullStringType
	var s string

	s = string(`{}`)
	data = TestNullStringType{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Valid, false, "value Valid should be false")

	s = string(`{"value":0}`)
	data = TestNullStringType{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Valid, false, "value Valid should be false")

	s = string(`{"value":1}`)
	data = TestNullStringType{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Valid, false, "value Valid should be false")

	s = string(`{"value":""}`)
	data = TestNullStringType{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.String, "", "value Valid should be empty")

	s = string(`{"value":"0"}`)
	data = TestNullStringType{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.String, "0", "value String should be '0'")

	s = string(`{"value":"1"}`)
	data = TestNullStringType{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.String, "1", "value String should be '1'")
}
