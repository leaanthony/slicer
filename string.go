// Package slicer cotains utility classes for handling slices
package slicer

// StringSlicer handles slices of strings
type StringSlicer struct {
	slice []string
}

// String creates a new StringSlicer
func String() *StringSlicer {
	return &StringSlicer{}
}

// Add a string value to the slicer
func (s *StringSlicer) Add(value string) {
	s.slice = append(s.slice, value)
}

// AddSlice adds a string slice to the slicer
func (s *StringSlicer) AddSlice(value []string) {
	s.slice = append(s.slice, value...)
}

// AsSlice returns the slice
func (s *StringSlicer) AsSlice() []string {
	return s.slice
}

// AddSlicer appends a StringSlicer to the slicer
func (s *StringSlicer) AddSlicer(value *StringSlicer) {
	s.slice = append(s.slice, value.AsSlice()...)
}

// Filter the slice based on the given function
func (s *StringSlicer) Filter(fn func(string) bool) *StringSlicer {
	result := &StringSlicer{}
	for _, elem := range s.slice {
		if fn(elem) {
			result.Add(elem)
		}
	}
	return result
}

// Each runs a function on every element of the slice
func (s *StringSlicer) Each(fn func(string)) {
	for _, elem := range s.slice {
		fn(elem)
	}
}
