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
