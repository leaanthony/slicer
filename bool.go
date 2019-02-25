package slicer

// BoolSlicer handles slices of bools
type BoolSlicer struct {
	slice []bool
}

// Bool creates a new BoolSlicer
func Bool() *BoolSlicer {
	return &BoolSlicer{}
}

// Add a bool value to the slicer
func (s *BoolSlicer) Add(value bool) {
	s.slice = append(s.slice, value)
}

// AddSlice adds a bool slice to the slicer
func (s *BoolSlicer) AddSlice(value []bool) {
	s.slice = append(s.slice, value...)
}

// AsSlice returns the slice
func (s *BoolSlicer) AsSlice() []bool {
	return s.slice
}

// AddSlicer appends a BoolSlicer to the slicer
func (s *BoolSlicer) AddSlicer(value *BoolSlicer) {
	s.slice = append(s.slice, value.AsSlice()...)
}

// Filter the slice based on the given function
func (s *BoolSlicer) Filter(fn func(bool) bool) *BoolSlicer {
	result := &BoolSlicer{}
	for _, elem := range s.slice {
		if fn(elem) {
			result.Add(elem)
		}
	}
	return result
}
