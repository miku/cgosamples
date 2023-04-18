# Cgo experiments

Lightning Talk at [Leipzig Gophers](https://golangleipzig.space) [#35](#), 2023-04-18 1900 CEST, Basislager Leipzig

## Cgo

> Cgo enables the creation of Go packages that call C code.

## Pseudo-package "C"

* compilation takes a different code path
* allows to include C code directly in go code and to reference C symbols from Go

```go
// #include <stdio.h>
// #include <errno.h>
import "C"
```

That is called the preamble. No newlines allowed.

## Types

> Cgo recognizes any use of a qualified identifier C.xxx and uses gcc to
find the definition of xxx.  If xxx is a type, cgo replaces C.xxx with a Go
translation.  C arithmetic types translate to precisely-sized Go arithmetic
types.  A C struct translates to a Go struct, field by field; unrepresentable
fields are replaced with opaque byte arrays.  A C union translates into a
struct containing the first union member and perhaps additional padding.  C
arrays become Go arrays.  C pointers become Go pointers.  C function pointers
become Go's uintptr.  C void pointers become Go's unsafe.Pointer.


## Basic example

> [x/xcgobasic](x/xcgobasic)

* preamble, with function
* stdlib

Helper functions:

> The C string is allocated in the C heap using malloc. -- [src/cmd/cgo/out.go;l=1640](https://cs.opensource.google/go/go/+/refs/tags/go1.20.3:src/cmd/cgo/out.go;l=1640)

## The cgo tool

> [x/xcgotool](x/xcgotool)

## Pointers

* [x/xcgopointers](x/xcgopointers)

## Slice example

* [x/xcgoslice](x/xcgoslice)

## Wrapping C++

* [x/xcgocpp](x/xcgocpp)

Manually, or via SWIG.

## Performance experiment

A pseudo-random walk.

* [x/xcgoloop](x/xcgoloop)

## A thread mystery

* [x/xcgothreads](x/xcgothreads)

## GGML

* [x/xcgoggml](x/xcgoggml)

