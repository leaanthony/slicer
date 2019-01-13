package slicer

import (
	"encoding/json"
	"testing"
)

func TestFloat64Add(t *testing.T) {

	s := Float64()
	s.Add(1.1)
	s.Add(2.5)

	expected := "[1.1,2.5]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestFloat64AddSlice(t *testing.T) {

	s := Float64()
	s.Add(1.1)
	s.Add(2.2)

	extras := []float64{3.3, 4.4}

	s.AddSlice(extras)

	expected := "[1.1,2.2,3.3,4.4]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestFloat64AddSlicer(t *testing.T) {

	s := Float64()
	s.Add(1.1)
	s.Add(2.2)

	p := Float64()
	p.Add(3.3)
	p.Add(4.4)

	s.AddSlicer(p)

	expected := "[1.1,2.2,3.3,4.4]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
