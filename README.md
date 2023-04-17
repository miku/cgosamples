# C and Go

* pseudo-package "C"

> Cgo recognizes any use of a qualified identifier C.xxx and uses gcc to
find the definition of xxx.  If xxx is a type, cgo replaces C.xxx with a Go
translation.  C arithmetic types translate to precisely-sized Go arithmetic
types.  A C struct translates to a Go struct, field by field; unrepresentable
fields are replaced with opaque byte arrays.  A C union translates into a
struct containing the first union member and perhaps additional padding.  C
arrays become Go arrays.  C pointers become Go pointers.  C function pointers
become Go's uintptr.  C void pointers become Go's unsafe.Pointer.


