package slicer

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestFloat32Add(t *testing.T) {

	s := Float32()
	s.Add(1.1)
	s.Add(2.5)

	expected := "[1.1,2.5]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestFloat32AddSlice(t *testing.T) {

	s := Float32()
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

	s := Float32()
	s.Add(1.1)
	s.Add(2.2)

	p := Float32()
	p.Add(3.3)
	p.Add(4.4)

	s.AddSlicer(p)

	expected := "[1.1,2.2,3.3,4.4]"
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestFloat32Filter(t *testing.T) {

	s := Float32()
	s.Add(0.1)
	s.Add(0.4)
	s.Add(0.3)
	s.Add(0.8)
	s.Add(0.5)
	s.Add(1.3)
	s.Add(3.9)
	s.Add(0.4)

	result := s.Filter(func(i float32) bool {
		return i < 0.5
	})

	expected := "[0.1,0.4,0.3,0.4]"
	actual, _ := json.Marshal(result.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestFloat32Each(t *testing.T) {

	s := Float32()
	s.Add(0.1)
	s.Add(0.4)
	s.Add(0.3)
	s.Add(0.8)
	s.Add(0.5)
	s.Add(1.3)
	s.Add(3.9)
	s.Add(0.4)

	var result float32
	s.Each(func(i float32) {
		result += i
	})

	var expected = 7.7
	if fmt.Sprintf("%f", expected) != fmt.Sprintf("%f", result) {
		t.Errorf("Expected '%f', but got '%f'", expected, result)
	}
}
