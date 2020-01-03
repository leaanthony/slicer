package slicer

import (
	"fmt"
	"strings"
)

// Int16Slicer handles slices of int16s
type Int16Slicer struct {
	slice []int16
}

// Int16 creates a new Int16Slicer
func Int16(slice ...[]int16) *Int16Slicer {
	if len(slice) > 0 {
		return &Int16Slicer{slice: slice[0]}
	}
	return &Int16Slicer{}
}

// Add a int16 value to the slicer
func (s *Int16Slicer) Add(value int16) {
	s.slice = append(s.slice, value)
}

// AddSlice adds a int16 slice to the slicer
func (s *Int16Slicer) AddSlice(value []int16) {
	s.slice = append(s.slice, value...)
}

// AsSlice returns the slice
func (s *Int16Slicer) AsSlice() []int16 {
	return s.slice
}

// AddSlicer appends a Int16Slicer to the slicer
func (s *Int16Slicer) AddSlicer(value *Int16Slicer) {
	s.slice = append(s.slice, value.AsSlice()...)
}

// Filter the slice based on the given function
func (s *Int16Slicer) Filter(fn func(int16) bool) *Int16Slicer {
	result := &Int16Slicer{}
	for _, elem := range s.slice {
		if fn(elem) {
			result.Add(elem)
		}
	}
	return result
}

// Each runs a function on every element of the slice
func (s *Int16Slicer) Each(fn func(int16)) {
	for _, elem := range s.slice {
		fn(elem)
	}
}

// Contains indicates if the given value is in the slice
func (s *Int16Slicer) Contains(matcher int16) bool {
	result := false
	for _, elem := range s.slice {
		if elem == matcher {
			result = true
		}
	}
	return result
}

// Join returns a string with the slicer elements separated by the given separator
func (s *Int16Slicer) Join(separator string) string {
	var builder strings.Builder

	// Shortcut no elements
	if len(s.slice) == 0 {
		return ""
	}

	// Iterate over length - 1
	index := 0
	for index = 0; index < len(s.slice)-1; index++ {
		builder.WriteString(fmt.Sprintf("%v%s", s.slice[index], separator))
	}
	builder.WriteString(fmt.Sprintf("%v", s.slice[index]))
	result := builder.String()
	return result
}

// Length returns the number of elements in the slice
func (s *Int16Slicer) Length() int {
	return len(s.slice)
}
