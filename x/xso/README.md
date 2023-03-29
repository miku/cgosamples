# C: object and shared object files

```
$ file main.o
main.o: ELF 64-bit LSB relocatable, x86-64, version 1 (SYSV), not stripped
```

There is a separate strip(1) command to discard symbols and other data from object files.

> https://stackoverflow.com/questions/3880924/how-to-view-symbols-in-object-files

```
$ objdump -t main.o

main.o:     file format elf64-x86-64

SYMBOL TABLE:
0000000000000000 l    df *ABS*  0000000000000000 main.c
0000000000000000 l    d  .text  0000000000000000 .text
0000000000000000 l    d  .rodata    0000000000000000 .rodata
0000000000000000 g     F .text  000000000000002e print_name
0000000000000000         *UND*  0000000000000000 printf
```

There are symbol and dynamic symbol tables. More headers.

```
$ objdump -x main.o

main.o:     file format elf64-x86-64
main.o
architecture: i386:x86-64, flags 0x00000011:
HAS_RELOC, HAS_SYMS
start address 0x0000000000000000

Sections:
Idx Name          Size      VMA               LMA               File off  Algn
  0 .text         0000002e  0000000000000000  0000000000000000  00000040  2**0
                  CONTENTS, ALLOC, LOAD, RELOC, READONLY, CODE
  1 .data         00000000  0000000000000000  0000000000000000  0000006e  2**0
                  CONTENTS, ALLOC, LOAD, DATA
  2 .bss          00000000  0000000000000000  0000000000000000  0000006e  2**0
                  ALLOC
  3 .rodata       0000000f  0000000000000000  0000000000000000  0000006e  2**0
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  4 .comment      0000002c  0000000000000000  0000000000000000  0000007d  2**0
                  CONTENTS, READONLY
  5 .note.GNU-stack 00000000  0000000000000000  0000000000000000  000000a9  2**0
                  CONTENTS, READONLY
  6 .note.gnu.property 00000020  0000000000000000  0000000000000000  000000b0  2**3
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  7 .eh_frame     00000038  0000000000000000  0000000000000000  000000d0  2**3
                  CONTENTS, ALLOC, LOAD, RELOC, READONLY, DATA
SYMBOL TABLE:
0000000000000000 l    df *ABS*  0000000000000000 main.c
0000000000000000 l    d  .text  0000000000000000 .text
0000000000000000 l    d  .rodata        0000000000000000 .rodata
0000000000000000 g     F .text  000000000000002e print_name
0000000000000000         *UND*  0000000000000000 printf


RELOCATION RECORDS FOR [.text]:
OFFSET           TYPE              VALUE
000000000000001a R_X86_64_PC32     .rodata-0x0000000000000004
0000000000000027 R_X86_64_PLT32    printf-0x0000000000000004


RELOCATION RECORDS FOR [.eh_frame]:
OFFSET           TYPE              VALUE
0000000000000020 R_X86_64_PC32     .text
```

Sections:

* [data](https://en.wikipedia.org/wiki/Data_segment), does not change at runtime
* [.bss](https://en.wikipedia.org/wiki/.bss), block starting symbol

Part of ELF.

* [https://refspecs.linuxfoundation.org/elf/elf.pdf](https://refspecs.linuxfoundation.org/elf/elf.pdf)


