package test

import (
	"encoding/json"
	"testing"

	"github.com/leaanthony/slicer"
)

func TestInterfaceAdd(t *testing.T) {

	s := slicer.Interface()
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

	s := slicer.Interface()
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

	s := slicer.Interface()
	var a interface{} = 1
	var b interface{} = "hello"
	s.Add(a)
	s.Add(b)

	p := slicer.Interface()
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
