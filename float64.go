package slicer

// Float64Slicer handles slices of float64s
type Float64Slicer struct {
	slice []float64
}

// Float64 creates a new Float64Slicer
func Float64() *Float64Slicer {
	return &Float64Slicer{}
}

// Add a float64 value to the slicer
func (s *Float64Slicer) Add(value float64) {
	s.slice = append(s.slice, value)
}

// AddSlice adds a float64 slice to the slicer
func (s *Float64Slicer) AddSlice(value []float64) {
	s.slice = append(s.slice, value...)
}

// AsSlice returns the slice
func (s *Float64Slicer) AsSlice() []float64 {
	return s.slice
}

// AddSlicer appends a Float64Slicer to the slicer
func (s *Float64Slicer) AddSlicer(value *Float64Slicer) {
	s.slice = append(s.slice, value.AsSlice()...)
}

// Filter the slice based on the given function
func (s *Float64Slicer) Filter(fn func(float64) bool) *Float64Slicer {
	result := &Float64Slicer{}
	for _, elem := range s.slice {
		if fn(elem) {
			result.Add(elem)
		}
	}
	return result
}

// Each runs a function on every element of the slice
func (s *Float64Slicer) Each(fn func(float64)) {
	for _, elem := range s.slice {
		fn(elem)
	}
}
