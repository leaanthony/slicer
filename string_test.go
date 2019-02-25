package slicer

import (
	"strings"
	"testing"
)

func TestStringAdd(t *testing.T) {

	s := String()
	s.Add("one")
	s.Add("two")

	expected := "one two"
	actual := strings.Join(s.AsSlice(), " ")
	if expected != actual {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestStringAddSlice(t *testing.T) {

	s := String()
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

	s := String()
	s.Add("one")
	s.Add("two")

	p := String()
	p.Add("three")
	p.Add("four")

	s.AddSlicer(p)

	expected := "one two three four"
	actual := strings.Join(s.AsSlice(), " ")
	if expected != actual {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestStringFilter(t *testing.T) {

	s := String()
	s.Add("one")
	s.Add("two")
	s.Add("one")
	s.Add("three")
	s.Add("one")

	expected := "one one one"
	filtered := s.Filter(func(v string) bool {
		return v == "one"
	})
	actual := strings.Join(filtered.AsSlice(), " ")
	if expected != actual {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
