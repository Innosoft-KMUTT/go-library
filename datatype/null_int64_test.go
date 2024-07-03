package datatype

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestNullInt64Type struct {
	Value NullInt64 `json:"value"`
}

func TestNullInt64(t *testing.T) {
	fmt.Println("TestNullInt64")

	var data TestNullInt64Type
	var s string

	s = string(`{}`)
	data = TestNullInt64Type{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Valid, false, "value Valid should be false")

	s = string(`{"value":0}`)
	data = TestNullInt64Type{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Int64, int64(0), "value Int64 should be 0")

	s = string(`{"value":1}`)
	data = TestNullInt64Type{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Int64, int64(1), "value Int64 should be 1")

	s = string(`{"value":""}`)
	data = TestNullInt64Type{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Valid, false, "value Valid should be false")

	s = string(`{"value":"0"}`)
	data = TestNullInt64Type{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Int64, int64(0), "value Int64 should be 0")

	s = string(`{"value":"1"}`)
	data = TestNullInt64Type{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Int64, int64(1), "value Int64 should be 1.1")
}
