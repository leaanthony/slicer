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

func TestFloat64Filter(t *testing.T) {

	s := Float64()
	s.Add(0.1)
	s.Add(0.4)
	s.Add(0.3)
	s.Add(0.8)
	s.Add(0.5)
	s.Add(1.3)
	s.Add(3.9)
	s.Add(0.4)

	result := s.Filter(func(i float64) bool {
		return i < 0.5
	})

	expected := "[0.1,0.4,0.3,0.4]"
	actual, _ := json.Marshal(result.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
