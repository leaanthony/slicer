// Package slicer cotains utility classes for handling slices
package slicer

// IntSlicer handles slices of ints
type IntSlicer struct {
	slice []int
}

// Int creates a new IntSlicer
func Int() *IntSlicer {
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
