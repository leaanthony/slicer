package slicer

import (
	"fmt"
	"strings"
)

// IntSlicer handles slices of ints
type IntSlicer struct {
	slice []int
}

// Int creates a new IntSlicer
func Int(slice ...[]int) *IntSlicer {
	if len(slice) > 0 {
		return &IntSlicer{slice: slice[0]}
	}
	return &IntSlicer{}
}

// Add a int value to the slicer
func (s *IntSlicer) Add(value int) {
	s.slice = append(s.slice, value)
}

// AddSlice adds a int slice to the slicer
func (s *IntSlicer) AddSlice(value []int) {
	s.slice = append(s.slice, value...)
}

// AsSlice returns the slice
func (s *IntSlicer) AsSlice() []int {
	return s.slice
}

// AddSlicer appends a IntSlicer to the slicer
func (s *IntSlicer) AddSlicer(value *IntSlicer) {
	s.slice = append(s.slice, value.AsSlice()...)
}

// Filter the slice based on the given function
func (s *IntSlicer) Filter(fn func(int) bool) *IntSlicer {
	result := &IntSlicer{}
	for _, elem := range s.slice {
		if fn(elem) {
			result.Add(elem)
		}
	}
	return result
}

// Each runs a function on every element of the slice
func (s *IntSlicer) Each(fn func(int)) {
	for _, elem := range s.slice {
		fn(elem)
	}
}

// Contains indicates if the given value is in the slice
func (s *IntSlicer) Contains(matcher int) bool {
	result := false
	for _, elem := range s.slice {
		if elem == matcher {
			result = true
		}
	}
	return result
}

// Join returns a string with the slicer elements separated by the given separator
func (s *IntSlicer) Join(separator string) string {
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
func (s *IntSlicer) Length() int {
	return len(s.slice)
}
