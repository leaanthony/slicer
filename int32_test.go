package slicer

import (
	"encoding/json"
	"testing"
)

func TestInt32Add(t *testing.T) {

	s := Int32()
	s.Add(1)
	s.Add(2)

	expected := "[1,2]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestInt32AddSlice(t *testing.T) {

	s := Int32()
	s.Add(1)
	s.Add(2)

	extras := []int32{3, 4}

	s.AddSlice(extras)

	expected := "[1,2,3,4]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestInt32AddSlicer(t *testing.T) {

	s := Int32()
	s.Add(1)
	s.Add(2)

	p := Int32()
	p.Add(3)
	p.Add(4)

	s.AddSlicer(p)

	expected := "[1,2,3,4]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestInt32Filter(t *testing.T) {

	s := Int32()
	s.Add(18)
	s.Add(180)
	s.Add(1)
	s.Add(10)
	s.Add(20)
	s.Add(3)
	s.Add(29)

	result := s.Filter(func(i int32) bool {
		return i > 19
	})

	expected := "[180,20,29]"
	actual, _ := json.Marshal(result.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
