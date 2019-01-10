package test

import (
	"encoding/json"
	"testing"

	"github.com/leaanthony/slicer"
)

func TestFloat32Add(t *testing.T) {

	s := slicer.Float32()
	s.Add(1.1)
	s.Add(2.5)

	expected := "[1.1,2.5]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestFloat32AddSlice(t *testing.T) {

	s := slicer.Float32()
	s.Add(1.1)
	s.Add(2.2)

	extras := []float32{3.3, 4.4}

	s.AddSlice(extras)

	expected := "[1.1,2.2,3.3,4.4]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestFloat32AddSlicer(t *testing.T) {

	s := slicer.Float32()
	s.Add(1.1)
	s.Add(2.2)

	p := slicer.Float32()
	p.Add(3.3)
	p.Add(4.4)

	s.AddSlicer(p)

	expected := "[1.1,2.2,3.3,4.4]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
