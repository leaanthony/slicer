package slicer

import (
	"encoding/json"
	"testing"
)

func TestInt8Add(t *testing.T) {

	s := Int8()
	s.Add(1)
	s.Add(2)

	expected := "[1,2]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestInt8AddSlice(t *testing.T) {

	s := Int8()
	s.Add(1)
	s.Add(2)

	extras := []int8{3, 4}

	s.AddSlice(extras)

	expected := "[1,2,3,4]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestInt8AddSlicer(t *testing.T) {

	s := Int8()
	s.Add(1)
	s.Add(2)

	p := Int8()
	p.Add(3)
	p.Add(4)

	s.AddSlicer(p)

	expected := "[1,2,3,4]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
