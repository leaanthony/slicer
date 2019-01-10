// Package slicer cotains utility classes for handling slices
package slicer

// Float32Slicer handles slices of float32s
type Float32Slicer struct {
	slice []float32
}

// Float32 creates a new Float32Slicer
func Float32() *Float32Slicer {
	return &Float32Slicer{}
}

// Add a float32 value to the slicer
func (s *Float32Slicer) Add(value float32) {
	s.slice = append(s.slice, value)
}

// AddSlice adds a float32 slice to the slicer
func (s *Float32Slicer) AddSlice(value []float32) {
	s.slice = append(s.slice, value...)
}

// AsSlice returns the slice
func (s *Float32Slicer) AsSlice() []float32 {
	return s.slice
}

// AddSlicer appends a Float32Slicer to the slicer
func (s *Float32Slicer) AddSlicer(value *Float32Slicer) {
	s.slice = append(s.slice, value.AsSlice()...)
}
