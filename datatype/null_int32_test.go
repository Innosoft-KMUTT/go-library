package datatype

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestNullInt32Type struct {
	Value NullInt32 `json:"value"`
}

func TestNullInt32(t *testing.T) {
	fmt.Println("TestNullInt32")

	var data TestNullInt32Type
	var s string

	s = string(`{}`)
	data = TestNullInt32Type{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Valid, false, "value Valid should be false")

	s = string(`{"value":0}`)
	data = TestNullInt32Type{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Int32, int32(0), "value Int32 should be 0")

	s = string(`{"value":1}`)
	data = TestNullInt32Type{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Int32, int32(1), "value Int32 should be 1")

	s = string(`{"value":""}`)
	data = TestNullInt32Type{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Valid, false, "value Valid should be false")

	s = string(`{"value":"0"}`)
	data = TestNullInt32Type{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Int32, int32(0), "value Int32 should be 0")

	s = string(`{"value":"1"}`)
	data = TestNullInt32Type{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Int32, int32(1), "value Int32 should be 1.1")
}
