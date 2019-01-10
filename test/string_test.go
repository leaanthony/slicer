package test

import (
	"strings"
	"testing"

	"github.com/leaanthony/slicer"
)

func TestStringAdd(t *testing.T) {

	s := slicer.String()
	s.Add("one")
	s.Add("two")

	expected := "one two"
	actual := strings.Join(s.AsSlice(), " ")
	if expected != actual {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestStringAddSlice(t *testing.T) {

	s := slicer.String()
	s.Add("one")
	s.Add("two")
	extras := []string{"three", "four"}
	s.AddSlice(extras)

	expected := "one two three four"
	actual := strings.Join(s.AsSlice(), " ")
	if expected != actual {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestStringAddSlicer(t *testing.T) {

	s := slicer.String()
	s.Add("one")
	s.Add("two")

	p := slicer.String()
	p.Add("three")
	p.Add("four")

	s.AddSlicer(p)

	expected := "one two three four"
	actual := strings.Join(s.AsSlice(), " ")
	if expected != actual {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
