# Go to C++

> It is feasible and not hard, but the C++ functionality needs to be
mapped through a C layer.
https://groups.google.com/g/golang-nuts/c/jl5TLFQlVjI/m/JZ48mxoqDqAJ (Oct 14,
2014, 9:19:22 AM)

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

## Reinterpret Cast

> [When to use reinterpret_cast?](https://stackoverflow.com/questions/573294/when-to-use-reinterpret-cast), 590 up, viewed 450K times

There:

> For casting to and from `void*`, `static_cast` should be preferred.

## Symbols

Since C++ name-mangles symbols, we need a C bridge.

```
$ objdump -t librect.so

librect.so:     file format elf64-x86-64

SYMBOL TABLE:
0000000000000000 l    df *ABS*  0000000000000000              crtstuff.c
00000000000010e0 l     F .text  0000000000000000              deregister_tm_clones
0000000000001110 l     F .text  0000000000000000              register_tm_clones
0000000000001150 l     F .text  0000000000000000              __do_global_dtors_aux
0000000000004050 l     O .bss   0000000000000001              completed.0
0000000000003e18 l     O .fini_array    0000000000000000              __do_global_dtors_aux_fini_array_entry
0000000000001190 l     F .text  0000000000000000              frame_dummy
0000000000003e10 l     O .init_array    0000000000000000              __frame_dummy_init_array_entry
0000000000000000 l    df *ABS*  0000000000000000              rect.cpp
0000000000000000 l    df *ABS*  0000000000000000              rect.c
0000000000000000 l    df *ABS*  0000000000000000              crtstuff.c
00000000000021b0 l     O .eh_frame      0000000000000000              __FRAME_END__
0000000000000000 l    df *ABS*  0000000000000000
0000000000003e20 l     O .dynamic       0000000000000000              _DYNAMIC
0000000000004050 l     O .data  0000000000000000              __TMC_END__
0000000000004040 l     O .data  0000000000000000              __dso_handle
0000000000001000 l     F .init  0000000000000000              _init
0000000000004048 l     O .data  0000000000000008              DW.ref.__gxx_personality_v0
0000000000002000 l       .eh_frame_hdr  0000000000000000              __GNU_EH_FRAME_HDR
000000000000129c l     F .fini  0000000000000000              _fini
0000000000004000 l     O .got.plt       0000000000000000              _GLOBAL_OFFSET_TABLE_
0000000000000000         *UND*  0000000000000000              _Znwm
000000000000119a g     F .text  0000000000000028              _ZN9RectangleC1Eii
0000000000000000  w      *UND*  0000000000000000              __cxa_finalize
00000000000011c2 g     F .text  000000000000000f              _ZN9RectangleD2Ev
000000000000127c g     F .text  000000000000001e              rect_area
0000000000000000         *UND*  0000000000000000              __gxx_personality_v0
000000000000119a g     F .text  0000000000000028              _ZN9RectangleC2Eii
00000000000011fa g     F .text  000000000000001e              _ZN9Rectangle4areaEv
0000000000000000  w      *UND*  0000000000000000              _ITM_registerTMCloneTable
0000000000000000         *UND*  0000000000000000              _ZdlPvm
0000000000000000  w      *UND*  0000000000000000              _ITM_deregisterTMCloneTable
00000000000011d2 g     F .text  0000000000000028              _ZN9Rectangle10set_valuesEii
00000000000011c2 g     F .text  000000000000000f              _ZN9RectangleD1Ev
0000000000000000       F *UND*  0000000000000000              _Unwind_Resume@GCC_3.0
0000000000001218 g     F .text  0000000000000064              rect_new
0000000000000000  w      *UND*  0000000000000000              __gmon_start__

```
