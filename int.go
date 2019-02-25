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
