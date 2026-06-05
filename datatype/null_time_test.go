package datatype

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestNullTimeType struct {
	Value NullTime `json:"value"`
}

func TestNullTime(t *testing.T) {
	fmt.Println("TestNullTime")

	var data TestNullTimeType
	var s string

	s = string(`{}`)
	data = TestNullTimeType{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Valid, false, "value Valid should be false")

	s = string(`{"value":null}`)
	data = TestNullTimeType{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Valid, false, "value Valid should be false")

	s = string(`{"value":""}`)
	data = TestNullTimeType{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Valid, false, "value Valid should be false")

	s = string(`{"value":"2024-01-02T03:04:05Z"}`)
	data = TestNullTimeType{}
	json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Equal(t, data.Value.Valid, true, "value Valid should be true")
	assert.Equal(t, data.Value.Time.UTC(), time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC), "value Time should be 2024-01-02T03:04:05Z")

	s = string(`{"value":"not-a-time"}`)
	data = TestNullTimeType{}
	err := json.Unmarshal([]byte(s), &data)
	fmt.Printf("s: %s\n", s)
	fmt.Printf("value: %v\n", data.Value)
	assert.Error(t, err, "invalid time should return an error")
}

func TestNullTimeMarshalJSON(t *testing.T) {
	fmt.Println("TestNullTimeMarshalJSON")

	now := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)

	nf := NullTime{}
	nf.Time = now
	nf.Valid = true
	b, err := json.Marshal(nf)
	fmt.Printf("value: %s\n", string(b))
	assert.NoError(t, err)
	assert.Equal(t, string(b), `"2024-01-02T03:04:05Z"`, "value should marshal to ISO 8601 string")

	nf = NullTime{}
	b, err = json.Marshal(nf)
	fmt.Printf("value: %s\n", string(b))
	assert.NoError(t, err)
	assert.Equal(t, string(b), `null`, "invalid value should marshal to null")
}

func TestNullTimeScan(t *testing.T) {
	fmt.Println("TestNullTimeScan")

	now := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)

	nf := NullTime{}
	err := nf.Scan(now)
	fmt.Printf("value: %v\n", nf)
	assert.NoError(t, err)
	assert.Equal(t, nf.Valid, true, "value Valid should be true")
	assert.Equal(t, nf.Time, now, "value Time should be 2024-01-02T03:04:05Z")

	nf = NullTime{}
	err = nf.Scan(nil)
	fmt.Printf("value: %v\n", nf)
	assert.NoError(t, err)
	assert.Equal(t, nf.Valid, false, "value Valid should be false")

	nf = NullTime{}
	err = nf.Scan("not a time")
	fmt.Printf("value: %v\n", nf)
	assert.Error(t, err, "scanning an invalid type should return an error")
}

func TestNullTimeValue(t *testing.T) {
	fmt.Println("TestNullTimeValue")

	now := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)

	nf := NullTime{}
	nf.Time = now
	nf.Valid = true
	v, err := nf.Value()
	fmt.Printf("value: %v\n", v)
	assert.NoError(t, err)
	assert.Equal(t, v, now, "value should be 2024-01-02T03:04:05Z")

	nf = NullTime{}
	v, err = nf.Value()
	fmt.Printf("value: %v\n", v)
	assert.NoError(t, err)
	assert.Nil(t, v, "invalid value should be nil")
}
