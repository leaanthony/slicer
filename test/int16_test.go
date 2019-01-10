package test

import (
	"encoding/json"
	"testing"

	"github.com/leaanthony/slicer"
)

func TestInt16Add(t *testing.T) {

	s := slicer.Int16()
	s.Add(1)
	s.Add(2)

	expected := "[1,2]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestInt16AddSlice(t *testing.T) {

	s := slicer.Int16()
	s.Add(1)
	s.Add(2)

	extras := []int16{3, 4}

	s.AddSlice(extras)

	expected := "[1,2,3,4]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestInt16AddSlicer(t *testing.T) {

	s := slicer.Int16()
	s.Add(1)
	s.Add(2)

	p := slicer.Int16()
	p.Add(3)
	p.Add(4)

	s.AddSlicer(p)

	expected := "[1,2,3,4]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
