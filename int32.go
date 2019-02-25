package slicer

// Int32Slicer handles slices of int32s
type Int32Slicer struct {
	slice []int32
}

// Int32 creates a new Int32Slicer
func Int32() *Int32Slicer {
	return &Int32Slicer{}
}

// Add a int32 value to the slicer
func (s *Int32Slicer) Add(value int32) {
	s.slice = append(s.slice, value)
}

// AddSlice adds a int32 slice to the slicer
func (s *Int32Slicer) AddSlice(value []int32) {
	s.slice = append(s.slice, value...)
}

// AsSlice returns the slice
func (s *Int32Slicer) AsSlice() []int32 {
	return s.slice
}

// AddSlicer appends a Int32Slicer to the slicer
func (s *Int32Slicer) AddSlicer(value *Int32Slicer) {
	s.slice = append(s.slice, value.AsSlice()...)
}

// Filter the slice based on the given function
func (s *Int32Slicer) Filter(fn func(int32) bool) *Int32Slicer {
	result := &Int32Slicer{}
	for _, elem := range s.slice {
		if fn(elem) {
			result.Add(elem)
		}
	}
	return result
}

// Each runs a function on every element of the slice
func (s *Int32Slicer) Each(fn func(int32)) {
	for _, elem := range s.slice {
		fn(elem)
	}
}
