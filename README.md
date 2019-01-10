
<div style="text-align:center; width:400px">
  <img src="logo.png"/>
  Utility class for handling slices.
</div>


[![Go Report Card](https://goreportcard.com/badge/github.com/leaanthony/slicer)](https://goreportcard.com/report/github.com/leaanthony/slicer)  [![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)]
  (http://godoc.org/github.com/leaanthony/slicer) [![CodeFactor](https://www.codefactor.io/repository/github/leaanthony/slicer/badge)](https://www.codefactor.io/repository/github/leaanthony/slicer) ![](https://img.shields.io/bower/l/slicer.svg)


## Install

`go get -u github.com/leaanthony/slicer`

## Quick Start

```
  import "github.com/leaanthony/slicer"

  func test() {
    s := slicer.String()
    s.Add("one")
    s.Add("two")
    s.AddSlice([]string{"three","four"})
    fmt.Printf("My slice = %+v\n", s.AsSlice())
    
    t := slicer.String()
    t.Add("zero")
    t.AddSlicer(s)
    fmt.Printf("My slice = %+v\n", t.AsSlice())
  }
```

## Available slicers

  - Int
  - Int8
  - Int16
  - Int32
  - Float32
  - Float64
  - String
  - Bool
  - Interface
  