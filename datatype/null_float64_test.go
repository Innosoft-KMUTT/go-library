package datatype

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestNullFloat64Type struct {
	Value NullFloat64 `json:"value"`
}

func TestNullFloat64(t *testing.T) {
	fmt.Println("TestNullFloat64")

	var data TestNullFloat64Type
	var s string

	s = string(`{}`)
	data = TestNullFloat64Type{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Valid, false, "value Valid should be false")

	s = string(`{"value":0}`)
	data = TestNullFloat64Type{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Float64, 0.0, "value Float64 should be 0")

	s = string(`{"value":1.1}`)
	data = TestNullFloat64Type{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Float64, 1.1, "value Float64 should be 1.1")

	s = string(`{"value":""}`)
	data = TestNullFloat64Type{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Valid, false, "value Valid should be false")

	s = string(`{"value":"0"}`)
	data = TestNullFloat64Type{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Float64, 0.0, "value Float64 should be 0")

	s = string(`{"value":"1.1"}`)
	data = TestNullFloat64Type{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Float64, 1.1, "value Float64 should be 1.1")
}
