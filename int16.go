// Package slicer cotains utility classes for handling slices
package slicer

// Int16Slicer handles slices of int16s
type Int16Slicer struct {
	slice []int16
}

// Int16 creates a new Int16Slicer
func Int16() *Int16Slicer {
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
