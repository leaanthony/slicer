// Package slicer cotains utility classes for handling slices
package slicer

// Int64Slicer handles slices of int64s
type Int64Slicer struct {
	slice []int64
}

// Int64 creates a new Int64Slicer
func Int64() *Int64Slicer {
	return &Int64Slicer{}
}

// Add a int64 value to the slicer
func (s *Int64Slicer) Add(value int64) {
	s.slice = append(s.slice, value)
}

// AddSlice adds a int64 slice to the slicer
func (s *Int64Slicer) AddSlice(value []int64) {
	s.slice = append(s.slice, value...)
}

// AsSlice returns the slice
func (s *Int64Slicer) AsSlice() []int64 {
	return s.slice
}

// AddSlicer appends a Int64Slicer to the slicer
func (s *Int64Slicer) AddSlicer(value *Int64Slicer) {
	s.slice = append(s.slice, value.AsSlice()...)
}
