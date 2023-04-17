# Go to C++

Not directly possible. Need to wrap manually, or via SWIG.

```
go run main.go
# command-line-arguments
./main.go:11:7: could not determine kind of name for C.Rectangle
cgo:
gcc errors for preamble:
In file included from ./main.go:4:
./rect.h:4:1: error: unknown type name 'class'
    4 | class Rectangle {
      | ^~~~~
./rect.h:4:17: error: expected '=', ',', ';', 'asm' or '__attribute__' before '{' token
    4 | class Rectangle {
      |                 ^

```

> This error? error: unknown type name ‘class’ You're probably compiling it as C rather than C++.

* [Unknown type name class](https://stackoverflow.com/a/16564736/89391)
