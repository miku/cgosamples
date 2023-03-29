# Cgo basic

Can use a single file.

```
package main

/*

#include <stdio.h>
#include <stdlib.h>

void print_name(char* s)
{
    printf("%s\n", s);
}

*/
import "C"
import "fmt"

func main() {
        fmt.Println("vim-go")
}
```

* comment above pseudo-package "C"
* env: `CGO_ENABLED=1`, enabled by default

What happens with `CGO_ENABLED=0` with a C comment? It fails.

```
$ CGO_ENABLED=0 go run main.go
go: no Go source files
```

We can run `cgo` as a standalone tool:

```
$ go tool cgo main.go
$ tree _obj
_obj
├── _cgo_export.c
├── _cgo_export.h
├── _cgo_flags
├── _cgo_gotypes.go
├── _cgo_main.c
├── main.cgo1.go
└── main.cgo2.c

0 directories, 7 files
```

For example in `_cgo_gotypes.go` we find a reference to the C function, via `cgocall`, [code](https://github.com/golang/go/blob/f57f02fcd519a86683c47a444e0e24a086fea8e0/src/runtime/cgocall.go#L117-L200).

When running with `-debug-gcc` turned on, we get to keep the object file.

```
$ go tool cgo -debug-gcc main.go 2> debug.log
```

```
$ objdump -t _obj/_cgo_.o

_obj/_cgo_.o:     file format elf64-x86-64

SYMBOL TABLE:
0000000000000000 l    df *ABS*  0000000000000000 cgo-gcc-input-1015433406.c
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

Calls:

```
$ grep gcc debug.log
$ gcc -E -dM -xc -m64 - <<EOF
$ gcc -w -Wno-error -o_obj/_cgo_.o -gdwarf-2 -c -xc -m64 -fno-lto -O0 -fdiagnostics-color=never - <<EOF
$ gcc -w -Wno-error -o_obj/_cgo_.o -gdwarf-2 -c -xc -m64 -fno-lto - <<EOF
$ gcc -w -Wno-error -o_obj/_cgo_.o -gdwarf-2 -c -xc -m64 -fno-lto -O0 -fdiagnostics-color=never - <<EOF
$ gcc -w -Wno-error -o_obj/_cgo_.o -gdwarf-2 -c -xc -m64 -fno-lto - <<EOF
```


