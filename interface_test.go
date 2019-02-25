package slicer

import (
	"encoding/json"
	"testing"
)

func TestInterfaceAdd(t *testing.T) {

	s := Interface()
	var a interface{} = 1
	var b interface{} = "hello"
	s.Add(a)
	s.Add(b)

	expected := `[1,"hello"]`
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestInterfaceAddSlice(t *testing.T) {

	s := Interface()
	var a interface{} = 1
	var b interface{} = "hello"
	s.Add(a)
	s.Add(b)

	extras := []interface{}{true, 6.6}

	s.AddSlice(extras)

	expected := `[1,"hello",true,6.6]`
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
func TestInterfaceAddSlicer(t *testing.T) {

	s := Interface()
	var a interface{} = 1
	var b interface{} = "hello"
	s.Add(a)
	s.Add(b)

	p := Interface()
	var c interface{} = true
	var d interface{} = 6.6
	p.Add(c)
	p.Add(d)

	s.AddSlicer(p)

	expected := `[1,"hello",true,6.6]`
	actual, _ := json.Marshal(s.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestInterfaceFilter(t *testing.T) {

	s := Interface()
	s.Add(1)
	s.Add("hello")
	s.Add(true)
	s.Add("world")
	s.Add(9.9)

	result := s.Filter(func(i interface{}) bool {
		_, ok := i.(string)
		return ok
	})

	expected := `["hello","world"]`
	actual, _ := json.Marshal(result.AsSlice())
	if expected != string(actual) {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
