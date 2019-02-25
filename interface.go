package slicer

// InterfaceSlicer handles slices of interface{}
type InterfaceSlicer struct {
	slice []interface{}
}

// Interface creates a new InterfaceSlicer
func Interface() *InterfaceSlicer {
	return &InterfaceSlicer{}
}

// Add a interface{} value to the slicer
func (s *InterfaceSlicer) Add(value interface{}) {
	s.slice = append(s.slice, value)
}

// AddSlice adds a interface{} slice to the slicer
func (s *InterfaceSlicer) AddSlice(value []interface{}) {
	s.slice = append(s.slice, value...)
}

// AsSlice returns the slice
func (s *InterfaceSlicer) AsSlice() []interface{} {
	return s.slice
}

// AddSlicer appends a InterfaceSlicer to the slicer
func (s *InterfaceSlicer) AddSlicer(value *InterfaceSlicer) {
	s.slice = append(s.slice, value.AsSlice()...)
}

// Filter the slice based on the given function
func (s *InterfaceSlicer) Filter(fn func(interface{}) bool) *InterfaceSlicer {
	result := &InterfaceSlicer{}
	for _, elem := range s.slice {
		if fn(elem) {
			result.Add(elem)
		}
	}
	return result
}

// Each runs a function on every element of the slice
func (s *InterfaceSlicer) Each(fn func(interface{})) {
	for _, elem := range s.slice {
		fn(elem)
	}
}
