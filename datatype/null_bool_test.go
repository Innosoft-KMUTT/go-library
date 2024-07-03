package datatype

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestNullBoolType struct {
	Value NullBool `json:"value"`
}

func TestNullBool(t *testing.T) {
	fmt.Println("TestNullBool")

	var data TestNullBoolType
	var s string

	s = string(`{}`)
	data = TestNullBoolType{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Valid, false, "value Valid should be false")

	s = string(`{"value":true}`)
	data = TestNullBoolType{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Bool, true, "value Int64 should be true")

	s = string(`{"value":false}`)
	data = TestNullBoolType{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Bool, false, "value Int64 should be false")

	s = string(`{"value":""}`)
	data = TestNullBoolType{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Valid, false, "value Valid should be false")

	s = string(`{"value":"true"}`)
	data = TestNullBoolType{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Bool, true, "value Int64 should be true")

	s = string(`{"value":"True"}`)
	data = TestNullBoolType{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Bool, true, "value Int64 should be true")

	s = string(`{"value":"false"}`)
	data = TestNullBoolType{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Bool, false, "value Int64 should be false")

	s = string(`{"value":"False"}`)
	data = TestNullBoolType{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Bool, false, "value Int64 should be false")
}
