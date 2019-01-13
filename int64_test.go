package slicer

import (
	"encoding/json"
	"testing"
)

func TestInt64Add(t *testing.T) {

	s := Int64()
	s.Add(1)
	s.Add(2)

	expected := "[1,2]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestInt64AddSlice(t *testing.T) {

	s := Int64()
	s.Add(1)
	s.Add(2)

	extras := []int64{3, 4}

	s.AddSlice(extras)

	expected := "[1,2,3,4]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestInt64AddSlicer(t *testing.T) {

	s := Int64()
	s.Add(1)
	s.Add(2)

	p := Int64()
	p.Add(3)
	p.Add(4)

	s.AddSlicer(p)

	expected := "[1,2,3,4]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
