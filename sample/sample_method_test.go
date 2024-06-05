package sample

import (
	"fmt"
	"testing"
)

func TestSampleMethod(t *testing.T) {
	fmt.Println("TestSampleMethod")

	err := SampleMethod()

	if err != nil {
		t.Error(err)
	}
    
}