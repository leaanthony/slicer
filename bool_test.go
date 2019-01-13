package slicer

import (
	"encoding/json"
	"testing"
)

func TestBoolAdd(t *testing.T) {

	s := Bool()
	s.Add(true)
	s.Add(false)

	expected := "[true,false]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestBoolAddSlice(t *testing.T) {

	s := Bool()
	s.Add(true)
	s.Add(false)

	extras := []bool{true, false}

	s.AddSlice(extras)

	expected := "[true,false,true,false]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestBoolAddSlicer(t *testing.T) {

	s := Bool()
	s.Add(true)
	s.Add(false)

	p := Bool()
	p.Add(true)
	p.Add(false)

	s.AddSlicer(p)

	expected := "[true,false,true,false]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
