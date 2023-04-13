# CGO run

Tool `go tool cgo`, about 8K go code for parsing "C" package, translating calls
to C with Go "cgo" helpers, writing out files.

Basic C/GO:

```go
package main

/*

#include <stdio.h>
#include <stdlib.h>

void hello(char* s)
{
    printf("hello %s\n", s);
}

*/
import "C"
import "unsafe"

func main() {
        cs := C.CString("gopher")
        C.hello(cs)
        C.free(unsafe.Pointer(cs))
}
```

Generate a set of aux files:

```
$ go tool cgo main.go
$ tree _obj
.
├── main.go
├── Makefile
├── _obj
│   ├── _cgo_export.c
│   ├── _cgo_export.h
│   ├── _cgo_flags
│   ├── _cgo_gotypes.go
│   ├── _cgo_main.c
│   ├── _cgo_.o
│   ├── main.cgo1.go
│   └── main.cgo2.c
└── README.md

1 directory, 11 files
```

A C file is synthesized, compiled. Here, it contains our `hello`.

```
$ objdump -t _obj/_cgo_.o

_obj/_cgo_.o:     file format elf64-x86-64

SYMBOL TABLE:
0000000000000000 l    df *ABS*  0000000000000000 cgo-gcc-input-1876643937.c
0000000000000000 l    d  .text  0000000000000000 .text
0000000000000000 l     F .text  000000000000001f _GoStringLen
000000000000001f l     F .text  000000000000001f _GoStringPtr
0000000000000000 l    d  .rodata        0000000000000000 .rodata
0000000000000000 l    d  .debug_info    0000000000000000 .debug_info
0000000000000000 l    d  .debug_abbrev  0000000000000000 .debug_abbrev
0000000000000000 l    d  .debug_loc     0000000000000000 .debug_loc
0000000000000000 l    d  .debug_line    0000000000000000 .debug_line
0000000000000000 l    d  .debug_str     0000000000000000 .debug_str
000000000000003e g     F .text  000000000000002e hello
0000000000000000         *UND*  0000000000000000 printf
0000000000000000 g     O .bss   0000000000000008 __cgo__0
0000000000000008 g     O .bss   0000000000000008 __cgo__1
0000000000000010 g     O .bss   0000000000000008 __cgo__2
0000000000000018 g     O .bss   0000000000000008 __cgo__3
0000000000000020 g     O .bss   0000000000000008 __cgo__4
0000000000000028 g     O .bss   0000000000000008 __cgo__5
0000000000000000 g     O .data  0000000000000038 __cgodebug_ints
0000000000000040 g     O .data  0000000000000038 __cgodebug_floats
```

So, what is happening? [cgocall.go](https://go.dev/src/runtime/cgocall.go)

```go
// To call into the C function f from Go, the cgo-generated code calls
// runtime.cgocall(_cgo_Cfunc_f, frame), where _cgo_Cfunc_f is a
// gcc-compiled function written by cgo.
```

Where's `cgocall` in the generated code? In `_obj/_cgo_gotypes.go` generated Go code.

```
//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32
```

Sidenote: Magic of [compiler directives](https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives):

```
//go:linkname localname [importpath.name]
```

> Instead, the `//go:linkname` directive instructs the compiler to use
> "importpath.name" as the object file symbol name for the variable or
> function declared as "localname" in the source code. [...] Because this
> directive can subvert the type system and package modularity, it is only
> enabled in files that have imported "unsafe".

The Go code that call C:

```
//go:cgo_import_static _cgo_1b32228e3b82_Cfunc_hello
//go:linkname __cgofn__cgo_1b32228e3b82_Cfunc_hello _cgo_1b32228e3b82_Cfunc_hello
var __cgofn__cgo_1b32228e3b82_Cfunc_hello byte
var _cgo_1b32228e3b82_Cfunc_hello = unsafe.Pointer(&__cgofn__cgo_1b32228e3b82_Cfunc_hello)

//go:cgo_unsafe_args
func _Cfunc_hello(p0 *_Ctype_char) (r1 _Ctype_void) {
        _cgo_runtime_cgocall(_cgo_1b32228e3b82_Cfunc_hello, uintptr(unsafe.Pointer(&p0)))
        if _Cgo_always_false {
                _Cgo_use(p0)
        }
        return
}
```
