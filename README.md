
<div style="text-align:center; width:400px">
  <img src="logo.png"/>
  Utility class for handling slices.
</div>


[![Go Report Card](https://goreportcard.com/badge/github.com/leaanthony/slicer)](https://goreportcard.com/report/github.com/leaanthony/slicer)  [![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/leaanthony/slicer) [![CodeFactor](https://www.codefactor.io/repository/github/leaanthony/slicer/badge)](https://www.codefactor.io/repository/github/leaanthony/slicer) ![](https://img.shields.io/bower/l/svg) [![cover.run](https://cover.run/go/github.com/leaanthony/slicer.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2Fleaanthony%2Fslicer) [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)  





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
  
