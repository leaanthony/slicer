package slicer

// Int8Slicer handles slices of int8s
type Int8Slicer struct {
	slice []int8
}

// Int8 creates a new Int8Slicer
func Int8() *Int8Slicer {
	return &Int8Slicer{}
}

// Add a int8 value to the slicer
func (s *Int8Slicer) Add(value int8) {
	s.slice = append(s.slice, value)
}

// AddSlice adds a int8 slice to the slicer
func (s *Int8Slicer) AddSlice(value []int8) {
	s.slice = append(s.slice, value...)
}

// AsSlice returns the slice
func (s *Int8Slicer) AsSlice() []int8 {
	return s.slice
}

// AddSlicer appends a Int8Slicer to the slicer
func (s *Int8Slicer) AddSlicer(value *Int8Slicer) {
	s.slice = append(s.slice, value.AsSlice()...)
}

// Filter the slice based on the given function
func (s *Int8Slicer) Filter(fn func(int8) bool) *Int8Slicer {
	result := &Int8Slicer{}
	for _, elem := range s.slice {
		if fn(elem) {
			result.Add(elem)
		}
	}
	return result
}
