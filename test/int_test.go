package test

import (
	"encoding/json"
	"testing"

	"github.com/leaanthony/slicer"
)

func TestIntAdd(t *testing.T) {

	s := slicer.Int()
	s.Add(1)
	s.Add(2)

	expected := "[1,2]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestIntAddSlice(t *testing.T) {

	s := slicer.Int()
	s.Add(1)
	s.Add(2)

	extras := []int{3, 4}

	s.AddSlice(extras)

	expected := "[1,2,3,4]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestIntAddSlicer(t *testing.T) {

	s := slicer.Int()
	s.Add(1)
	s.Add(2)

	p := slicer.Int()
	p.Add(3)
	p.Add(4)

	s.AddSlicer(p)

	expected := "[1,2,3,4]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
