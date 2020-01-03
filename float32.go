package slicer

// Float32Slicer handles slices of float32s
type Float32Slicer struct {
	slice []float32
}

// Float32 creates a new Float32Slicer
func Float32(slice ...[]float32) *Float32Slicer {
	if len(slice) > 0 {
		return &Float32Slicer{slice: slice[0]}
	}
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

// Filter the slice based on the given function
func (s *Float32Slicer) Filter(fn func(float32) bool) *Float32Slicer {
	result := &Float32Slicer{}
	for _, elem := range s.slice {
		if fn(elem) {
			result.Add(elem)
		}
	}
	return result
}

// Each runs a function on every element of the slice
func (s *Float32Slicer) Each(fn func(float32)) {
	for _, elem := range s.slice {
		fn(elem)
	}
}

// Contains indicates if the given value is in the slice
func (s *Float32Slicer) Contains(matcher float32) bool {
	result := false
	for _, elem := range s.slice {
		if elem == matcher {
			result = true
		}
	}
	return result
}
