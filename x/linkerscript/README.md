## Object files

A C hello world binary has 26 sections, whereas the Go output has 18.

```
Sections:
Idx Name          Size      VMA               LMA               File off  Algn
  0 .interp       0000001c  0000000000000318  0000000000000318  00000318  2**0
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  1 .note.gnu.property 00000030  0000000000000338  0000000000000338  00000338  2**3
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  2 .note.gnu.build-id 00000024  0000000000000368  0000000000000368  00000368  2**2
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  3 .note.ABI-tag 00000020  000000000000038c  000000000000038c  0000038c  2**2
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  4 .gnu.hash     00000024  00000000000003b0  00000000000003b0  000003b0  2**3
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  5 .dynsym       000000a8  00000000000003d8  00000000000003d8  000003d8  2**3
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  6 .dynstr       0000008d  0000000000000480  0000000000000480  00000480  2**0
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  7 .gnu.version  0000000e  000000000000050e  000000000000050e  0000050e  2**1
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  8 .gnu.version_r 00000030  0000000000000520  0000000000000520  00000520  2**3
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  9 .rela.dyn     000000c0  0000000000000550  0000000000000550  00000550  2**3
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
 10 .rela.plt     00000018  0000000000000610  0000000000000610  00000610  2**3
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
 11 .init         0000001b  0000000000001000  0000000000001000  00001000  2**2
                  CONTENTS, ALLOC, LOAD, READONLY, CODE
 12 .plt          00000020  0000000000001020  0000000000001020  00001020  2**4
                  CONTENTS, ALLOC, LOAD, READONLY, CODE
 13 .plt.got      00000010  0000000000001040  0000000000001040  00001040  2**4
                  CONTENTS, ALLOC, LOAD, READONLY, CODE
 14 .plt.sec      00000010  0000000000001050  0000000000001050  00001050  2**4
                  CONTENTS, ALLOC, LOAD, READONLY, CODE
 15 .text         00000112  0000000000001060  0000000000001060  00001060  2**4
                  CONTENTS, ALLOC, LOAD, READONLY, CODE
 16 .fini         0000000d  0000000000001174  0000000000001174  00001174  2**2
                  CONTENTS, ALLOC, LOAD, READONLY, CODE
 17 .rodata       00000012  0000000000002000  0000000000002000  00002000  2**2
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
 18 .eh_frame_hdr 00000034  0000000000002014  0000000000002014  00002014  2**2
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
 19 .eh_frame     000000ac  0000000000002048  0000000000002048  00002048  2**3
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
 20 .init_array   00000008  0000000000003db8  0000000000003db8  00002db8  2**3
                  CONTENTS, ALLOC, LOAD, DATA
 21 .fini_array   00000008  0000000000003dc0  0000000000003dc0  00002dc0  2**3
                  CONTENTS, ALLOC, LOAD, DATA
 22 .dynamic      000001f0  0000000000003dc8  0000000000003dc8  00002dc8  2**3
                  CONTENTS, ALLOC, LOAD, DATA
 23 .got          00000048  0000000000003fb8  0000000000003fb8  00002fb8  2**3
                  CONTENTS, ALLOC, LOAD, DATA
 24 .data         00000010  0000000000004000  0000000000004000  00003000  2**3
                  CONTENTS, ALLOC, LOAD, DATA
 25 .bss          00000008  0000000000004010  0000000000004010  00003010  2**0
                  ALLOC
 26 .comment      0000002b  0000000000000000  0000000000000000  00003010  2**0
                  CONTENTS, READONLY
```

Go hello world.

```
Sections:
Idx Name          Size      VMA               LMA               File off  Algn
  0 .text         000800c7  0000000000401000  0000000000401000  00001000  2**5
                  CONTENTS, ALLOC, LOAD, READONLY, CODE
  1 .rodata       00036f45  0000000000482000  0000000000482000  00082000  2**5
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  2 .typelink     00000508  00000000004b90e0  00000000004b90e0  000b90e0  2**5
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  3 .itablink     00000058  00000000004b9600  00000000004b9600  000b9600  2**5
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  4 .gosymtab     00000000  00000000004b9658  00000000004b9658  000b9658  2**0
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  5 .gopclntab    0005a1b0  00000000004b9660  00000000004b9660  000b9660  2**5
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
  6 .go.buildinfo 00000130  0000000000514000  0000000000514000  00114000  2**4
                  CONTENTS, ALLOC, LOAD, DATA
  7 .noptrdata    00010600  0000000000514140  0000000000514140  00114140  2**5
                  CONTENTS, ALLOC, LOAD, DATA
  8 .data         00007790  0000000000524740  0000000000524740  00124740  2**5
                  CONTENTS, ALLOC, LOAD, DATA
  9 .bss          0002df80  000000000052bee0  000000000052bee0  0012bee0  2**5
                  ALLOC
 10 .noptrbss     000039d0  0000000000559e60  0000000000559e60  00159e60  2**5
                  ALLOC
 11 .debug_abbrev 00000210  0000000000000000  0000000000000000  0012c000  2**0
                  CONTENTS, READONLY, DEBUGGING, OCTETS
 12 .debug_line   00032caf  0000000000000000  0000000000000000  0012c133  2**0
                  CONTENTS, READONLY, DEBUGGING, OCTETS
 13 .debug_frame  0000f69c  0000000000000000  0000000000000000  00148f6d  2**0
                  CONTENTS, READONLY, DEBUGGING, OCTETS
 14 .debug_gdb_scripts 0000002a  0000000000000000  0000000000000000  0014e777  2**0
                  CONTENTS, READONLY, DEBUGGING, OCTETS
 15 .debug_info   0008a8c3  0000000000000000  0000000000000000  0014e7a1  2**0
                  CONTENTS, READONLY, DEBUGGING, OCTETS
 16 .debug_loc    0009cf3f  0000000000000000  0000000000000000  0018673a  2**0
                  CONTENTS, READONLY, DEBUGGING, OCTETS
 17 .debug_ranges 00033300  0000000000000000  0000000000000000  001a168d  2**0
                  CONTENTS, READONLY, DEBUGGING, OCTETS
 18 .note.go.buildid 00000064  0000000000400f9c  0000000000400f9c  00000f9c  2**2
                  CONTENTS, ALLOC, LOAD, READONLY, DATA
```

Go does not have:

* `eh_frame...`
* `dynamic`
* `gnu.version...`
* ...

Using readelf.

```
ELF Header:
  Magic:   7f 45 4c 46 02 01 01 00 00 00 00 00 00 00 00 00
  Class:                             ELF64
  Data:                              2's complement, little endian
  Version:                           1 (current)
  OS/ABI:                            UNIX - System V
  ABI Version:                       0
  Type:                              EXEC (Executable file)
  Machine:                           Advanced Micro Devices X86-64
  Version:                           0x1
  Entry point address:               0x45ed80
  Start of program headers:          64 (bytes into file)
  Start of section headers:          456 (bytes into file)
  Flags:                             0x0
  Size of this header:               64 (bytes)
  Size of program headers:           56 (bytes)
  Number of program headers:         7
  Size of section headers:           64 (bytes)
  Number of section headers:         23
  Section header string table index: 3

Section Headers:
  [Nr] Name              Type             Address           Offset
       Size              EntSize          Flags  Link  Info  Align
  [ 0]                   NULL             0000000000000000  00000000
       0000000000000000  0000000000000000           0     0     0
  [ 1] .text             PROGBITS         0000000000401000  00001000
       00000000000800c7  0000000000000000  AX       0     0     32
  [ 2] .rodata           PROGBITS         0000000000482000  00082000
       0000000000036f45  0000000000000000   A       0     0     32
  [ 3] .shstrtab         STRTAB           0000000000000000  000b8f60
       000000000000016d  0000000000000000           0     0     1
  [ 4] .typelink         PROGBITS         00000000004b90e0  000b90e0
       0000000000000508  0000000000000000   A       0     0     32
  [ 5] .itablink         PROGBITS         00000000004b9600  000b9600
       0000000000000058  0000000000000000   A       0     0     32
  [ 6] .gosymtab         PROGBITS         00000000004b9658  000b9658
       0000000000000000  0000000000000000   A       0     0     1
  [ 7] .gopclntab        PROGBITS         00000000004b9660  000b9660
       000000000005a1b0  0000000000000000   A       0     0     32
  [ 8] .go.buildinfo     PROGBITS         0000000000514000  00114000
       0000000000000130  0000000000000000  WA       0     0     16
  [ 9] .noptrdata        PROGBITS         0000000000514140  00114140
       0000000000010600  0000000000000000  WA       0     0     32
  [10] .data             PROGBITS         0000000000524740  00124740
       0000000000007790  0000000000000000  WA       0     0     32
  [11] .bss              NOBITS           000000000052bee0  0012bee0
       000000000002df80  0000000000000000  WA       0     0     32
  [12] .noptrbss         NOBITS           0000000000559e60  00159e60
       00000000000039d0  0000000000000000  WA       0     0     32
  [13] .debug_abbrev     PROGBITS         0000000000000000  0012c000
       0000000000000133  0000000000000000   C       0     0     1
  [14] .debug_line       PROGBITS         0000000000000000  0012c133
       000000000001ce3a  0000000000000000   C       0     0     1
  [15] .debug_frame      PROGBITS         0000000000000000  00148f6d
       000000000000580a  0000000000000000   C       0     0     1
  [16] .debug_gdb_s[...] PROGBITS         0000000000000000  0014e777
       000000000000002a  0000000000000000           0     0     1
  [17] .debug_info       PROGBITS         0000000000000000  0014e7a1
       0000000000037f99  0000000000000000   C       0     0     1
  [18] .debug_loc        PROGBITS         0000000000000000  0018673a
       000000000001af53  0000000000000000   C       0     0     1
  [19] .debug_ranges     PROGBITS         0000000000000000  001a168d
       0000000000009473  0000000000000000   C       0     0     1
  [20] .note.go.buildid  NOTE             0000000000400f9c  00000f9c
       0000000000000064  0000000000000000   A       0     0     4
  [21] .symtab           SYMTAB           0000000000000000  001aab00
       000000000000c8d0  0000000000000018          22   104     8
  [22] .strtab           STRTAB           0000000000000000  001b73d0
       000000000000b6ba  0000000000000000           0     0     1
Key to Flags:
  W (write), A (alloc), X (execute), M (merge), S (strings), I (info),
  L (link order), O (extra OS processing required), G (group), T (TLS),
  C (compressed), x (unknown), o (OS specific), E (exclude),
  D (mbind), l (large), p (processor specific)

There are no section groups in this file.

Program Headers:
  Type           Offset             VirtAddr           PhysAddr
                 FileSiz            MemSiz              Flags  Align
  PHDR           0x0000000000000040 0x0000000000400040 0x0000000000400040
                 0x0000000000000188 0x0000000000000188  R      0x1000
  NOTE           0x0000000000000f9c 0x0000000000400f9c 0x0000000000400f9c
                 0x0000000000000064 0x0000000000000064  R      0x4
  LOAD           0x0000000000000000 0x0000000000400000 0x0000000000400000
                 0x00000000000810c7 0x00000000000810c7  R E    0x1000
  LOAD           0x0000000000082000 0x0000000000482000 0x0000000000482000
                 0x0000000000091810 0x0000000000091810  R      0x1000
  LOAD           0x0000000000114000 0x0000000000514000 0x0000000000514000
                 0x0000000000017ee0 0x0000000000049830  RW     0x1000
  GNU_STACK      0x0000000000000000 0x0000000000000000 0x0000000000000000
                 0x0000000000000000 0x0000000000000000  RW     0x8
  LOOS+0x5041580 0x0000000000000000 0x0000000000000000 0x0000000000000000
                 0x0000000000000000 0x0000000000000000         0x8

 Section to Segment mapping:
  Segment Sections...
   00
   01     .note.go.buildid
   02     .text .note.go.buildid
   03     .rodata .typelink .itablink .gosymtab .gopclntab
   04     .go.buildinfo .noptrdata .data .bss .noptrbss
   05
   06

There is no dynamic section in this file.

There are no relocations in this file.
No processor specific unwind information to decode

Symbol table '.symtab' contains 2142 entries:
   Num:    Value          Size Type    Bind   Vis      Ndx Name
     0: 0000000000000000     0 NOTYPE  LOCAL  DEFAULT  UND
     1: 0000000000000000     0 FILE    LOCAL  DEFAULT  ABS go.go
     2: 0000000000401000     0 FUNC    LOCAL  DEFAULT    1 runtime.text
     3: 00000000004021a0   557 FUNC    LOCAL  DEFAULT    1 cmpbody
     4: 0000000000402400   318 FUNC    LOCAL  DEFAULT    1 memeqbody
     5: 0000000000402580   279 FUNC    LOCAL  DEFAULT    1 indexbytebody
     6: 000000000045a760    64 FUNC    LOCAL  DEFAULT    1 gogo
     7: 000000000045a7a0    53 FUNC    LOCAL  DEFAULT    1 callRet
     8: 000000000045a7e0    47 FUNC    LOCAL  DEFAULT    1 gosave_systemsta[...]
     9: 000000000045a820    13 FUNC    LOCAL  DEFAULT    1 setg_gcc
    10: 000000000045a840  1370 FUNC    LOCAL  DEFAULT    1 aeshashbody
    11: 000000000045ada0    75 FUNC    LOCAL  DEFAULT    1 debugCall32
    12: 000000000045ae00    75 FUNC    LOCAL  DEFAULT    1 debugCall64
    13: 000000000045ae60   111 FUNC    LOCAL  DEFAULT    1 debugCall128
    14: 000000000045aee0   114 FUNC    LOCAL  DEFAULT    1 debugCall256
    15: 000000000045af60   114 FUNC    LOCAL  DEFAULT    1 debugCall512
    16: 000000000045afe0   114 FUNC    LOCAL  DEFAULT    1 debugCall1024
    17: 000000000045b060   114 FUNC    LOCAL  DEFAULT    1 debugCall2048
    18: 000000000045b0e0   118 FUNC    LOCAL  DEFAULT    1 debugCall4096
    19: 000000000045b160   118 FUNC    LOCAL  DEFAULT    1 debugCall8192
    20: 000000000045b1e0   118 FUNC    LOCAL  DEFAULT    1 debugCall16384
    21: 000000000045b260   118 FUNC    LOCAL  DEFAULT    1 debugCall32768
    22: 000000000045b2e0   118 FUNC    LOCAL  DEFAULT    1 debugCall65536
    23: 000000000045eda0   108 FUNC    LOCAL  DEFAULT    1 runtime.sigprofN[...]
    24: 00000000004810c7     0 FUNC    LOCAL  DEFAULT    1 runtime.etext
    25: 00000000005146d0    16 OBJECT  LOCAL  DEFAULT    9 runtime..gobytes.2
    26: 0000000000515cc0    51 OBJECT  LOCAL  DEFAULT    9 runtime..gobytes.3
    27: 00000000005146b8    13 OBJECT  LOCAL  DEFAULT    9 runtime..gobytes.4
    28: 00000000004b7c80    84 OBJECT  LOCAL  DEFAULT    2 bad_cpu_msg
    29: 00000000004b8220   256 OBJECT  LOCAL  DEFAULT    2 masks
    30: 00000000004b8320   256 OBJECT  LOCAL  DEFAULT    2 shifts
    31: 00000000004b75c0    20 OBJECT  LOCAL  DEFAULT    2 debugCallFrameTo[...]
    32: 00000000004b8040   216 OBJECT  LOCAL  DEFAULT    2 fmt.(*pp).printV[...]
    33: 00000000004b7110     8 OBJECT  LOCAL  DEFAULT    2 $f64.8000000000000000
    34: 00000000004b7e80   208 OBJECT  LOCAL  DEFAULT    2 internal/fmtsort[...]
    35: 00000000004b7100     8 OBJECT  LOCAL  DEFAULT    2 $f64.43e0000000000000
    36: 00000000004b7f60   208 OBJECT  LOCAL  DEFAULT    2 reflect.(*abiSeq[...]
    37: 00000000004b7740    72 OBJECT  LOCAL  DEFAULT    2 reflect.(*rtype)[...]
    38: 00000000004b7080     8 OBJECT  LOCAL  DEFAULT    2 $f64.3fd0000000000000
    39: 00000000004b7090     8 OBJECT  LOCAL  DEFAULT    2 $f64.3fe0000000000000
    40: 00000000004b70a0     8 OBJECT  LOCAL  DEFAULT    2 $f64.3ff0000000000000
    41: 00000000004b70b8     8 OBJECT  LOCAL  DEFAULT    2 $f64.4010000000000000
    42: 00000000004b70c0     8 OBJECT  LOCAL  DEFAULT    2 $f64.4014000000000000
    43: 00000000004b7108     8 OBJECT  LOCAL  DEFAULT    2 $f64.7ff0000000000000
    44: 00000000004b7068     8 OBJECT  LOCAL  DEFAULT    2 $f64.3eb0000000000000
    45: 00000000004b7070     8 OBJECT  LOCAL  DEFAULT    2 $f64.3f50624dd2f1a9fc
    46: 00000000004b7078     8 OBJECT  LOCAL  DEFAULT    2 $f64.3f847ae147ae147b
    47: 00000000004b7088     8 OBJECT  LOCAL  DEFAULT    2 $f64.3fd3333333333333
    48: 00000000004b7098     8 OBJECT  LOCAL  DEFAULT    2 $f64.3fe8000000000000
    49: 00000000004b70a8     8 OBJECT  LOCAL  DEFAULT    2 $f64.3ff199999999999a
    50: 00000000004b70b0     8 OBJECT  LOCAL  DEFAULT    2 $f64.3ff3333333333333
    51: 00000000004b70c8     8 OBJECT  LOCAL  DEFAULT    2 $f64.4024000000000000
    52: 00000000004b70d0     8 OBJECT  LOCAL  DEFAULT    2 $f64.403a000000000000
    53: 00000000004b70d8     8 OBJECT  LOCAL  DEFAULT    2 $f64.4057c00000000000
    54: 00000000004b70e0     8 OBJECT  LOCAL  DEFAULT    2 $f64.4059000000000000
    55: 00000000004b70e8     8 OBJECT  LOCAL  DEFAULT    2 $f64.40c3880000000000
    56: 00000000004b70f0     8 OBJECT  LOCAL  DEFAULT    2 $f64.40f0000000000000
    57: 00000000004b70f8     8 OBJECT  LOCAL  DEFAULT    2 $f64.412e848000000000
    58: 00000000004b7118     8 OBJECT  LOCAL  DEFAULT    2 $f64.bfd3333333333333
    59: 00000000004b7120     8 OBJECT  LOCAL  DEFAULT    2 $f64.bfe62e42fefa39ef
    60: 00000000004b7ce0   104 OBJECT  LOCAL  DEFAULT    2 runtime.typehash[...]
    61: 00000000004b7dc0   192 OBJECT  LOCAL  DEFAULT    2 runtime.printany[...]
    62: 00000000004b7800    72 OBJECT  LOCAL  DEFAULT    2 runtime.(*itab).[...]
    63: 00000000004b78c0    72 OBJECT  LOCAL  DEFAULT    2 runtime.SetFinal[...]
    64: 00000000004b7860    72 OBJECT  LOCAL  DEFAULT    2 runtime.SetFinal[...]
    65: 00000000004b7aa0    80 OBJECT  LOCAL  DEFAULT    2 runtime.deltimer[...]
    66: 00000000004b7b00    80 OBJECT  LOCAL  DEFAULT    2 runtime.modtimer[...]
    67: 00000000004b7b60    80 OBJECT  LOCAL  DEFAULT    2 runtime.moveTime[...]
    68: 00000000004b79e0    80 OBJECT  LOCAL  DEFAULT    2 runtime.adjustti[...]
    69: 00000000004b7bc0    80 OBJECT  LOCAL  DEFAULT    2 runtime.runtimer[...]
    70: 00000000004b7a40    80 OBJECT  LOCAL  DEFAULT    2 runtime.clearDel[...]
    71: 00000000004b77a0    72 OBJECT  LOCAL  DEFAULT    2 runtime.(*_type)[...]
    72: 00000000004b7920    72 OBJECT  LOCAL  DEFAULT    2 runtime.typesEqu[...]
    73: 00000000004b7980    72 OBJECT  LOCAL  DEFAULT    2 runtime.typesEqu[...]
    74: 00000000004b7c20    80 OBJECT  LOCAL  DEFAULT    2 runtime.typesEqu[...]
    75: 00000000004b90e0  1288 OBJECT  LOCAL  DEFAULT    4 runtime.typelink
    76: 00000000004b9600    88 OBJECT  LOCAL  DEFAULT    5 runtime.itablink
    77: 00000000004b9660     0 OBJECT  LOCAL  DEFAULT    7 runtime.pclntab
    78: 00000000004b8540  2565 OBJECT  LOCAL  DEFAULT    2 runtime.findfunctab
    79: 0000000000482000     0 OBJECT  LOCAL  DEFAULT    2 runtime.rodata
    80: 00000000004b8f45     0 OBJECT  LOCAL  DEFAULT    2 runtime.erodata
    81: 0000000000482000     0 OBJECT  LOCAL  DEFAULT    2 runtime.types
    82: 00000000004b8f45     0 OBJECT  LOCAL  DEFAULT    2 runtime.etypes
    83: 0000000000514140     0 OBJECT  LOCAL  DEFAULT    9 runtime.noptrdata
    84: 0000000000524740     0 OBJECT  LOCAL  DEFAULT    9 runtime.enoptrdata
    85: 0000000000524740     0 OBJECT  LOCAL  DEFAULT   10 runtime.data
    86: 000000000052bed0     0 OBJECT  LOCAL  DEFAULT   10 runtime.edata
    87: 000000000052bee0     0 OBJECT  LOCAL  DEFAULT   11 runtime.bss
    88: 0000000000559e60     0 OBJECT  LOCAL  DEFAULT   11 runtime.ebss
    89: 0000000000559e60     0 OBJECT  LOCAL  DEFAULT   12 runtime.noptrbss
    90: 000000000055d830     0 OBJECT  LOCAL  DEFAULT   12 runtime.enoptrbss
    91: 000000000055d830     0 OBJECT  LOCAL  DEFAULT   12 runtime.covctrs
    92: 000000000055d830     0 OBJECT  LOCAL  DEFAULT   12 runtime.ecovctrs
    93: 000000000055d830     0 OBJECT  LOCAL  DEFAULT   12 runtime.end
    94: 0000000000513810     0 OBJECT  LOCAL  DEFAULT    7 runtime.epclntab
    95: 00000000004b9658     0 OBJECT  LOCAL  DEFAULT    6 runtime.esymtab
    96: 00000000004b6e65   514 OBJECT  LOCAL  DEFAULT    2 runtime.gcdata
    97: 00000000004b7067     0 OBJECT  LOCAL  DEFAULT    2 runtime.egcdata
    98: 00000000004b66f8  1901 OBJECT  LOCAL  DEFAULT    2 runtime.gcbss
    99: 00000000004b6e65     0 OBJECT  LOCAL  DEFAULT    2 runtime.egcbss
   100: 0000000000499418     0 OBJECT  LOCAL  DEFAULT    2 go:string.*
   101: 00000000004a1980     0 OBJECT  LOCAL  DEFAULT    2 go:func.*
   102: 00000000004b6420     0 OBJECT  LOCAL  DEFAULT    2 runtime.gcbits.*
   103: 00000000004b9658     0 OBJECT  LOCAL  DEFAULT    6 runtime.symtab
   104: 0000000000401000    89 FUNC    GLOBAL DEFAULT    1 internal/cpu.Ini[...]
   105: 0000000000401060  1367 FUNC    GLOBAL DEFAULT    1 internal/cpu.pro[...]
   106: 00000000004015c0  2155 FUNC    GLOBAL DEFAULT    1 internal/cpu.doinit
   107: 0000000000401e40    27 FUNC    GLOBAL DEFAULT    1 internal/cpu.cpu[...]
   108: 0000000000401e60    17 FUNC    GLOBAL DEFAULT    1 internal/cpu.xge[...]
   109: 0000000000401e80     9 FUNC    GLOBAL DEFAULT    1 internal/cpu.get[...]
   110: 0000000000401ea0   122 FUNC    GLOBAL DEFAULT    1 type:.eq.interna[...]
   111: 0000000000401f20   230 FUNC    GLOBAL DEFAULT    1 type:.eq.[6]inte[...]
   112: 0000000000402020    10 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
   113: 0000000000402040    10 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
   114: 0000000000402060   110 FUNC    GLOBAL DEFAULT    1 runtime/internal[...]
   115: 00000000004020e0     6 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
   116: 0000000000402100    67 FUNC    GLOBAL DEFAULT    1 type:.eq.interna[...]
   117: 0000000000402160    34 FUNC    GLOBAL DEFAULT    1 internal/bytealg[...]
   118: 00000000004023e0    14 FUNC    GLOBAL DEFAULT    1 runtime.cmpstring
   119: 0000000000402540    27 FUNC    GLOBAL DEFAULT    1 runtime.memequal
   120: 0000000000402560    28 FUNC    GLOBAL DEFAULT    1 runtime.memequal[...]
   121: 00000000004026a0    24 FUNC    GLOBAL DEFAULT    1 internal/bytealg[...]
   122: 00000000004026c0   152 FUNC    GLOBAL DEFAULT    1 runtime/internal[...]
   123: 0000000000402760    29 FUNC    GLOBAL DEFAULT    1 syscall.RawSyscall6
   124: 0000000000402780    54 FUNC    GLOBAL DEFAULT    1 runtime/internal[...]
   125: 00000000004027c0    67 FUNC    GLOBAL DEFAULT    1 runtime.memhash128
   126: 0000000000402820    74 FUNC    GLOBAL DEFAULT    1 runtime.strhashF[...]
   127: 0000000000402880   229 FUNC    GLOBAL DEFAULT    1 runtime.f32hash
   128: 0000000000402980   229 FUNC    GLOBAL DEFAULT    1 runtime.f64hash
   129: 0000000000402a80    94 FUNC    GLOBAL DEFAULT    1 runtime.c64hash
   130: 0000000000402ae0    94 FUNC    GLOBAL DEFAULT    1 runtime.c128hash
   131: 0000000000402b40   261 FUNC    GLOBAL DEFAULT    1 runtime.interhash
   132: 0000000000402c60   261 FUNC    GLOBAL DEFAULT    1 runtime.nilinterhash
   133: 0000000000402d80   654 FUNC    GLOBAL DEFAULT    1 runtime.typehash
   134: 0000000000403020     6 FUNC    GLOBAL DEFAULT    1 runtime.memequal0
   135: 0000000000403040     9 FUNC    GLOBAL DEFAULT    1 runtime.memequal8
   136: 0000000000403060    10 FUNC    GLOBAL DEFAULT    1 runtime.memequal16
   137: 0000000000403080     8 FUNC    GLOBAL DEFAULT    1 runtime.memequal32
   138: 00000000004030a0    10 FUNC    GLOBAL DEFAULT    1 runtime.memequal64
   139: 00000000004030c0    26 FUNC    GLOBAL DEFAULT    1 runtime.memequal128
   140: 00000000004030e0    20 FUNC    GLOBAL DEFAULT    1 runtime.f32equal
   141: 0000000000403100    21 FUNC    GLOBAL DEFAULT    1 runtime.f64equal
   142: 0000000000403120    43 FUNC    GLOBAL DEFAULT    1 runtime.c64equal
   143: 0000000000403160    45 FUNC    GLOBAL DEFAULT    1 runtime.c128equal
   144: 00000000004031a0    87 FUNC    GLOBAL DEFAULT    1 runtime.strequal
   145: 0000000000403200    89 FUNC    GLOBAL DEFAULT    1 runtime.interequal
   146: 0000000000403260    89 FUNC    GLOBAL DEFAULT    1 runtime.nilinterequal
   147: 00000000004032c0   197 FUNC    GLOBAL DEFAULT    1 runtime.efaceeq
   148: 00000000004033a0   197 FUNC    GLOBAL DEFAULT    1 runtime.ifaceeq
   149: 0000000000403480   165 FUNC    GLOBAL DEFAULT    1 runtime.alginit
   150: 0000000000403540    92 FUNC    GLOBAL DEFAULT    1 runtime.init.0
   151: 00000000004035a0   403 FUNC    GLOBAL DEFAULT    1 runtime.(*mspan)[...]
   152: 0000000000403740    91 FUNC    GLOBAL DEFAULT    1 runtime.(*mspan)[...]
   153: 00000000004037a0    92 FUNC    GLOBAL DEFAULT    1 runtime.atomicwb
   154: 0000000000403800    61 FUNC    GLOBAL DEFAULT    1 runtime.atomicstorep
   155: 0000000000403840   235 FUNC    GLOBAL DEFAULT    1 runtime.mmap
   156: 0000000000403940   124 FUNC    GLOBAL DEFAULT    1 runtime.mmap.func1
   157: 00000000004039c0   180 FUNC    GLOBAL DEFAULT    1 runtime.munmap
   158: 0000000000403a80    72 FUNC    GLOBAL DEFAULT    1 runtime.munmap.func1
   159: 0000000000403ae0   342 FUNC    GLOBAL DEFAULT    1 runtime.sigaction
   160: 0000000000403c40   104 FUNC    GLOBAL DEFAULT    1 runtime.sigactio[...]
   161: 0000000000403cc0   230 FUNC    GLOBAL DEFAULT    1 runtime.cgocall
   162: 0000000000403dc0   193 FUNC    GLOBAL DEFAULT    1 runtime.cgoIsGoP[...]
   163: 0000000000403ea0   255 FUNC    GLOBAL DEFAULT    1 runtime.cgoCheck[...]
   164: 0000000000403fa0   149 FUNC    GLOBAL DEFAULT    1 runtime.cgoCheck[...]
   165: 0000000000404040   139 FUNC    GLOBAL DEFAULT    1 runtime.cgoCheck[...]
   166: 00000000004040e0   180 FUNC    GLOBAL DEFAULT    1 runtime.cgoCheck[...]
   167: 00000000004041a0   544 FUNC    GLOBAL DEFAULT    1 runtime.cgoCheck[...]
   168: 00000000004043c0    58 FUNC    GLOBAL DEFAULT    1 runtime.cgoCheck[...]
   169: 0000000000404400   198 FUNC    GLOBAL DEFAULT    1 runtime.cgoCheckBits
   170: 00000000004044e0   517 FUNC    GLOBAL DEFAULT    1 runtime.cgoCheck[...]
   171: 0000000000404700   469 FUNC    GLOBAL DEFAULT    1 runtime.makechan
   172: 00000000004048e0    39 FUNC    GLOBAL DEFAULT    1 runtime.chansend1
   173: 0000000000404920  1610 FUNC    GLOBAL DEFAULT    1 runtime.chansend
   174: 0000000000404f80    54 FUNC    GLOBAL DEFAULT    1 runtime.chansend[...]
   175: 0000000000404fc0   311 FUNC    GLOBAL DEFAULT    1 runtime.send
   176: 0000000000405100   118 FUNC    GLOBAL DEFAULT    1 runtime.sendDirect
   177: 0000000000405180   123 FUNC    GLOBAL DEFAULT    1 runtime.recvDirect
   178: 0000000000405200  1061 FUNC    GLOBAL DEFAULT    1 runtime.closechan
   179: 0000000000405640    34 FUNC    GLOBAL DEFAULT    1 runtime.chanrecv1
   180: 0000000000405680  1743 FUNC    GLOBAL DEFAULT    1 runtime.chanrecv
   181: 0000000000405d60    54 FUNC    GLOBAL DEFAULT    1 runtime.chanrecv[...]
   182: 0000000000405da0   464 FUNC    GLOBAL DEFAULT    1 runtime.recv
   183: 0000000000405f80    92 FUNC    GLOBAL DEFAULT    1 runtime.chanpark[...]
   184: 0000000000405fe0    84 FUNC    GLOBAL DEFAULT    1 runtime.init.1
   185: 0000000000406040   357 FUNC    GLOBAL DEFAULT    1 runtime.(*cpuPro[...]
   186: 00000000004061c0   247 FUNC    GLOBAL DEFAULT    1 runtime.(*cpuPro[...]
   187: 00000000004062c0   538 FUNC    GLOBAL DEFAULT    1 runtime.(*cpuPro[...]
   188: 00000000004064e0   184 FUNC    GLOBAL DEFAULT    1 runtime.GOMAXPROCS
   189: 00000000004065a0   210 FUNC    GLOBAL DEFAULT    1 runtime.debugCal[...]
   190: 0000000000406680   798 FUNC    GLOBAL DEFAULT    1 runtime.debugCal[...]
   191: 00000000004069a0   218 FUNC    GLOBAL DEFAULT    1 runtime.debugCallWrap
   192: 0000000000406a80   336 FUNC    GLOBAL DEFAULT    1 runtime.debugCal[...]
   193: 0000000000406be0   125 FUNC    GLOBAL DEFAULT    1 runtime.debugCal[...]
   194: 0000000000406c60   188 FUNC    GLOBAL DEFAULT    1 runtime.debugCal[...]
   195: 0000000000406d20    83 FUNC    GLOBAL DEFAULT    1 runtime.debugCal[...]
   196: 0000000000406d80   314 FUNC    GLOBAL DEFAULT    1 runtime.gogetenv
   197: 0000000000406ec0   937 FUNC    GLOBAL DEFAULT    1 runtime.(*TypeAs[...]
   198: 0000000000407280    87 FUNC    GLOBAL DEFAULT    1 runtime.errorStr[...]
   199: 00000000004072e0    98 FUNC    GLOBAL DEFAULT    1 runtime.errorAdd[...]
   200: 0000000000407360     6 FUNC    GLOBAL DEFAULT    1 runtime.plainErr[...]
   201: 0000000000407380  1585 FUNC    GLOBAL DEFAULT    1 runtime.boundsEr[...]
   202: 00000000004079c0  1404 FUNC    GLOBAL DEFAULT    1 runtime.printany
   203: 0000000000407f40  1852 FUNC    GLOBAL DEFAULT    1 runtime.printany[...]
   204: 0000000000408680   885 FUNC    GLOBAL DEFAULT    1 runtime.panicwrap
   205: 0000000000408a00   314 FUNC    GLOBAL DEFAULT    1 runtime.runExitHooks
   206: 0000000000408b40   501 FUNC    GLOBAL DEFAULT    1 runtime.memhashF[...]
   207: 0000000000408d40    69 FUNC    GLOBAL DEFAULT    1 runtime.memhash3[...]
   208: 0000000000408da0    70 FUNC    GLOBAL DEFAULT    1 runtime.memhash6[...]
   209: 0000000000408e00   198 FUNC    GLOBAL DEFAULT    1 runtime.(*timeHi[...]
   210: 0000000000408ee0   750 FUNC    GLOBAL DEFAULT    1 runtime.getitab
   211: 00000000004091e0    70 FUNC    GLOBAL DEFAULT    1 runtime.(*itabTa[...]
   212: 0000000000409240   314 FUNC    GLOBAL DEFAULT    1 runtime.itabAdd
   213: 0000000000409380    78 FUNC    GLOBAL DEFAULT    1 runtime.(*itabTa[...]
   214: 00000000004093e0  1029 FUNC    GLOBAL DEFAULT    1 runtime.(*itab).init
   215: 0000000000409800   209 FUNC    GLOBAL DEFAULT    1 runtime.itabsinit
   216: 00000000004098e0   197 FUNC    GLOBAL DEFAULT    1 runtime.panicdottypeE
   217: 00000000004099c0   217 FUNC    GLOBAL DEFAULT    1 runtime.panicdottypeI
   218: 0000000000409aa0   117 FUNC    GLOBAL DEFAULT    1 runtime.convT
   219: 0000000000409b20   114 FUNC    GLOBAL DEFAULT    1 runtime.convTnoptr
   220: 0000000000409ba0   110 FUNC    GLOBAL DEFAULT    1 runtime.convT64
   221: 0000000000409c20   152 FUNC    GLOBAL DEFAULT    1 runtime.convTstring
   222: 0000000000409cc0   169 FUNC    GLOBAL DEFAULT    1 runtime.convTslice
   223: 0000000000409d80   140 FUNC    GLOBAL DEFAULT    1 runtime.assertE2I2
   224: 0000000000409e20   119 FUNC    GLOBAL DEFAULT    1 runtime.iterate_itabs
   225: 0000000000409ea0    45 FUNC    GLOBAL DEFAULT    1 runtime.unreacha[...]
   226: 0000000000409ee0   325 FUNC    GLOBAL DEFAULT    1 runtime.(*lfstac[...]
   227: 000000000040a040   207 FUNC    GLOBAL DEFAULT    1 runtime.lfnodeVa[...]
   228: 000000000040a120    54 FUNC    GLOBAL DEFAULT    1 runtime.lock
   229: 000000000040a160   399 FUNC    GLOBAL DEFAULT    1 runtime.lock2
   230: 000000000040a300    52 FUNC    GLOBAL DEFAULT    1 runtime.unlock
   231: 000000000040a340   165 FUNC    GLOBAL DEFAULT    1 runtime.unlock2
   232: 000000000040a400   148 FUNC    GLOBAL DEFAULT    1 runtime.notewakeup
   233: 000000000040a4a0   231 FUNC    GLOBAL DEFAULT    1 runtime.notesleep
   234: 000000000040a5a0   452 FUNC    GLOBAL DEFAULT    1 runtime.notetsle[...]
   235: 000000000040a780   107 FUNC    GLOBAL DEFAULT    1 runtime.notetsleep
   236: 000000000040a800   141 FUNC    GLOBAL DEFAULT    1 runtime.notetsleepg
   237: 000000000040a8a0    97 FUNC    GLOBAL DEFAULT    1 runtime.lockRank[...]
   238: 000000000040a920    62 FUNC    GLOBAL DEFAULT    1 runtime.lockWithRank
   239: 000000000040a960    52 FUNC    GLOBAL DEFAULT    1 runtime.unlockWi[...]
   240: 000000000040a9a0   793 FUNC    GLOBAL DEFAULT    1 runtime.mallocinit
   241: 000000000040acc0  1848 FUNC    GLOBAL DEFAULT    1 runtime.(*mheap)[...]
   242: 000000000040b400   261 FUNC    GLOBAL DEFAULT    1 runtime.sysReser[...]
   243: 000000000040b520   532 FUNC    GLOBAL DEFAULT    1 runtime.(*mcache[...]
   244: 000000000040b740  2127 FUNC    GLOBAL DEFAULT    1 runtime.mallocgc
   245: 000000000040bfa0   133 FUNC    GLOBAL DEFAULT    1 runtime.deductAs[...]
   246: 000000000040c040   165 FUNC    GLOBAL DEFAULT    1 runtime.memclrNo[...]
   247: 000000000040c100    66 FUNC    GLOBAL DEFAULT    1 runtime.newobject
   248: 000000000040c160   171 FUNC    GLOBAL DEFAULT    1 runtime.newarray
   249: 000000000040c220   206 FUNC    GLOBAL DEFAULT    1 runtime.profilealloc
   250: 000000000040c300   329 FUNC    GLOBAL DEFAULT    1 runtime.fastexprand
   251: 000000000040c460   169 FUNC    GLOBAL DEFAULT    1 runtime.persiste[...]
   252: 000000000040c520    72 FUNC    GLOBAL DEFAULT    1 runtime.persiste[...]
   253: 000000000040c580   703 FUNC    GLOBAL DEFAULT    1 runtime.persiste[...]
   254: 000000000040c840   238 FUNC    GLOBAL DEFAULT    1 runtime.(*linear[...]
   255: 000000000040c940   752 FUNC    GLOBAL DEFAULT    1 runtime.(*hmap).[...]
   256: 000000000040cc40   116 FUNC    GLOBAL DEFAULT    1 runtime.makemap_small
   257: 000000000040ccc0   421 FUNC    GLOBAL DEFAULT    1 runtime.makemap
   258: 000000000040ce80   589 FUNC    GLOBAL DEFAULT    1 runtime.makeBuck[...]
   259: 000000000040d0e0   553 FUNC    GLOBAL DEFAULT    1 runtime.mapaccess2
   260: 000000000040d320   486 FUNC    GLOBAL DEFAULT    1 runtime.mapaccessK
   261: 000000000040d520  1373 FUNC    GLOBAL DEFAULT    1 runtime.mapassign
   262: 000000000040da80   627 FUNC    GLOBAL DEFAULT    1 runtime.mapiterinit
   263: 000000000040dd00  1295 FUNC    GLOBAL DEFAULT    1 runtime.mapiternext
   264: 000000000040e220   517 FUNC    GLOBAL DEFAULT    1 runtime.hashGrow
   265: 000000000040e440   153 FUNC    GLOBAL DEFAULT    1 runtime.growWork
   266: 000000000040e4e0  1426 FUNC    GLOBAL DEFAULT    1 runtime.evacuate
   267: 000000000040ea80   235 FUNC    GLOBAL DEFAULT    1 runtime.advanceE[...]
   268: 000000000040eb80   400 FUNC    GLOBAL DEFAULT    1 runtime.mapacces[...]
   269: 000000000040ed20   408 FUNC    GLOBAL DEFAULT    1 runtime.mapacces[...]
   270: 000000000040eec0   805 FUNC    GLOBAL DEFAULT    1 runtime.mapassig[...]
   271: 000000000040f200   153 FUNC    GLOBAL DEFAULT    1 runtime.growWork[...]
   272: 000000000040f2a0   947 FUNC    GLOBAL DEFAULT    1 runtime.evacuate[...]
   273: 000000000040f660   410 FUNC    GLOBAL DEFAULT    1 runtime.mapacces[...]
   274: 000000000040f800   825 FUNC    GLOBAL DEFAULT    1 runtime.mapassig[...]
   275: 000000000040fb40   153 FUNC    GLOBAL DEFAULT    1 runtime.growWork[...]
   276: 000000000040fbe0  1061 FUNC    GLOBAL DEFAULT    1 runtime.evacuate[...]
   277: 0000000000410020  1033 FUNC    GLOBAL DEFAULT    1 runtime.mapassig[...]
   278: 0000000000410440   153 FUNC    GLOBAL DEFAULT    1 runtime.growWork[...]
   279: 00000000004104e0  1029 FUNC    GLOBAL DEFAULT    1 runtime.evacuate[...]
   280: 0000000000410900   155 FUNC    GLOBAL DEFAULT    1 runtime.typedmemmove
   281: 00000000004109a0   264 FUNC    GLOBAL DEFAULT    1 runtime.reflectc[...]
   282: 0000000000410ac0   206 FUNC    GLOBAL DEFAULT    1 runtime.typedsli[...]
   283: 0000000000410ba0    91 FUNC    GLOBAL DEFAULT    1 runtime.typedmemclr
   284: 0000000000410c00    62 FUNC    GLOBAL DEFAULT    1 runtime.memclrHa[...]
   285: 0000000000410c40    18 FUNC    GLOBAL DEFAULT    1 runtime.(*mspan)[...]
   286: 0000000000410c60   348 FUNC    GLOBAL DEFAULT    1 runtime.(*mspan)[...]
   287: 0000000000410dc0   549 FUNC    GLOBAL DEFAULT    1 runtime.badPointer
   288: 0000000000411000   316 FUNC    GLOBAL DEFAULT    1 runtime.findObject
   289: 0000000000411140   278 FUNC    GLOBAL DEFAULT    1 runtime.heapBits[...]
   290: 0000000000411260    85 FUNC    GLOBAL DEFAULT    1 runtime.heapBits.next
   291: 00000000004112c0   752 FUNC    GLOBAL DEFAULT    1 runtime.bulkBarr[...]
   292: 00000000004115c0   294 FUNC    GLOBAL DEFAULT    1 runtime.bulkBarr[...]
   293: 0000000000411700   399 FUNC    GLOBAL DEFAULT    1 runtime.bulkBarr[...]
   294: 00000000004118a0   567 FUNC    GLOBAL DEFAULT    1 runtime.typeBits[...]
   295: 0000000000411ae0   332 FUNC    GLOBAL DEFAULT    1 runtime.(*mspan)[...]
   296: 0000000000411c40   294 FUNC    GLOBAL DEFAULT    1 runtime.writeHea[...]
   297: 0000000000411d80   154 FUNC    GLOBAL DEFAULT    1 runtime.writeHea[...]
   298: 0000000000411e20   480 FUNC    GLOBAL DEFAULT    1 runtime.writeHea[...]
   299: 0000000000412000  1423 FUNC    GLOBAL DEFAULT    1 runtime.heapBits[...]
   300: 00000000004125a0   218 FUNC    GLOBAL DEFAULT    1 runtime.progToPo[...]
   301: 0000000000412680  1005 FUNC    GLOBAL DEFAULT    1 runtime.runGCProg
   302: 0000000000412a80   140 FUNC    GLOBAL DEFAULT    1 runtime.material[...]
   303: 0000000000412b20   180 FUNC    GLOBAL DEFAULT    1 runtime.allocmcache
   304: 0000000000412be0   106 FUNC    GLOBAL DEFAULT    1 runtime.allocmca[...]
   305: 0000000000412c60    98 FUNC    GLOBAL DEFAULT    1 runtime.freemcache
   306: 0000000000412ce0   139 FUNC    GLOBAL DEFAULT    1 runtime.freemcac[...]
   307: 0000000000412d80   596 FUNC    GLOBAL DEFAULT    1 runtime.(*mcache[...]
   308: 0000000000412fe0   433 FUNC    GLOBAL DEFAULT    1 runtime.(*mcache[...]
   309: 00000000004131a0   497 FUNC    GLOBAL DEFAULT    1 runtime.(*mcache[...]
   310: 00000000004133a0   222 FUNC    GLOBAL DEFAULT    1 runtime.(*mcache[...]
   311: 0000000000413480   859 FUNC    GLOBAL DEFAULT    1 runtime.(*mcentr[...]
   312: 00000000004137e0   229 FUNC    GLOBAL DEFAULT    1 runtime.(*mcentr[...]
   313: 00000000004138e0   199 FUNC    GLOBAL DEFAULT    1 runtime.(*mcentr[...]
   314: 00000000004139c0   255 FUNC    GLOBAL DEFAULT    1 runtime.startChe[...]
   315: 0000000000413ac0   113 FUNC    GLOBAL DEFAULT    1 runtime.endCheckmarks
   316: 0000000000413b40   492 FUNC    GLOBAL DEFAULT    1 runtime.setCheckmark
   317: 0000000000413d40    79 FUNC    GLOBAL DEFAULT    1 runtime.sysAlloc
   318: 0000000000413da0    83 FUNC    GLOBAL DEFAULT    1 runtime.sysUnused
   319: 0000000000413e00    86 FUNC    GLOBAL DEFAULT    1 runtime.sysUsed
   320: 0000000000413e60    84 FUNC    GLOBAL DEFAULT    1 runtime.sysFree
   321: 0000000000413ec0    99 FUNC    GLOBAL DEFAULT    1 runtime.sysFault
   322: 0000000000413f40    89 FUNC    GLOBAL DEFAULT    1 runtime.sysReserve
   323: 0000000000413fa0   100 FUNC    GLOBAL DEFAULT    1 runtime.sysMap
   324: 0000000000414020   203 FUNC    GLOBAL DEFAULT    1 runtime.sysAllocOS
   325: 0000000000414100   629 FUNC    GLOBAL DEFAULT    1 runtime.sysUnusedOS
   326: 0000000000414380   181 FUNC    GLOBAL DEFAULT    1 runtime.sysUsedOS
   327: 0000000000414440   148 FUNC    GLOBAL DEFAULT    1 runtime.sysHugePageOS
   328: 00000000004144e0   314 FUNC    GLOBAL DEFAULT    1 runtime.sysMapOS
   329: 0000000000414620   619 FUNC    GLOBAL DEFAULT    1 runtime.queuefin[...]
   330: 00000000004148a0    86 FUNC    GLOBAL DEFAULT    1 runtime.createfing
   331: 0000000000414900    87 FUNC    GLOBAL DEFAULT    1 runtime.finalize[...]
   332: 0000000000414960  1148 FUNC    GLOBAL DEFAULT    1 runtime.runfinq
   333: 0000000000414de0  1848 FUNC    GLOBAL DEFAULT    1 runtime.SetFinalizer
   334: 0000000000415520    92 FUNC    GLOBAL DEFAULT    1 runtime.SetFinal[...]
   335: 0000000000415580    54 FUNC    GLOBAL DEFAULT    1 runtime.SetFinal[...]
   336: 00000000004155c0   275 FUNC    GLOBAL DEFAULT    1 runtime.(*fixall[...]
   337: 00000000004156e0   293 FUNC    GLOBAL DEFAULT    1 runtime.(*fixall[...]
   338: 0000000000415820   104 FUNC    GLOBAL DEFAULT    1 runtime.gcinit
   339: 00000000004158a0   229 FUNC    GLOBAL DEFAULT    1 runtime.gcenable
   340: 00000000004159a0    76 FUNC    GLOBAL DEFAULT    1 runtime.gcenable[...]
   341: 0000000000415a00    76 FUNC    GLOBAL DEFAULT    1 runtime.gcenable[...]
   342: 0000000000415a60   177 FUNC    GLOBAL DEFAULT    1 runtime.pollFrac[...]
   343: 0000000000415b20   270 FUNC    GLOBAL DEFAULT    1 runtime.gcTrigge[...]
   344: 0000000000415c40  1413 FUNC    GLOBAL DEFAULT    1 runtime.gcStart
   345: 00000000004161e0   141 FUNC    GLOBAL DEFAULT    1 runtime.gcStart.func2
   346: 0000000000416280   773 FUNC    GLOBAL DEFAULT    1 runtime.gcMarkDone
   347: 00000000004165a0   182 FUNC    GLOBAL DEFAULT    1 runtime.gcMarkDo[...]
   348: 0000000000416660  3429 FUNC    GLOBAL DEFAULT    1 runtime.gcMarkTe[...]
   349: 00000000004173e0    46 FUNC    GLOBAL DEFAULT    1 runtime.gcMarkTe[...]
   350: 0000000000417420   105 FUNC    GLOBAL DEFAULT    1 runtime.gcBgMark[...]
   351: 00000000004174a0  1189 FUNC    GLOBAL DEFAULT    1 runtime.gcBgMark[...]
   352: 0000000000417960   401 FUNC    GLOBAL DEFAULT    1 runtime.gcBgMark[...]
   353: 0000000000417b00  1138 FUNC    GLOBAL DEFAULT    1 runtime.gcMark
   354: 0000000000417f80   489 FUNC    GLOBAL DEFAULT    1 runtime.gcSweep
   355: 0000000000418180   240 FUNC    GLOBAL DEFAULT    1 runtime.gcResetM[...]
   356: 0000000000418280   399 FUNC    GLOBAL DEFAULT    1 runtime.clearpools
   357: 0000000000418420   544 FUNC    GLOBAL DEFAULT    1 runtime.fmtNSAsMS
   358: 0000000000418640   173 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCPUL[...]
   359: 0000000000418700   151 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCPUL[...]
   360: 00000000004187a0   138 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCPUL[...]
   361: 0000000000418840   504 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCPUL[...]
   362: 0000000000418a40   135 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCPUL[...]
   363: 0000000000418ae0    76 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCPUL[...]
   364: 0000000000418b40   217 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCPUL[...]
   365: 0000000000418c20   141 FUNC    GLOBAL DEFAULT    1 runtime.(*limite[...]
   366: 0000000000418cc0   397 FUNC    GLOBAL DEFAULT    1 runtime.(*limite[...]
   367: 0000000000418e60   497 FUNC    GLOBAL DEFAULT    1 runtime.gcMarkRo[...]
   368: 0000000000419060   202 FUNC    GLOBAL DEFAULT    1 runtime.gcMarkRo[...]
   369: 0000000000419140   287 FUNC    GLOBAL DEFAULT    1 runtime.gcMarkRo[...]
   370: 0000000000419260  1080 FUNC    GLOBAL DEFAULT    1 runtime.markroot
   371: 00000000004196a0   310 FUNC    GLOBAL DEFAULT    1 runtime.markroot[...]
   372: 00000000004197e0   176 FUNC    GLOBAL DEFAULT    1 runtime.markrootBlock
   373: 00000000004198a0   266 FUNC    GLOBAL DEFAULT    1 runtime.markroot[...]
   374: 00000000004199c0   869 FUNC    GLOBAL DEFAULT    1 runtime.markrootSpans
   375: 0000000000419d40   731 FUNC    GLOBAL DEFAULT    1 runtime.gcAssistAlloc
   376: 000000000041a020    54 FUNC    GLOBAL DEFAULT    1 runtime.gcAssist[...]
   377: 000000000041a060   901 FUNC    GLOBAL DEFAULT    1 runtime.gcAssist[...]
   378: 000000000041a400   103 FUNC    GLOBAL DEFAULT    1 runtime.gcWakeAl[...]
   379: 000000000041a480   293 FUNC    GLOBAL DEFAULT    1 runtime.gcParkAssist
   380: 000000000041a5c0   389 FUNC    GLOBAL DEFAULT    1 runtime.gcFlushB[...]
   381: 000000000041a760  1829 FUNC    GLOBAL DEFAULT    1 runtime.scanstack
   382: 000000000041aea0    81 FUNC    GLOBAL DEFAULT    1 runtime.scanstac[...]
   383: 000000000041af00   587 FUNC    GLOBAL DEFAULT    1 runtime.scanfram[...]
   384: 000000000041b160  1029 FUNC    GLOBAL DEFAULT    1 runtime.gcDrain
   385: 000000000041b580   553 FUNC    GLOBAL DEFAULT    1 runtime.gcDrainN
   386: 000000000041b7c0   391 FUNC    GLOBAL DEFAULT    1 runtime.scanblock
   387: 000000000041b960   694 FUNC    GLOBAL DEFAULT    1 runtime.scanobject
   388: 000000000041bc20   581 FUNC    GLOBAL DEFAULT    1 runtime.scanCons[...]
   389: 000000000041be80   100 FUNC    GLOBAL DEFAULT    1 runtime.shade
   390: 000000000041bf00   748 FUNC    GLOBAL DEFAULT    1 runtime.greyobject
   391: 000000000041c200  1032 FUNC    GLOBAL DEFAULT    1 runtime.gcDumpObject
   392: 000000000041c620   247 FUNC    GLOBAL DEFAULT    1 runtime.gcmarkne[...]
   393: 000000000041c720   184 FUNC    GLOBAL DEFAULT    1 runtime.gcMarkTi[...]
   394: 000000000041c7e0   167 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCont[...]
   395: 000000000041c8a0   778 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCont[...]
   396: 000000000041cbc0   466 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCont[...]
   397: 000000000041cda0  1003 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCont[...]
   398: 000000000041d1a0   350 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCont[...]
   399: 000000000041d300   691 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCont[...]
   400: 000000000041d5c0   180 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCont[...]
   401: 000000000041d680   156 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCont[...]
   402: 000000000041d720   171 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCont[...]
   403: 000000000041d7e0   134 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCont[...]
   404: 000000000041d880   139 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCont[...]
   405: 000000000041d920   389 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCont[...]
   406: 000000000041dac0   245 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCont[...]
   407: 000000000041dbc0   120 FUNC    GLOBAL DEFAULT    1 runtime.readGOGC
   408: 000000000041dc40   197 FUNC    GLOBAL DEFAULT    1 runtime.readGOME[...]
   409: 000000000041dd20   206 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCont[...]
   410: 000000000041de00   218 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCont[...]
   411: 000000000041dee0   229 FUNC    GLOBAL DEFAULT    1 runtime.(*gcCont[...]
   412: 000000000041dfe0   155 FUNC    GLOBAL DEFAULT    1 runtime.gcContro[...]
   413: 000000000041e080   401 FUNC    GLOBAL DEFAULT    1 runtime.gcPaceSc[...]
   414: 000000000041e220   542 FUNC    GLOBAL DEFAULT    1 runtime.(*scaven[...]
   415: 000000000041e440   133 FUNC    GLOBAL DEFAULT    1 runtime.(*scaven[...]
   416: 000000000041e4e0   143 FUNC    GLOBAL DEFAULT    1 runtime.(*scaven[...]
   417: 000000000041e580   629 FUNC    GLOBAL DEFAULT    1 runtime.(*scaven[...]
   418: 000000000041e800    80 FUNC    GLOBAL DEFAULT    1 runtime.(*scaven[...]
   419: 000000000041e860   431 FUNC    GLOBAL DEFAULT    1 runtime.(*scaven[...]
   420: 000000000041ea20   150 FUNC    GLOBAL DEFAULT    1 runtime.bgscavenge
   421: 000000000041eac0   293 FUNC    GLOBAL DEFAULT    1 runtime.(*pageAl[...]
   422: 000000000041ec00    78 FUNC    GLOBAL DEFAULT    1 runtime.(*pageAl[...]
   423: 000000000041ec60   404 FUNC    GLOBAL DEFAULT    1 runtime.printSca[...]
   424: 000000000041ee00   843 FUNC    GLOBAL DEFAULT    1 runtime.(*pageAl[...]
   425: 000000000041f160   421 FUNC    GLOBAL DEFAULT    1 runtime.fillAligned
   426: 000000000041f320   830 FUNC    GLOBAL DEFAULT    1 runtime.(*palloc[...]
   427: 000000000041f660   372 FUNC    GLOBAL DEFAULT    1 runtime.(*scaven[...]
   428: 000000000041f7e0   464 FUNC    GLOBAL DEFAULT    1 runtime.(*scaven[...]
   429: 000000000041f9c0   268 FUNC    GLOBAL DEFAULT    1 runtime.(*piCont[...]
   430: 000000000041fae0   293 FUNC    GLOBAL DEFAULT    1 runtime.(*stackS[...]
   431: 000000000041fc20   314 FUNC    GLOBAL DEFAULT    1 runtime.(*stackS[...]
   432: 000000000041fd60   389 FUNC    GLOBAL DEFAULT    1 runtime.(*stackS[...]
   433: 000000000041ff00   281 FUNC    GLOBAL DEFAULT    1 runtime.binarySe[...]
   434: 0000000000420020   423 FUNC    GLOBAL DEFAULT    1 runtime.(*mheap)[...]
   435: 00000000004201e0   427 FUNC    GLOBAL DEFAULT    1 runtime.(*active[...]
   436: 00000000004203a0   266 FUNC    GLOBAL DEFAULT    1 runtime.finishsweep_m
   437: 00000000004204c0   381 FUNC    GLOBAL DEFAULT    1 runtime.bgsweep
   438: 0000000000420640   149 FUNC    GLOBAL DEFAULT    1 runtime.(*sweepL[...]
   439: 00000000004206e0   613 FUNC    GLOBAL DEFAULT    1 runtime.sweepone
   440: 0000000000420960   335 FUNC    GLOBAL DEFAULT    1 runtime.(*mspan)[...]
   441: 0000000000420ac0  3320 FUNC    GLOBAL DEFAULT    1 runtime.(*sweepL[...]
   442: 00000000004217c0   153 FUNC    GLOBAL DEFAULT    1 runtime.(*sweepL[...]
   443: 0000000000421860   762 FUNC    GLOBAL DEFAULT    1 runtime.(*mspan)[...]
   444: 0000000000421b60   358 FUNC    GLOBAL DEFAULT    1 runtime.deductSw[...]
   445: 0000000000421ce0   145 FUNC    GLOBAL DEFAULT    1 runtime.gcPaceSweeper
   446: 0000000000421d80    86 FUNC    GLOBAL DEFAULT    1 runtime.(*gcWork[...]
   447: 0000000000421de0   262 FUNC    GLOBAL DEFAULT    1 runtime.(*gcWork).put
   448: 0000000000421f00   486 FUNC    GLOBAL DEFAULT    1 runtime.(*gcWork[...]
   449: 0000000000422100   218 FUNC    GLOBAL DEFAULT    1 runtime.(*gcWork[...]
   450: 00000000004221e0   229 FUNC    GLOBAL DEFAULT    1 runtime.(*gcWork[...]
   451: 00000000004222e0   185 FUNC    GLOBAL DEFAULT    1 runtime.(*gcWork[...]
   452: 00000000004223a0    76 FUNC    GLOBAL DEFAULT    1 runtime.(*workbu[...]
   453: 0000000000422400    76 FUNC    GLOBAL DEFAULT    1 runtime.(*workbu[...]
   454: 0000000000422460   511 FUNC    GLOBAL DEFAULT    1 runtime.getempty
   455: 0000000000422660    76 FUNC    GLOBAL DEFAULT    1 runtime.getempty[...]
   456: 00000000004226c0    76 FUNC    GLOBAL DEFAULT    1 runtime.putempty
   457: 0000000000422720    76 FUNC    GLOBAL DEFAULT    1 runtime.putfull
   458: 0000000000422780   139 FUNC    GLOBAL DEFAULT    1 runtime.trygetfull
   459: 0000000000422820   175 FUNC    GLOBAL DEFAULT    1 runtime.handoff
   460: 00000000004228e0   246 FUNC    GLOBAL DEFAULT    1 runtime.prepareF[...]
   461: 00000000004229e0   212 FUNC    GLOBAL DEFAULT    1 runtime.freeSomeWbufs
   462: 0000000000422ac0   179 FUNC    GLOBAL DEFAULT    1 runtime.freeSome[...]
   463: 0000000000422b80   429 FUNC    GLOBAL DEFAULT    1 runtime.recordspan
   464: 0000000000422d40   115 FUNC    GLOBAL DEFAULT    1 runtime.inHeapOrStack
   465: 0000000000422dc0   121 FUNC    GLOBAL DEFAULT    1 runtime.spanOfHeap
   466: 0000000000422e40   349 FUNC    GLOBAL DEFAULT    1 runtime.(*mheap).init
   467: 0000000000422fa0   556 FUNC    GLOBAL DEFAULT    1 runtime.(*mheap)[...]
   468: 00000000004231e0  1036 FUNC    GLOBAL DEFAULT    1 runtime.(*mheap)[...]
   469: 0000000000423600   168 FUNC    GLOBAL DEFAULT    1 runtime.(*mheap)[...]
   470: 00000000004236c0   126 FUNC    GLOBAL DEFAULT    1 runtime.(*mheap)[...]
   471: 0000000000423740    95 FUNC    GLOBAL DEFAULT    1 runtime.(*mheap)[...]
   472: 00000000004237a0   188 FUNC    GLOBAL DEFAULT    1 runtime.(*mheap)[...]
   473: 0000000000423860   325 FUNC    GLOBAL DEFAULT    1 runtime.(*mheap)[...]
   474: 00000000004239c0   231 FUNC    GLOBAL DEFAULT    1 runtime.(*mheap)[...]
   475: 0000000000423ac0  1406 FUNC    GLOBAL DEFAULT    1 runtime.(*mheap)[...]
   476: 0000000000424040   652 FUNC    GLOBAL DEFAULT    1 runtime.(*mheap)[...]
   477: 00000000004242e0   709 FUNC    GLOBAL DEFAULT    1 runtime.(*mheap).grow
   478: 00000000004245c0   121 FUNC    GLOBAL DEFAULT    1 runtime.(*mheap)[...]
   479: 0000000000424640   103 FUNC    GLOBAL DEFAULT    1 runtime.(*mheap)[...]
   480: 00000000004246c0   133 FUNC    GLOBAL DEFAULT    1 runtime.(*mheap)[...]
   481: 0000000000424760   926 FUNC    GLOBAL DEFAULT    1 runtime.(*mheap)[...]
   482: 0000000000424b00    82 FUNC    GLOBAL DEFAULT    1 runtime.(*mspan).init
   483: 0000000000424b60   361 FUNC    GLOBAL DEFAULT    1 runtime.(*mSpanL[...]
   484: 0000000000424ce0   261 FUNC    GLOBAL DEFAULT    1 runtime.(*mSpanL[...]
   485: 0000000000424e00   533 FUNC    GLOBAL DEFAULT    1 runtime.addspecial
   486: 0000000000425020   436 FUNC    GLOBAL DEFAULT    1 runtime.removespecial
   487: 00000000004251e0   581 FUNC    GLOBAL DEFAULT    1 runtime.addfinalizer
   488: 0000000000425440   159 FUNC    GLOBAL DEFAULT    1 runtime.removefi[...]
   489: 00000000004254e0   168 FUNC    GLOBAL DEFAULT    1 runtime.setprofi[...]
   490: 00000000004255a0   316 FUNC    GLOBAL DEFAULT    1 runtime.freeSpecial
   491: 00000000004256e0   795 FUNC    GLOBAL DEFAULT    1 runtime.newMarkBits
   492: 0000000000425a00    52 FUNC    GLOBAL DEFAULT    1 runtime.newAllocBits
   493: 0000000000425a40   174 FUNC    GLOBAL DEFAULT    1 runtime.nextMark[...]
   494: 0000000000425b00   176 FUNC    GLOBAL DEFAULT    1 runtime.newArena[...]
   495: 0000000000425bc0   366 FUNC    GLOBAL DEFAULT    1 runtime.(*pageAl[...]
   496: 0000000000425d40   558 FUNC    GLOBAL DEFAULT    1 runtime.(*pageAl[...]
   497: 0000000000425f80  1477 FUNC    GLOBAL DEFAULT    1 runtime.(*pageAl[...]
   498: 0000000000426560   942 FUNC    GLOBAL DEFAULT    1 runtime.(*pageAl[...]
   499: 0000000000426920   183 FUNC    GLOBAL DEFAULT    1 runtime.(*pageAl[...]
   500: 00000000004269e0  3031 FUNC    GLOBAL DEFAULT    1 runtime.(*pageAl[...]
   501: 00000000004275c0   389 FUNC    GLOBAL DEFAULT    1 runtime.(*pageAl[...]
   502: 0000000000427760   700 FUNC    GLOBAL DEFAULT    1 runtime.(*pageAl[...]
   503: 0000000000427a20   798 FUNC    GLOBAL DEFAULT    1 runtime.(*pageAl[...]
   504: 0000000000427d40   349 FUNC    GLOBAL DEFAULT    1 runtime.mergeSum[...]
   505: 0000000000427ea0   421 FUNC    GLOBAL DEFAULT    1 runtime.(*pageAl[...]
   506: 0000000000428060  1166 FUNC    GLOBAL DEFAULT    1 runtime.(*pageAl[...]
   507: 0000000000428500   118 FUNC    GLOBAL DEFAULT    1 runtime.(*pageAl[...]
   508: 0000000000428580   116 FUNC    GLOBAL DEFAULT    1 runtime.(*pageAl[...]
   509: 0000000000428600   792 FUNC    GLOBAL DEFAULT    1 runtime.(*scaven[...]
   510: 0000000000428920   180 FUNC    GLOBAL DEFAULT    1 runtime.(*pageCa[...]
   511: 00000000004289e0   302 FUNC    GLOBAL DEFAULT    1 runtime.(*pageCa[...]
   512: 0000000000428b20   500 FUNC    GLOBAL DEFAULT    1 runtime.(*pageCa[...]
   513: 0000000000428d20   732 FUNC    GLOBAL DEFAULT    1 runtime.(*pageAl[...]
   514: 0000000000429000   253 FUNC    GLOBAL DEFAULT    1 runtime.(*pageBi[...]
   515: 0000000000429100   262 FUNC    GLOBAL DEFAULT    1 runtime.(*pageBi[...]
   516: 0000000000429220   542 FUNC    GLOBAL DEFAULT    1 runtime.(*pageBi[...]
   517: 0000000000429440   468 FUNC    GLOBAL DEFAULT    1 runtime.(*palloc[...]
   518: 0000000000429620   167 FUNC    GLOBAL DEFAULT    1 runtime.(*palloc[...]
   519: 00000000004296e0   244 FUNC    GLOBAL DEFAULT    1 runtime.(*palloc[...]
   520: 00000000004297e0   296 FUNC    GLOBAL DEFAULT    1 runtime.(*palloc[...]
   521: 0000000000429920   118 FUNC    GLOBAL DEFAULT    1 runtime.(*palloc[...]
   522: 00000000004299a0   156 FUNC    GLOBAL DEFAULT    1 runtime.newBucket
   523: 0000000000429a40    87 FUNC    GLOBAL DEFAULT    1 runtime.(*bucket).mp
   524: 0000000000429aa0    95 FUNC    GLOBAL DEFAULT    1 runtime.(*bucket).bp
   525: 0000000000429b00   946 FUNC    GLOBAL DEFAULT    1 runtime.stkbucket
   526: 0000000000429ec0   220 FUNC    GLOBAL DEFAULT    1 runtime.mProf_Flush
   527: 0000000000429fa0   197 FUNC    GLOBAL DEFAULT    1 runtime.mProf_Fl[...]
   528: 000000000042a080   474 FUNC    GLOBAL DEFAULT    1 runtime.mProf_Malloc
   529: 000000000042a260    54 FUNC    GLOBAL DEFAULT    1 runtime.mProf_Ma[...]
   530: 000000000042a2a0   206 FUNC    GLOBAL DEFAULT    1 runtime.mProf_Free
   531: 000000000042a380   215 FUNC    GLOBAL DEFAULT    1 runtime.blockevent
   532: 000000000042a460   518 FUNC    GLOBAL DEFAULT    1 runtime.savebloc[...]
   533: 000000000042a680    94 FUNC    GLOBAL DEFAULT    1 runtime.tryRecor[...]
   534: 000000000042a6e0   326 FUNC    GLOBAL DEFAULT    1 runtime.tryRecor[...]
   535: 000000000042a840   346 FUNC    GLOBAL DEFAULT    1 runtime.doRecord[...]
   536: 000000000042a9a0    98 FUNC    GLOBAL DEFAULT    1 runtime.doRecord[...]
   537: 000000000042aa20   158 FUNC    GLOBAL DEFAULT    1 runtime.saveg
   538: 000000000042aac0   575 FUNC    GLOBAL DEFAULT    1 runtime.tracealloc
   539: 000000000042ad00    59 FUNC    GLOBAL DEFAULT    1 runtime.traceall[...]
   540: 000000000042ad40   340 FUNC    GLOBAL DEFAULT    1 runtime.tracefree
   541: 000000000042aea0    59 FUNC    GLOBAL DEFAULT    1 runtime.tracefre[...]
   542: 000000000042aee0   185 FUNC    GLOBAL DEFAULT    1 runtime.tracegc
   543: 000000000042afa0   113 FUNC    GLOBAL DEFAULT    1 runtime.makeAddrRange
   544: 000000000042b020   229 FUNC    GLOBAL DEFAULT    1 runtime.addrRang[...]
   545: 000000000042b120   169 FUNC    GLOBAL DEFAULT    1 runtime.(*addrRa[...]
   546: 000000000042b1e0   252 FUNC    GLOBAL DEFAULT    1 runtime.(*addrRa[...]
   547: 000000000042b2e0   270 FUNC    GLOBAL DEFAULT    1 runtime.(*addrRa[...]
   548: 000000000042b400  1349 FUNC    GLOBAL DEFAULT    1 runtime.(*addrRa[...]
   549: 000000000042b960   436 FUNC    GLOBAL DEFAULT    1 runtime.(*spanSe[...]
   550: 000000000042bb20   303 FUNC    GLOBAL DEFAULT    1 runtime.(*spanSe[...]
   551: 000000000042bc60   325 FUNC    GLOBAL DEFAULT    1 runtime.(*spanSe[...]
   552: 000000000042bdc0   149 FUNC    GLOBAL DEFAULT    1 runtime.(*spanSe[...]
   553: 000000000042be60   165 FUNC    GLOBAL DEFAULT    1 runtime.(*atomic[...]
   554: 000000000042bf20     1 FUNC    GLOBAL DEFAULT    1 runtime.init.4
   555: 000000000042bf40   166 FUNC    GLOBAL DEFAULT    1 runtime.(*sysMem[...]
   556: 000000000042c000   215 FUNC    GLOBAL DEFAULT    1 runtime.(*consis[...]
   557: 000000000042c0e0   142 FUNC    GLOBAL DEFAULT    1 runtime.(*consis[...]
   558: 000000000042c180   119 FUNC    GLOBAL DEFAULT    1 runtime.(*wbBuf)[...]
   559: 000000000042c200   164 FUNC    GLOBAL DEFAULT    1 runtime.wbBufFlush
   560: 000000000042c2c0   687 FUNC    GLOBAL DEFAULT    1 runtime.wbBufFlush1
   561: 000000000042c580   105 FUNC    GLOBAL DEFAULT    1 runtime.netpollG[...]
   562: 000000000042c600   106 FUNC    GLOBAL DEFAULT    1 runtime.(*pollCa[...]
   563: 000000000042c680   235 FUNC    GLOBAL DEFAULT    1 runtime.netpollready
   564: 000000000042c780    38 FUNC    GLOBAL DEFAULT    1 runtime.netpollb[...]
   565: 000000000042c7c0   369 FUNC    GLOBAL DEFAULT    1 runtime.netpollblock
   566: 000000000042c940   175 FUNC    GLOBAL DEFAULT    1 runtime.(*pollCa[...]
   567: 000000000042ca00   442 FUNC    GLOBAL DEFAULT    1 runtime.netpollinit
   568: 000000000042cbc0   135 FUNC    GLOBAL DEFAULT    1 runtime.netpollopen
   569: 000000000042cc60   109 FUNC    GLOBAL DEFAULT    1 runtime.netpollclose
   570: 000000000042cce0   208 FUNC    GLOBAL DEFAULT    1 runtime.netpollBreak
   571: 000000000042cdc0   863 FUNC    GLOBAL DEFAULT    1 runtime.netpoll
   572: 000000000042d120   198 FUNC    GLOBAL DEFAULT    1 runtime.futexsleep
   573: 000000000042d200   173 FUNC    GLOBAL DEFAULT    1 runtime.futexwakeup
   574: 000000000042d2c0   133 FUNC    GLOBAL DEFAULT    1 runtime.futexwak[...]
   575: 000000000042d360   247 FUNC    GLOBAL DEFAULT    1 runtime.getproccount
   576: 000000000042d460   423 FUNC    GLOBAL DEFAULT    1 runtime.newosproc
   577: 000000000042d620   122 FUNC    GLOBAL DEFAULT    1 runtime.newospro[...]
   578: 000000000042d6a0   604 FUNC    GLOBAL DEFAULT    1 runtime.sysargs
   579: 000000000042d900   286 FUNC    GLOBAL DEFAULT    1 runtime.sysauxv
   580: 000000000042da20   298 FUNC    GLOBAL DEFAULT    1 runtime.getHugeP[...]
   581: 000000000042db60   114 FUNC    GLOBAL DEFAULT    1 runtime.osinit
   582: 000000000042dbe0   370 FUNC    GLOBAL DEFAULT    1 runtime.getRandomData
   583: 000000000042dd60   143 FUNC    GLOBAL DEFAULT    1 runtime.mpreinit
   584: 000000000042de00    72 FUNC    GLOBAL DEFAULT    1 runtime.minit
   585: 000000000042de60   129 FUNC    GLOBAL DEFAULT    1 runtime.setsig
   586: 000000000042df00    99 FUNC    GLOBAL DEFAULT    1 runtime.setsigstack
   587: 000000000042df80   126 FUNC    GLOBAL DEFAULT    1 runtime.sysSigaction
   588: 000000000042e000   133 FUNC    GLOBAL DEFAULT    1 runtime.signalM
   589: 000000000042e0a0  1061 FUNC    GLOBAL DEFAULT    1 runtime.setThrea[...]
   590: 000000000042e4e0   868 FUNC    GLOBAL DEFAULT    1 runtime.runPerTh[...]
   591: 000000000042e860   176 FUNC    GLOBAL DEFAULT    1 runtime.panicCheck1
   592: 000000000042e920    98 FUNC    GLOBAL DEFAULT    1 runtime.panicCheck2
   593: 000000000042e9a0   158 FUNC    GLOBAL DEFAULT    1 runtime.goPanicIndex
   594: 000000000042ea40   158 FUNC    GLOBAL DEFAULT    1 runtime.goPanicIndexU
   595: 000000000042eae0   158 FUNC    GLOBAL DEFAULT    1 runtime.goPanicS[...]
   596: 000000000042eb80   158 FUNC    GLOBAL DEFAULT    1 runtime.goPanicS[...]
   597: 000000000042ec20   158 FUNC    GLOBAL DEFAULT    1 runtime.goPanicS[...]
   598: 000000000042ecc0   158 FUNC    GLOBAL DEFAULT    1 runtime.goPanicS[...]
   599: 000000000042ed60   158 FUNC    GLOBAL DEFAULT    1 runtime.goPanicSliceB
   600: 000000000042ee00   158 FUNC    GLOBAL DEFAULT    1 runtime.goPanicS[...]
   601: 000000000042eea0   158 FUNC    GLOBAL DEFAULT    1 runtime.goPanicS[...]
   602: 000000000042ef40   158 FUNC    GLOBAL DEFAULT    1 runtime.goPanicS[...]
   603: 000000000042efe0   158 FUNC    GLOBAL DEFAULT    1 runtime.goPanicS[...]
   604: 000000000042f080    81 FUNC    GLOBAL DEFAULT    1 runtime.panicshift
   605: 000000000042f0e0    77 FUNC    GLOBAL DEFAULT    1 runtime.panicdivide
   606: 000000000042f140   136 FUNC    GLOBAL DEFAULT    1 runtime.deferpro[...]
   607: 000000000042f1e0   542 FUNC    GLOBAL DEFAULT    1 runtime.newdefer
   608: 000000000042f400   710 FUNC    GLOBAL DEFAULT    1 runtime.freedefer
   609: 000000000042f6e0    45 FUNC    GLOBAL DEFAULT    1 runtime.freedefe[...]
   610: 000000000042f720    45 FUNC    GLOBAL DEFAULT    1 runtime.freedeferfn
   611: 000000000042f760   304 FUNC    GLOBAL DEFAULT    1 runtime.deferreturn
   612: 000000000042f8a0   359 FUNC    GLOBAL DEFAULT    1 runtime.preprint[...]
   613: 000000000042fa20   252 FUNC    GLOBAL DEFAULT    1 runtime.printpanics
   614: 000000000042fb20   169 FUNC    GLOBAL DEFAULT    1 runtime.addOneOp[...]
   615: 000000000042fbe0   124 FUNC    GLOBAL DEFAULT    1 runtime.addOneOp[...]
   616: 000000000042fc60   629 FUNC    GLOBAL DEFAULT    1 runtime.addOneOp[...]
   617: 000000000042fee0   767 FUNC    GLOBAL DEFAULT    1 runtime.runOpenD[...]
   618: 00000000004301e0   223 FUNC    GLOBAL DEFAULT    1 runtime.deferCallSave
   619: 00000000004302c0  1844 FUNC    GLOBAL DEFAULT    1 runtime.gopanic
   620: 0000000000430a00     6 FUNC    GLOBAL DEFAULT    1 runtime.getargp
   621: 0000000000430a20    47 FUNC    GLOBAL DEFAULT    1 runtime.gorecover
   622: 0000000000430a60   103 FUNC    GLOBAL DEFAULT    1 runtime.throw
   623: 0000000000430ae0   103 FUNC    GLOBAL DEFAULT    1 runtime.throw.func1
   624: 0000000000430b60   103 FUNC    GLOBAL DEFAULT    1 runtime.fatal
   625: 0000000000430be0   103 FUNC    GLOBAL DEFAULT    1 runtime.fatal.func1
   626: 0000000000430c60   304 FUNC    GLOBAL DEFAULT    1 runtime.recovery
   627: 0000000000430da0   140 FUNC    GLOBAL DEFAULT    1 runtime.fatalthrow
   628: 0000000000430e40   123 FUNC    GLOBAL DEFAULT    1 runtime.fatalthr[...]
   629: 0000000000430ec0   185 FUNC    GLOBAL DEFAULT    1 runtime.fatalpanic
   630: 0000000000430f80   165 FUNC    GLOBAL DEFAULT    1 runtime.fatalpan[...]
   631: 0000000000431040   389 FUNC    GLOBAL DEFAULT    1 runtime.startpanic_m
   632: 00000000004311e0   773 FUNC    GLOBAL DEFAULT    1 runtime.dopanic_m
   633: 0000000000431500   266 FUNC    GLOBAL DEFAULT    1 runtime.canpanic
   634: 0000000000431620   165 FUNC    GLOBAL DEFAULT    1 runtime.shouldPu[...]
   635: 00000000004316e0    53 FUNC    GLOBAL DEFAULT    1 runtime.isAbortPC
   636: 0000000000431720  1425 FUNC    GLOBAL DEFAULT    1 runtime.suspendG
   637: 0000000000431cc0   466 FUNC    GLOBAL DEFAULT    1 runtime.resumeG
   638: 0000000000431ea0    85 FUNC    GLOBAL DEFAULT    1 runtime.asyncPreempt2
   639: 0000000000431f00   187 FUNC    GLOBAL DEFAULT    1 runtime.init.5
   640: 0000000000431fc0   943 FUNC    GLOBAL DEFAULT    1 runtime.isAsyncS[...]
   641: 0000000000432380   366 FUNC    GLOBAL DEFAULT    1 runtime.recordFo[...]
   642: 0000000000432500    98 FUNC    GLOBAL DEFAULT    1 runtime.printlock
   643: 0000000000432580    76 FUNC    GLOBAL DEFAULT    1 runtime.printunlock
   644: 00000000004325e0   313 FUNC    GLOBAL DEFAULT    1 runtime.gwrite
   645: 0000000000432720    54 FUNC    GLOBAL DEFAULT    1 runtime.printsp
   646: 0000000000432760    54 FUNC    GLOBAL DEFAULT    1 runtime.printnl
   647: 00000000004327a0    86 FUNC    GLOBAL DEFAULT    1 runtime.printbool
   648: 0000000000432800   595 FUNC    GLOBAL DEFAULT    1 runtime.printfloat
   649: 0000000000432a60   150 FUNC    GLOBAL DEFAULT    1 runtime.printcomplex
   650: 0000000000432b00   241 FUNC    GLOBAL DEFAULT    1 runtime.printuint
   651: 0000000000432c00    87 FUNC    GLOBAL DEFAULT    1 runtime.printint
   652: 0000000000432c60   273 FUNC    GLOBAL DEFAULT    1 runtime.printhex
   653: 0000000000432d80    52 FUNC    GLOBAL DEFAULT    1 runtime.printpointer
   654: 0000000000432dc0    52 FUNC    GLOBAL DEFAULT    1 runtime.printuintptr
   655: 0000000000432e00   127 FUNC    GLOBAL DEFAULT    1 runtime.printstring
   656: 0000000000432e80   198 FUNC    GLOBAL DEFAULT    1 runtime.printslice
   657: 0000000000432f60   569 FUNC    GLOBAL DEFAULT    1 runtime.hexdumpWords
   658: 00000000004331a0   816 FUNC    GLOBAL DEFAULT    1 runtime.main
   659: 00000000004334e0    53 FUNC    GLOBAL DEFAULT    1 runtime.main.func2
   660: 0000000000433520    54 FUNC    GLOBAL DEFAULT    1 runtime.init.6
   661: 0000000000433560   245 FUNC    GLOBAL DEFAULT    1 runtime.forcegchelper
   662: 0000000000433660    36 FUNC    GLOBAL DEFAULT    1 runtime.Gosched
   663: 00000000004336a0    68 FUNC    GLOBAL DEFAULT    1 runtime.goschedIfBusy
   664: 0000000000433700   302 FUNC    GLOBAL DEFAULT    1 runtime.gopark
   665: 0000000000433840   121 FUNC    GLOBAL DEFAULT    1 runtime.goready
   666: 00000000004338c0    55 FUNC    GLOBAL DEFAULT    1 runtime.goready.func1
   667: 0000000000433900   763 FUNC    GLOBAL DEFAULT    1 runtime.acquireSudog
   668: 0000000000433c00   825 FUNC    GLOBAL DEFAULT    1 runtime.releaseSudog
   669: 0000000000433f40    45 FUNC    GLOBAL DEFAULT    1 runtime.badmcall
   670: 0000000000433f80    45 FUNC    GLOBAL DEFAULT    1 runtime.badmcall2
   671: 0000000000433fc0    47 FUNC    GLOBAL DEFAULT    1 runtime.badrefle[...]
   672: 0000000000434000    47 FUNC    GLOBAL DEFAULT    1 runtime.badmores[...]
   673: 0000000000434040    47 FUNC    GLOBAL DEFAULT    1 runtime.badmores[...]
   674: 0000000000434080    32 FUNC    GLOBAL DEFAULT    1 runtime.badctxt
   675: 00000000004340a0   351 FUNC    GLOBAL DEFAULT    1 runtime.allgadd
   676: 0000000000434200   165 FUNC    GLOBAL DEFAULT    1 runtime.forEachG
   677: 00000000004342c0   135 FUNC    GLOBAL DEFAULT    1 runtime.forEachGRace
   678: 0000000000434360   113 FUNC    GLOBAL DEFAULT    1 runtime.cpuinit
   679: 00000000004343e0   298 FUNC    GLOBAL DEFAULT    1 runtime.getGodeb[...]
   680: 0000000000434520   805 FUNC    GLOBAL DEFAULT    1 runtime.schedinit
   681: 0000000000434860   151 FUNC    GLOBAL DEFAULT    1 runtime.checkmcount
   682: 0000000000434900    94 FUNC    GLOBAL DEFAULT    1 runtime.mReserveID
   683: 0000000000434960   493 FUNC    GLOBAL DEFAULT    1 runtime.mcommoninit
   684: 0000000000434b60   531 FUNC    GLOBAL DEFAULT    1 runtime.ready
   685: 0000000000434d80   202 FUNC    GLOBAL DEFAULT    1 runtime.freezeth[...]
   686: 0000000000434e60   944 FUNC    GLOBAL DEFAULT    1 runtime.casfrom_[...]
   687: 0000000000435220   208 FUNC    GLOBAL DEFAULT    1 runtime.castogsc[...]
   688: 0000000000435300   870 FUNC    GLOBAL DEFAULT    1 runtime.casgstatus
   689: 0000000000435680   144 FUNC    GLOBAL DEFAULT    1 runtime.casgstat[...]
   690: 0000000000435720   138 FUNC    GLOBAL DEFAULT    1 runtime.casGToPr[...]
   691: 00000000004357c0   122 FUNC    GLOBAL DEFAULT    1 runtime.casGFrom[...]
   692: 0000000000435840   218 FUNC    GLOBAL DEFAULT    1 runtime.stopTheWorld
   693: 0000000000435920    94 FUNC    GLOBAL DEFAULT    1 runtime.stopTheW[...]
   694: 0000000000435980   216 FUNC    GLOBAL DEFAULT    1 runtime.startTheWorld
   695: 0000000000435a60   106 FUNC    GLOBAL DEFAULT    1 runtime.stopTheW[...]
   696: 0000000000435ae0    59 FUNC    GLOBAL DEFAULT    1 runtime.startThe[...]
   697: 0000000000435b20   628 FUNC    GLOBAL DEFAULT    1 runtime.stopTheW[...]
   698: 0000000000435da0   486 FUNC    GLOBAL DEFAULT    1 runtime.startThe[...]
   699: 0000000000435fa0   143 FUNC    GLOBAL DEFAULT    1 runtime.mstart0
   700: 0000000000436040   243 FUNC    GLOBAL DEFAULT    1 runtime.mstart1
   701: 0000000000436140    74 FUNC    GLOBAL DEFAULT    1 runtime.mstartm0
   702: 00000000004361a0    67 FUNC    GLOBAL DEFAULT    1 runtime.mPark
   703: 0000000000436200   615 FUNC    GLOBAL DEFAULT    1 runtime.mexit
   704: 0000000000436480   826 FUNC    GLOBAL DEFAULT    1 runtime.forEachP
   705: 00000000004367c0   165 FUNC    GLOBAL DEFAULT    1 runtime.runSafeP[...]
   706: 0000000000436880   710 FUNC    GLOBAL DEFAULT    1 runtime.allocm
   707: 0000000000436b60    59 FUNC    GLOBAL DEFAULT    1 runtime.allocm.func1
   708: 0000000000436ba0   359 FUNC    GLOBAL DEFAULT    1 runtime.needm
   709: 0000000000436d20   124 FUNC    GLOBAL DEFAULT    1 runtime.newextram
   710: 0000000000436da0   472 FUNC    GLOBAL DEFAULT    1 runtime.oneNewExtraM
   711: 0000000000436f80   269 FUNC    GLOBAL DEFAULT    1 runtime.dropm
   712: 00000000004370a0   210 FUNC    GLOBAL DEFAULT    1 runtime.lockextra
   713: 0000000000437180   411 FUNC    GLOBAL DEFAULT    1 runtime.newm
   714: 0000000000437320   271 FUNC    GLOBAL DEFAULT    1 runtime.newm1
   715: 0000000000437440   208 FUNC    GLOBAL DEFAULT    1 runtime.startTem[...]
   716: 0000000000437520   244 FUNC    GLOBAL DEFAULT    1 runtime.template[...]
   717: 0000000000437620   274 FUNC    GLOBAL DEFAULT    1 runtime.stopm
   718: 0000000000437740    12 FUNC    GLOBAL DEFAULT    1 runtime.mspinning
   719: 0000000000437760   664 FUNC    GLOBAL DEFAULT    1 runtime.startm
   720: 0000000000437a00   805 FUNC    GLOBAL DEFAULT    1 runtime.handoffp
   721: 0000000000437d40   330 FUNC    GLOBAL DEFAULT    1 runtime.wakep
   722: 0000000000437ea0   586 FUNC    GLOBAL DEFAULT    1 runtime.stoplockedm
   723: 0000000000438100   175 FUNC    GLOBAL DEFAULT    1 runtime.startlockedm
   724: 00000000004381c0   229 FUNC    GLOBAL DEFAULT    1 runtime.gcstopm
   725: 00000000004382c0   372 FUNC    GLOBAL DEFAULT    1 runtime.execute
   726: 0000000000438440  3077 FUNC    GLOBAL DEFAULT    1 runtime.findRunnable
   727: 0000000000439060   229 FUNC    GLOBAL DEFAULT    1 runtime.pollWork
   728: 0000000000439160   879 FUNC    GLOBAL DEFAULT    1 runtime.stealWork
   729: 00000000004394e0   338 FUNC    GLOBAL DEFAULT    1 runtime.checkRunqsNoP
   730: 0000000000439640   154 FUNC    GLOBAL DEFAULT    1 runtime.checkTim[...]
   731: 00000000004396e0   453 FUNC    GLOBAL DEFAULT    1 runtime.checkIdl[...]
   732: 00000000004398c0    86 FUNC    GLOBAL DEFAULT    1 runtime.wakeNetPoller
   733: 0000000000439920   126 FUNC    GLOBAL DEFAULT    1 runtime.resetspinning
   734: 00000000004399a0   719 FUNC    GLOBAL DEFAULT    1 runtime.injectglist
   735: 0000000000439c80   581 FUNC    GLOBAL DEFAULT    1 runtime.schedule
   736: 0000000000439ee0   461 FUNC    GLOBAL DEFAULT    1 runtime.checkTimers
   737: 000000000043a0c0    71 FUNC    GLOBAL DEFAULT    1 runtime.parkunlock_c
   738: 000000000043a120   335 FUNC    GLOBAL DEFAULT    1 runtime.park_m
   739: 000000000043a280   517 FUNC    GLOBAL DEFAULT    1 runtime.goschedImpl
   740: 000000000043a4a0    76 FUNC    GLOBAL DEFAULT    1 runtime.gosched_m
   741: 000000000043a500   197 FUNC    GLOBAL DEFAULT    1 runtime.goschedg[...]
   742: 000000000043a5e0   112 FUNC    GLOBAL DEFAULT    1 runtime.gopreempt_m
   743: 000000000043a660   645 FUNC    GLOBAL DEFAULT    1 runtime.preemptPark
   744: 000000000043a900   200 FUNC    GLOBAL DEFAULT    1 runtime.goyield_m
   745: 000000000043a9e0    86 FUNC    GLOBAL DEFAULT    1 runtime.goexit1
   746: 000000000043aa40   752 FUNC    GLOBAL DEFAULT    1 runtime.goexit0
   747: 000000000043ad40    91 FUNC    GLOBAL DEFAULT    1 runtime.save
   748: 000000000043ada0   542 FUNC    GLOBAL DEFAULT    1 runtime.reenters[...]
   749: 000000000043afc0   202 FUNC    GLOBAL DEFAULT    1 runtime.reenters[...]
   750: 000000000043b0a0   104 FUNC    GLOBAL DEFAULT    1 runtime.entersys[...]
   751: 000000000043b120   202 FUNC    GLOBAL DEFAULT    1 runtime.entersys[...]
   752: 000000000043b200   414 FUNC    GLOBAL DEFAULT    1 runtime.entersys[...]
   753: 000000000043b3a0   247 FUNC    GLOBAL DEFAULT    1 runtime.entersys[...]
   754: 000000000043b4a0   247 FUNC    GLOBAL DEFAULT    1 runtime.entersys[...]
   755: 000000000043b5a0    98 FUNC    GLOBAL DEFAULT    1 runtime.entersys[...]
   756: 000000000043b620    46 FUNC    GLOBAL DEFAULT    1 runtime.exitsysc[...]
   757: 000000000043b660   237 FUNC    GLOBAL DEFAULT    1 runtime.exitsysc[...]
   758: 000000000043b760   165 FUNC    GLOBAL DEFAULT    1 runtime.exitsysc[...]
   759: 000000000043b820   133 FUNC    GLOBAL DEFAULT    1 runtime.exitsysc[...]
   760: 000000000043b8c0    66 FUNC    GLOBAL DEFAULT    1 runtime.exitsysc[...]
   761: 000000000043b920   171 FUNC    GLOBAL DEFAULT    1 runtime.exitsysc[...]
   762: 000000000043b9e0   381 FUNC    GLOBAL DEFAULT    1 runtime.exitsyscall0
   763: 000000000043bb60   219 FUNC    GLOBAL DEFAULT    1 runtime.malg
   764: 000000000043bc40    71 FUNC    GLOBAL DEFAULT    1 runtime.malg.func1
   765: 000000000043bca0   121 FUNC    GLOBAL DEFAULT    1 runtime.newproc
   766: 000000000043bd20   103 FUNC    GLOBAL DEFAULT    1 runtime.newproc.func1
   767: 000000000043bda0   965 FUNC    GLOBAL DEFAULT    1 runtime.newproc1
   768: 000000000043c180   698 FUNC    GLOBAL DEFAULT    1 runtime.saveAncestors
   769: 000000000043c440   506 FUNC    GLOBAL DEFAULT    1 runtime.gfput
   770: 000000000043c640   465 FUNC    GLOBAL DEFAULT    1 runtime.gfget
   771: 000000000043c820    71 FUNC    GLOBAL DEFAULT    1 runtime.gfget.func2
   772: 000000000043c880    75 FUNC    GLOBAL DEFAULT    1 runtime.gfget.func1
   773: 000000000043c8e0   341 FUNC    GLOBAL DEFAULT    1 runtime.gfpurge
   774: 000000000043ca40   133 FUNC    GLOBAL DEFAULT    1 runtime.unlockOS[...]
   775: 000000000043cae0    45 FUNC    GLOBAL DEFAULT    1 runtime.badunloc[...]
   776: 000000000043cb20    40 FUNC    GLOBAL DEFAULT    1 runtime._System
   777: 000000000043cb60    40 FUNC    GLOBAL DEFAULT    1 runtime._ExternalCode
   778: 000000000043cba0    40 FUNC    GLOBAL DEFAULT    1 runtime._LostExt[...]
   779: 000000000043cbe0    40 FUNC    GLOBAL DEFAULT    1 runtime._GC
   780: 000000000043cc20    40 FUNC    GLOBAL DEFAULT    1 runtime._LostSIG[...]
   781: 000000000043cc60    40 FUNC    GLOBAL DEFAULT    1 runtime._VDSO
   782: 000000000043cca0  1107 FUNC    GLOBAL DEFAULT    1 runtime.sigprof
   783: 000000000043d100   410 FUNC    GLOBAL DEFAULT    1 runtime.(*p).init
   784: 000000000043d2a0   783 FUNC    GLOBAL DEFAULT    1 runtime.(*p).destroy
   785: 000000000043d5c0   202 FUNC    GLOBAL DEFAULT    1 runtime.(*p).des[...]
   786: 000000000043d6a0  2161 FUNC    GLOBAL DEFAULT    1 runtime.procresize
   787: 000000000043df20    85 FUNC    GLOBAL DEFAULT    1 runtime.acquirep
   788: 000000000043df80   286 FUNC    GLOBAL DEFAULT    1 runtime.wirep
   789: 000000000043e0a0   330 FUNC    GLOBAL DEFAULT    1 runtime.releasep
   790: 000000000043e200    95 FUNC    GLOBAL DEFAULT    1 runtime.incidlelocked
   791: 000000000043e260   784 FUNC    GLOBAL DEFAULT    1 runtime.checkdead
   792: 000000000043e580   250 FUNC    GLOBAL DEFAULT    1 runtime.checkdea[...]
   793: 000000000043e680  1081 FUNC    GLOBAL DEFAULT    1 runtime.sysmon
   794: 000000000043eac0   559 FUNC    GLOBAL DEFAULT    1 runtime.retake
   795: 000000000043ed00   133 FUNC    GLOBAL DEFAULT    1 runtime.preemptall
   796: 000000000043eda0   186 FUNC    GLOBAL DEFAULT    1 runtime.preemptone
   797: 000000000043ee60  2280 FUNC    GLOBAL DEFAULT    1 runtime.schedtrace
   798: 000000000043f760   313 FUNC    GLOBAL DEFAULT    1 runtime.schedEna[...]
   799: 000000000043f8a0    77 FUNC    GLOBAL DEFAULT    1 runtime.mput
   800: 000000000043f900   325 FUNC    GLOBAL DEFAULT    1 runtime.globrunqget
   801: 000000000043fa60   229 FUNC    GLOBAL DEFAULT    1 runtime.updateTi[...]
   802: 000000000043fb60   389 FUNC    GLOBAL DEFAULT    1 runtime.pidleput
   803: 000000000043fd00   415 FUNC    GLOBAL DEFAULT    1 runtime.pidleget
   804: 000000000043fea0    84 FUNC    GLOBAL DEFAULT    1 runtime.pidleget[...]
   805: 000000000043ff00   230 FUNC    GLOBAL DEFAULT    1 runtime.runqput
   806: 0000000000440000   476 FUNC    GLOBAL DEFAULT    1 runtime.runqputslow
   807: 00000000004401e0   303 FUNC    GLOBAL DEFAULT    1 runtime.runqputbatch
   808: 0000000000440320   112 FUNC    GLOBAL DEFAULT    1 runtime.runqget
   809: 00000000004403a0   317 FUNC    GLOBAL DEFAULT    1 runtime.runqdrain
   810: 00000000004404e0   399 FUNC    GLOBAL DEFAULT    1 runtime.runqgrab
   811: 0000000000440680   229 FUNC    GLOBAL DEFAULT    1 runtime.runqsteal
   812: 0000000000440780  1276 FUNC    GLOBAL DEFAULT    1 runtime.doInit
   813: 0000000000440c80    75 FUNC    GLOBAL DEFAULT    1 runtime.(*profBu[...]
   814: 0000000000440ce0   159 FUNC    GLOBAL DEFAULT    1 runtime.(*profBu[...]
   815: 0000000000440d80   230 FUNC    GLOBAL DEFAULT    1 runtime.(*profBu[...]
   816: 0000000000440e80  1297 FUNC    GLOBAL DEFAULT    1 runtime.(*profBu[...]
   817: 00000000004413a0   102 FUNC    GLOBAL DEFAULT    1 runtime.(*profBu[...]
   818: 0000000000441420   149 FUNC    GLOBAL DEFAULT    1 runtime.retryOnEAGAIN
   819: 00000000004414c0    98 FUNC    GLOBAL DEFAULT    1 runtime.args
   820: 0000000000441540   274 FUNC    GLOBAL DEFAULT    1 runtime.goargs
   821: 0000000000441660   298 FUNC    GLOBAL DEFAULT    1 runtime.goenvs_unix
   822: 00000000004417a0   500 FUNC    GLOBAL DEFAULT    1 runtime.testAtomic64
   823: 00000000004419a0   976 FUNC    GLOBAL DEFAULT    1 runtime.check
   824: 0000000000441d80   848 FUNC    GLOBAL DEFAULT    1 runtime.parsedeb[...]
   825: 00000000004420e0   293 FUNC    GLOBAL DEFAULT    1 runtime.extendRandom
   826: 0000000000442220    41 FUNC    GLOBAL DEFAULT    1 runtime.waitReas[...]
   827: 0000000000442260   133 FUNC    GLOBAL DEFAULT    1 runtime.(*rwmute[...]
   828: 0000000000442300   165 FUNC    GLOBAL DEFAULT    1 runtime.(*rwmute[...]
   829: 00000000004423c0   208 FUNC    GLOBAL DEFAULT    1 runtime.(*rwmute[...]
   830: 00000000004424a0   121 FUNC    GLOBAL DEFAULT    1 runtime.readyWithTime
   831: 0000000000442520   823 FUNC    GLOBAL DEFAULT    1 runtime.semacquire1
   832: 0000000000442860   529 FUNC    GLOBAL DEFAULT    1 runtime.semrelease1
   833: 0000000000442a80   920 FUNC    GLOBAL DEFAULT    1 runtime.(*semaRo[...]
   834: 0000000000442e20   837 FUNC    GLOBAL DEFAULT    1 runtime.(*semaRo[...]
   835: 0000000000443180   340 FUNC    GLOBAL DEFAULT    1 runtime.(*semaRo[...]
   836: 00000000004432e0   340 FUNC    GLOBAL DEFAULT    1 runtime.(*semaRo[...]
   837: 0000000000443440  1519 FUNC    GLOBAL DEFAULT    1 runtime.dumpregs
   838: 0000000000443a40   216 FUNC    GLOBAL DEFAULT    1 runtime.(*sigctx[...]
   839: 0000000000443b20   454 FUNC    GLOBAL DEFAULT    1 runtime.initsig
   840: 0000000000443d00    86 FUNC    GLOBAL DEFAULT    1 runtime.sigpipe
   841: 0000000000443d60   276 FUNC    GLOBAL DEFAULT    1 runtime.doSigPreempt
   842: 0000000000443e80   551 FUNC    GLOBAL DEFAULT    1 runtime.sigtrampgo
   843: 00000000004440c0   170 FUNC    GLOBAL DEFAULT    1 runtime.sigprofNonGo
   844: 0000000000444180    85 FUNC    GLOBAL DEFAULT    1 runtime.sigprofN[...]
   845: 00000000004441e0   668 FUNC    GLOBAL DEFAULT    1 runtime.adjustSi[...]
   846: 0000000000444480  2313 FUNC    GLOBAL DEFAULT    1 runtime.sighandler
   847: 0000000000444da0  1012 FUNC    GLOBAL DEFAULT    1 runtime.sigpanic
   848: 00000000004451a0   274 FUNC    GLOBAL DEFAULT    1 runtime.dieFromSignal
   849: 00000000004452c0   278 FUNC    GLOBAL DEFAULT    1 runtime.raisebad[...]
   850: 00000000004453e0    34 FUNC    GLOBAL DEFAULT    1 runtime.crash
   851: 0000000000445420   117 FUNC    GLOBAL DEFAULT    1 runtime.noSignalStack
   852: 00000000004454a0   117 FUNC    GLOBAL DEFAULT    1 runtime.sigNotOnStack
   853: 0000000000445520   117 FUNC    GLOBAL DEFAULT    1 runtime.signalDu[...]
   854: 00000000004455a0   155 FUNC    GLOBAL DEFAULT    1 runtime.badsignal
   855: 0000000000445640   401 FUNC    GLOBAL DEFAULT    1 runtime.sigfwdgo
   856: 00000000004457e0   146 FUNC    GLOBAL DEFAULT    1 runtime.sigblock
   857: 0000000000445880   143 FUNC    GLOBAL DEFAULT    1 runtime.unblocksig
   858: 0000000000445920    44 FUNC    GLOBAL DEFAULT    1 runtime.minitSignals
   859: 0000000000445960   318 FUNC    GLOBAL DEFAULT    1 runtime.minitSig[...]
   860: 0000000000445aa0   293 FUNC    GLOBAL DEFAULT    1 runtime.minitSig[...]
   861: 0000000000445be0   149 FUNC    GLOBAL DEFAULT    1 runtime.unminitS[...]
   862: 0000000000445c80    92 FUNC    GLOBAL DEFAULT    1 runtime.signalstack
   863: 0000000000445ce0   358 FUNC    GLOBAL DEFAULT    1 runtime.sigsend
   864: 0000000000445e60   350 FUNC    GLOBAL DEFAULT    1 runtime.makeslicecopy
   865: 0000000000445fc0   217 FUNC    GLOBAL DEFAULT    1 runtime.makeslice
   866: 00000000004460a0  1749 FUNC    GLOBAL DEFAULT    1 runtime.growslice
   867: 0000000000446780   409 FUNC    GLOBAL DEFAULT    1 runtime.stackpoo[...]
   868: 0000000000446920   404 FUNC    GLOBAL DEFAULT    1 runtime.stackpoolfree
   869: 0000000000446ac0   276 FUNC    GLOBAL DEFAULT    1 runtime.stackcac[...]
   870: 0000000000446be0   296 FUNC    GLOBAL DEFAULT    1 runtime.stackcac[...]
   871: 0000000000446d20   229 FUNC    GLOBAL DEFAULT    1 runtime.stackcac[...]
   872: 0000000000446e20   721 FUNC    GLOBAL DEFAULT    1 runtime.stackalloc
   873: 0000000000447100   741 FUNC    GLOBAL DEFAULT    1 runtime.stackfree
   874: 0000000000447400   525 FUNC    GLOBAL DEFAULT    1 runtime.adjustpo[...]
   875: 0000000000447620   869 FUNC    GLOBAL DEFAULT    1 runtime.adjustframe
   876: 00000000004479a0   230 FUNC    GLOBAL DEFAULT    1 runtime.adjustdefers
   877: 0000000000447aa0   367 FUNC    GLOBAL DEFAULT    1 runtime.syncadju[...]
   878: 0000000000447c20   981 FUNC    GLOBAL DEFAULT    1 runtime.copystack
   879: 0000000000448000  3058 FUNC    GLOBAL DEFAULT    1 runtime.newstack
   880: 0000000000448c00     6 FUNC    GLOBAL DEFAULT    1 runtime.nilfunc
   881: 0000000000448c20   120 FUNC    GLOBAL DEFAULT    1 runtime.gostartcallfn
   882: 0000000000448ca0   431 FUNC    GLOBAL DEFAULT    1 runtime.shrinkstack
   883: 0000000000448e60   399 FUNC    GLOBAL DEFAULT    1 runtime.freeStac[...]
   884: 0000000000449000   173 FUNC    GLOBAL DEFAULT    1 runtime.gcComput[...]
   885: 00000000004490c0    92 FUNC    GLOBAL DEFAULT    1 runtime.(*stkfra[...]
   886: 0000000000449120   719 FUNC    GLOBAL DEFAULT    1 runtime.(*stkfra[...]
   887: 0000000000449400  1657 FUNC    GLOBAL DEFAULT    1 runtime.(*stkfra[...]
   888: 0000000000449a80   304 FUNC    GLOBAL DEFAULT    1 runtime.stkobjinit
   889: 0000000000449bc0   561 FUNC    GLOBAL DEFAULT    1 runtime.concatstrings
   890: 0000000000449e00   150 FUNC    GLOBAL DEFAULT    1 runtime.concatstring2
   891: 0000000000449ea0   197 FUNC    GLOBAL DEFAULT    1 runtime.concatstring3
   892: 0000000000449f80   247 FUNC    GLOBAL DEFAULT    1 runtime.concatstring4
   893: 000000000044a080   281 FUNC    GLOBAL DEFAULT    1 runtime.concatstring5
   894: 000000000044a1a0   275 FUNC    GLOBAL DEFAULT    1 runtime.slicebyt[...]
   895: 000000000044a2c0   197 FUNC    GLOBAL DEFAULT    1 runtime.rawstringtmp
   896: 000000000044a3a0   209 FUNC    GLOBAL DEFAULT    1 runtime.atoi64
   897: 000000000044a480   518 FUNC    GLOBAL DEFAULT    1 runtime.parseByt[...]
   898: 000000000044a6a0   184 FUNC    GLOBAL DEFAULT    1 runtime.findnull
   899: 000000000044a760    47 FUNC    GLOBAL DEFAULT    1 runtime.badsyste[...]
   900: 000000000044a7a0    51 FUNC    GLOBAL DEFAULT    1 runtime.fastrand
   901: 000000000044a7e0  1821 FUNC    GLOBAL DEFAULT    1 runtime.(*Frames[...]
   902: 000000000044af00   639 FUNC    GLOBAL DEFAULT    1 runtime.expandCg[...]
   903: 000000000044b180   613 FUNC    GLOBAL DEFAULT    1 runtime.modulesinit
   904: 000000000044b400  2075 FUNC    GLOBAL DEFAULT    1 runtime.moduleda[...]
   905: 000000000044bc20   268 FUNC    GLOBAL DEFAULT    1 runtime.(*module[...]
   906: 000000000044bd40   133 FUNC    GLOBAL DEFAULT    1 runtime.(*Func).Entry
   907: 000000000044bde0    67 FUNC    GLOBAL DEFAULT    1 runtime.(*Func).[...]
   908: 000000000044be40    81 FUNC    GLOBAL DEFAULT    1 runtime.funcInfo[...]
   909: 000000000044bea0   385 FUNC    GLOBAL DEFAULT    1 runtime.findfunc
   910: 000000000044c040  1578 FUNC    GLOBAL DEFAULT    1 runtime.pcvalue
   911: 000000000044c680   153 FUNC    GLOBAL DEFAULT    1 runtime.funcname
   912: 000000000044c720   151 FUNC    GLOBAL DEFAULT    1 runtime.funcpkgpath
   913: 000000000044c7c0   157 FUNC    GLOBAL DEFAULT    1 runtime.funcname[...]
   914: 000000000044c860   223 FUNC    GLOBAL DEFAULT    1 runtime.funcfile
   915: 000000000044c940   278 FUNC    GLOBAL DEFAULT    1 runtime.funcline1
   916: 000000000044ca60    87 FUNC    GLOBAL DEFAULT    1 runtime.funcline
   917: 000000000044cac0   109 FUNC    GLOBAL DEFAULT    1 runtime.funcspdelta
   918: 000000000044cb40   357 FUNC    GLOBAL DEFAULT    1 runtime.funcMaxS[...]
   919: 000000000044ccc0   147 FUNC    GLOBAL DEFAULT    1 runtime.pcdatavalue
   920: 000000000044cd60   152 FUNC    GLOBAL DEFAULT    1 runtime.pcdatavalue1
   921: 000000000044ce00   141 FUNC    GLOBAL DEFAULT    1 runtime.pcdatavalue2
   922: 000000000044cea0   390 FUNC    GLOBAL DEFAULT    1 runtime.step
   923: 000000000044d040   398 FUNC    GLOBAL DEFAULT    1 runtime.doaddtimer
   924: 000000000044d1e0   623 FUNC    GLOBAL DEFAULT    1 runtime.deltimer
   925: 000000000044d460   541 FUNC    GLOBAL DEFAULT    1 runtime.dodeltimer
   926: 000000000044d680   389 FUNC    GLOBAL DEFAULT    1 runtime.dodeltimer0
   927: 000000000044d820  1240 FUNC    GLOBAL DEFAULT    1 runtime.modtimer
   928: 000000000044dd00   542 FUNC    GLOBAL DEFAULT    1 runtime.moveTimers
   929: 000000000044df20   756 FUNC    GLOBAL DEFAULT    1 runtime.adjusttimers
   930: 000000000044e220   204 FUNC    GLOBAL DEFAULT    1 runtime.addAdjus[...]
   931: 000000000044e300   606 FUNC    GLOBAL DEFAULT    1 runtime.runtimer
   932: 000000000044e560   409 FUNC    GLOBAL DEFAULT    1 runtime.runOneTimer
   933: 000000000044e700  1103 FUNC    GLOBAL DEFAULT    1 runtime.clearDel[...]
   934: 000000000044eb60   176 FUNC    GLOBAL DEFAULT    1 runtime.timeSlee[...]
   935: 000000000044ec20   357 FUNC    GLOBAL DEFAULT    1 runtime.siftupTimer
   936: 000000000044eda0   549 FUNC    GLOBAL DEFAULT    1 runtime.siftdownTimer
   937: 000000000044efe0    45 FUNC    GLOBAL DEFAULT    1 runtime.badTimer
   938: 000000000044f020    46 FUNC    GLOBAL DEFAULT    1 runtime.nanotime
   939: 000000000044f060    86 FUNC    GLOBAL DEFAULT    1 runtime.write
   940: 000000000044f0c0   234 FUNC    GLOBAL DEFAULT    1 runtime.traceReader
   941: 000000000044f1c0   175 FUNC    GLOBAL DEFAULT    1 runtime.traceProcFree
   942: 000000000044f280   261 FUNC    GLOBAL DEFAULT    1 runtime.traceEvent
   943: 000000000044f3a0  1413 FUNC    GLOBAL DEFAULT    1 runtime.traceEve[...]
   944: 000000000044f940    66 FUNC    GLOBAL DEFAULT    1 runtime.traceEve[...]
   945: 000000000044f9a0   368 FUNC    GLOBAL DEFAULT    1 runtime.traceCPU[...]
   946: 000000000044fb20   278 FUNC    GLOBAL DEFAULT    1 runtime.traceStackID
   947: 000000000044fc40   147 FUNC    GLOBAL DEFAULT    1 runtime.traceAcq[...]
   948: 000000000044fce0   107 FUNC    GLOBAL DEFAULT    1 runtime.traceRel[...]
   949: 000000000044fd60   634 FUNC    GLOBAL DEFAULT    1 runtime.traceFlush
   950: 000000000044ffe0   429 FUNC    GLOBAL DEFAULT    1 runtime.(*traceS[...]
   951: 00000000004501a0   432 FUNC    GLOBAL DEFAULT    1 runtime.(*traceS[...]
   952: 0000000000450360    81 FUNC    GLOBAL DEFAULT    1 runtime.(*traceS[...]
   953: 00000000004503c0   262 FUNC    GLOBAL DEFAULT    1 runtime.(*traceA[...]
   954: 00000000004504e0    92 FUNC    GLOBAL DEFAULT    1 runtime.tracePro[...]
   955: 0000000000450540   183 FUNC    GLOBAL DEFAULT    1 runtime.traceProcStop
   956: 0000000000450600    92 FUNC    GLOBAL DEFAULT    1 runtime.traceGCS[...]
   957: 0000000000450660   133 FUNC    GLOBAL DEFAULT    1 runtime.traceGCS[...]
   958: 0000000000450700   176 FUNC    GLOBAL DEFAULT    1 runtime.traceGCS[...]
   959: 00000000004507c0   216 FUNC    GLOBAL DEFAULT    1 runtime.traceGoCreate
   960: 00000000004508a0   336 FUNC    GLOBAL DEFAULT    1 runtime.traceGoStart
   961: 0000000000450a00    77 FUNC    GLOBAL DEFAULT    1 runtime.traceGoSched
   962: 0000000000450a60   118 FUNC    GLOBAL DEFAULT    1 runtime.traceGoPark
   963: 0000000000450ae0   198 FUNC    GLOBAL DEFAULT    1 runtime.traceGoUnpark
   964: 0000000000450bc0    59 FUNC    GLOBAL DEFAULT    1 runtime.traceGoS[...]
   965: 0000000000450c00   184 FUNC    GLOBAL DEFAULT    1 runtime.traceGoS[...]
   966: 0000000000450cc0   183 FUNC    GLOBAL DEFAULT    1 runtime.traceGoS[...]
   967: 0000000000450d80    98 FUNC    GLOBAL DEFAULT    1 runtime.traceHea[...]
   968: 0000000000450e00   157 FUNC    GLOBAL DEFAULT    1 runtime.traceHeapGoal
   969: 0000000000450ea0   175 FUNC    GLOBAL DEFAULT    1 runtime.startPCf[...]
   970: 0000000000450f60  7340 FUNC    GLOBAL DEFAULT    1 runtime.gentraceback
   971: 0000000000452c20  1029 FUNC    GLOBAL DEFAULT    1 runtime.printArgs
   972: 0000000000453040   217 FUNC    GLOBAL DEFAULT    1 runtime.printArg[...]
   973: 0000000000453120    71 FUNC    GLOBAL DEFAULT    1 runtime.printArg[...]
   974: 0000000000453180   625 FUNC    GLOBAL DEFAULT    1 runtime.tracebac[...]
   975: 0000000000453400   156 FUNC    GLOBAL DEFAULT    1 runtime.printcre[...]
   976: 00000000004534a0   433 FUNC    GLOBAL DEFAULT    1 runtime.printcre[...]
   977: 0000000000453660    84 FUNC    GLOBAL DEFAULT    1 runtime.traceback
   978: 00000000004536c0   143 FUNC    GLOBAL DEFAULT    1 runtime.tracebacktrap
   979: 0000000000453760   817 FUNC    GLOBAL DEFAULT    1 runtime.traceback1
   980: 0000000000453aa0   427 FUNC    GLOBAL DEFAULT    1 runtime.printAnc[...]
   981: 0000000000453c60   686 FUNC    GLOBAL DEFAULT    1 runtime.printAnc[...]
   982: 0000000000453f20   230 FUNC    GLOBAL DEFAULT    1 runtime.callers
   983: 0000000000454020   115 FUNC    GLOBAL DEFAULT    1 runtime.callers.func1
   984: 00000000004540a0   154 FUNC    GLOBAL DEFAULT    1 runtime.gcallers
   985: 0000000000454140   182 FUNC    GLOBAL DEFAULT    1 runtime.showframe
   986: 0000000000454200   485 FUNC    GLOBAL DEFAULT    1 runtime.showfuncinfo
   987: 0000000000454400   569 FUNC    GLOBAL DEFAULT    1 runtime.goroutin[...]
   988: 0000000000454640   261 FUNC    GLOBAL DEFAULT    1 runtime.tracebac[...]
   989: 0000000000454760   261 FUNC    GLOBAL DEFAULT    1 runtime.tracebac[...]
   990: 0000000000454880   469 FUNC    GLOBAL DEFAULT    1 runtime.tracebac[...]
   991: 0000000000454a60    46 FUNC    GLOBAL DEFAULT    1 runtime.tracebac[...]
   992: 0000000000454aa0   212 FUNC    GLOBAL DEFAULT    1 runtime.isSystem[...]
   993: 0000000000454b80   271 FUNC    GLOBAL DEFAULT    1 runtime.printCgo[...]
   994: 0000000000454ca0   492 FUNC    GLOBAL DEFAULT    1 runtime.printOne[...]
   995: 0000000000454ea0   110 FUNC    GLOBAL DEFAULT    1 runtime.callCgoS[...]
   996: 0000000000454f20   216 FUNC    GLOBAL DEFAULT    1 runtime.cgoContextPCs
   997: 0000000000455000   137 FUNC    GLOBAL DEFAULT    1 runtime.(*_type)[...]
   998: 00000000004550a0   249 FUNC    GLOBAL DEFAULT    1 runtime.(*_type)[...]
   999: 00000000004551a0   645 FUNC    GLOBAL DEFAULT    1 runtime.resolveN[...]
  1000: 0000000000455440   679 FUNC    GLOBAL DEFAULT    1 runtime.resolveT[...]
  1001: 0000000000455700   503 FUNC    GLOBAL DEFAULT    1 runtime.(*_type)[...]
  1002: 0000000000455900   218 FUNC    GLOBAL DEFAULT    1 runtime.name.name
  1003: 00000000004559e0   293 FUNC    GLOBAL DEFAULT    1 runtime.name.tag
  1004: 0000000000455b20   346 FUNC    GLOBAL DEFAULT    1 runtime.name.pkgPath
  1005: 0000000000455c80   137 FUNC    GLOBAL DEFAULT    1 runtime.name.isBlank
  1006: 0000000000455d20  1456 FUNC    GLOBAL DEFAULT    1 runtime.typelinksinit
  1007: 00000000004562e0  3246 FUNC    GLOBAL DEFAULT    1 runtime.typesEqual
  1008: 0000000000456fa0    47 FUNC    GLOBAL DEFAULT    1 runtime.panicuns[...]
  1009: 0000000000456fe0    47 FUNC    GLOBAL DEFAULT    1 runtime.panicuns[...]
  1010: 0000000000457020    47 FUNC    GLOBAL DEFAULT    1 runtime.panicuns[...]
  1011: 0000000000457060    47 FUNC    GLOBAL DEFAULT    1 runtime.panicuns[...]
  1012: 00000000004570a0   475 FUNC    GLOBAL DEFAULT    1 runtime.decoderune
  1013: 0000000000457280  1128 FUNC    GLOBAL DEFAULT    1 runtime.vdsoInit[...]
  1014: 0000000000457700   276 FUNC    GLOBAL DEFAULT    1 runtime.vdsoFind[...]
  1015: 0000000000457820   883 FUNC    GLOBAL DEFAULT    1 runtime.vdsoPars[...]
  1016: 0000000000457ba0   422 FUNC    GLOBAL DEFAULT    1 runtime.vdsoPars[...]
  1017: 0000000000457d60   186 FUNC    GLOBAL DEFAULT    1 runtime.vdsoauxv
  1018: 0000000000457e20   471 FUNC    GLOBAL DEFAULT    1 runtime.schedtra[...]
  1019: 0000000000458000   285 FUNC    GLOBAL DEFAULT    1 runtime.injectgl[...]
  1020: 0000000000458120    44 FUNC    GLOBAL DEFAULT    1 runtime.startThe[...]
  1021: 0000000000458160    58 FUNC    GLOBAL DEFAULT    1 runtime.main.func1
  1022: 00000000004581a0    71 FUNC    GLOBAL DEFAULT    1 runtime.fatalpan[...]
  1023: 0000000000458200   221 FUNC    GLOBAL DEFAULT    1 runtime.preprint[...]
  1024: 00000000004582e0    45 FUNC    GLOBAL DEFAULT    1 runtime.sysSigac[...]
  1025: 0000000000458320    54 FUNC    GLOBAL DEFAULT    1 runtime.wbBufFlu[...]
  1026: 0000000000458360    94 FUNC    GLOBAL DEFAULT    1 runtime.sweepone[...]
  1027: 00000000004583c0   118 FUNC    GLOBAL DEFAULT    1 runtime.(*scaven[...]
  1028: 0000000000458440   189 FUNC    GLOBAL DEFAULT    1 runtime.(*scaven[...]
  1029: 0000000000458500    56 FUNC    GLOBAL DEFAULT    1 runtime.(*scaven[...]
  1030: 0000000000458540     7 FUNC    GLOBAL DEFAULT    1 runtime.(*scaven[...]
  1031: 0000000000458560    19 FUNC    GLOBAL DEFAULT    1 runtime.gcResetM[...]
  1032: 0000000000458580   122 FUNC    GLOBAL DEFAULT    1 runtime.gcBgMark[...]
  1033: 0000000000458600   217 FUNC    GLOBAL DEFAULT    1 runtime.gcMarkTe[...]
  1034: 00000000004586e0    54 FUNC    GLOBAL DEFAULT    1 runtime.gcMarkTe[...]
  1035: 0000000000458720    56 FUNC    GLOBAL DEFAULT    1 runtime.gcMarkTe[...]
  1036: 0000000000458760    54 FUNC    GLOBAL DEFAULT    1 runtime.gcMarkTe[...]
  1037: 00000000004587a0   116 FUNC    GLOBAL DEFAULT    1 runtime.gcMarkDo[...]
  1038: 0000000000458820   108 FUNC    GLOBAL DEFAULT    1 runtime.gcMarkDo[...]
  1039: 00000000004588a0    98 FUNC    GLOBAL DEFAULT    1 runtime.gcMarkDo[...]
  1040: 0000000000458920    40 FUNC    GLOBAL DEFAULT    1 runtime.gcStart.func1
  1041: 0000000000458960   169 FUNC    GLOBAL DEFAULT    1 runtime.runExitH[...]
  1042: 0000000000458a20    71 FUNC    GLOBAL DEFAULT    1 runtime.runExitH[...]
  1043: 0000000000458a80   373 FUNC    GLOBAL DEFAULT    1 runtime.debugCal[...]
  1044: 0000000000458c00   181 FUNC    GLOBAL DEFAULT    1 runtime.debugCal[...]
  1045: 0000000000458cc0    13 FUNC    GLOBAL DEFAULT    1 runtime.(*mheap)[...]
  1046: 0000000000458ce0   159 FUNC    GLOBAL DEFAULT    1 runtime.(*pageAl[...]
  1047: 0000000000458d80   313 FUNC    GLOBAL DEFAULT    1 runtime.init
  1048: 0000000000458ec0    68 FUNC    GLOBAL DEFAULT    1 sync/atomic.Stor[...]
  1049: 0000000000458f20    91 FUNC    GLOBAL DEFAULT    1 sync/atomic.Comp[...]
  1050: 0000000000458f80    12 FUNC    GLOBAL DEFAULT    1 reflect.chanlen
  1051: 0000000000458fa0    66 FUNC    GLOBAL DEFAULT    1 reflect.unsafe_New
  1052: 0000000000459000    72 FUNC    GLOBAL DEFAULT    1 reflect.mapiterinit
  1053: 0000000000459060    52 FUNC    GLOBAL DEFAULT    1 reflect.mapiternext
  1054: 00000000004590a0     4 FUNC    GLOBAL DEFAULT    1 reflect.mapiterkey
  1055: 00000000004590c0     5 FUNC    GLOBAL DEFAULT    1 reflect.mapiterelem
  1056: 00000000004590e0    12 FUNC    GLOBAL DEFAULT    1 reflect.maplen
  1057: 0000000000459100    72 FUNC    GLOBAL DEFAULT    1 reflect.typedmemmove
  1058: 0000000000459160    62 FUNC    GLOBAL DEFAULT    1 reflect.typedmemclr
  1059: 00000000004591a0   104 FUNC    GLOBAL DEFAULT    1 reflect.verifyNo[...]
  1060: 0000000000459220    76 FUNC    GLOBAL DEFAULT    1 sync.runtime_reg[...]
  1061: 0000000000459280   180 FUNC    GLOBAL DEFAULT    1 sync.event
  1062: 0000000000459340    40 FUNC    GLOBAL DEFAULT    1 internal/poll.ru[...]
  1063: 0000000000459380   431 FUNC    GLOBAL DEFAULT    1 internal/poll.ru[...]
  1064: 0000000000459540   188 FUNC    GLOBAL DEFAULT    1 internal/poll.ru[...]
  1065: 0000000000459600   116 FUNC    GLOBAL DEFAULT    1 internal/poll.ru[...]
  1066: 0000000000459680   293 FUNC    GLOBAL DEFAULT    1 internal/poll.ru[...]
  1067: 00000000004597c0   581 FUNC    GLOBAL DEFAULT    1 internal/poll.ru[...]
  1068: 0000000000459a20    58 FUNC    GLOBAL DEFAULT    1 sync.throw
  1069: 0000000000459a60    67 FUNC    GLOBAL DEFAULT    1 sync.fatal
  1070: 0000000000459ac0    39 FUNC    GLOBAL DEFAULT    1 runtime.entersyscall
  1071: 0000000000459b00   601 FUNC    GLOBAL DEFAULT    1 runtime.exitsyscall
  1072: 0000000000459d60    21 FUNC    GLOBAL DEFAULT    1 sync.runtime_procPin
  1073: 0000000000459d80    12 FUNC    GLOBAL DEFAULT    1 sync.runtime_pro[...]
  1074: 0000000000459da0   104 FUNC    GLOBAL DEFAULT    1 sync.runtime_canSpin
  1075: 0000000000459e20    49 FUNC    GLOBAL DEFAULT    1 sync.runtime_doSpin
  1076: 0000000000459e60   202 FUNC    GLOBAL DEFAULT    1 syscall.runtime_envs
  1077: 0000000000459f40   202 FUNC    GLOBAL DEFAULT    1 os.runtime_args
  1078: 000000000045a020   287 FUNC    GLOBAL DEFAULT    1 runtime/debug.Se[...]
  1079: 000000000045a140    60 FUNC    GLOBAL DEFAULT    1 reflect.resolveN[...]
  1080: 000000000045a180    60 FUNC    GLOBAL DEFAULT    1 reflect.resolveT[...]
  1081: 000000000045a1c0    60 FUNC    GLOBAL DEFAULT    1 reflect.resolveT[...]
  1082: 000000000045a200    60 FUNC    GLOBAL DEFAULT    1 internal/reflect[...]
  1083: 000000000045a240   314 FUNC    GLOBAL DEFAULT    1 reflect.addReflectOff
  1084: 000000000045a380    66 FUNC    GLOBAL DEFAULT    1 internal/poll.ru[...]
  1085: 000000000045a3e0    71 FUNC    GLOBAL DEFAULT    1 sync.runtime_Sem[...]
  1086: 000000000045a440    86 FUNC    GLOBAL DEFAULT    1 sync.runtime_Sem[...]
  1087: 000000000045a4a0    56 FUNC    GLOBAL DEFAULT    1 internal/poll.ru[...]
  1088: 000000000045a4e0   159 FUNC    GLOBAL DEFAULT    1 sync.runtime_not[...]
  1089: 000000000045a580    59 FUNC    GLOBAL DEFAULT    1 sync.runtime_nanotime
  1090: 000000000045a5c0    71 FUNC    GLOBAL DEFAULT    1 os.sigpipe
  1091: 000000000045a620    32 FUNC    GLOBAL DEFAULT    1 runtime.morestackc
  1092: 000000000045a640   175 FUNC    GLOBAL DEFAULT    1 runtime.gostring
  1093: 000000000045a700    72 FUNC    GLOBAL DEFAULT    1 reflect.memmove
  1094: 000000000045b360    14 FUNC    GLOBAL DEFAULT    1 _rt0_amd64
  1095: 000000000045b380   370 FUNC    GLOBAL DEFAULT    1 runtime.rt0_go.abi0
  1096: 000000000045b500     1 FUNC    GLOBAL DEFAULT    1 runtime.asminit.abi0
  1097: 000000000045b520     6 FUNC    GLOBAL DEFAULT    1 runtime.mstart.abi0
  1098: 000000000045b540    17 FUNC    GLOBAL DEFAULT    1 runtime.gogo.abi0
  1099: 000000000045b560    74 FUNC    GLOBAL DEFAULT    1 runtime.mcall
  1100: 000000000045b5c0     1 FUNC    GLOBAL DEFAULT    1 runtime.systemst[...]
  1101: 000000000045b5e0   134 FUNC    GLOBAL DEFAULT    1 runtime.systemst[...]
  1102: 000000000045b680   145 FUNC    GLOBAL DEFAULT    1 runtime.morestac[...]
  1103: 000000000045b720    10 FUNC    GLOBAL DEFAULT    1 runtime.morestac[...]
  1104: 000000000045b740   174 FUNC    GLOBAL DEFAULT    1 runtime.spillArg[...]
  1105: 000000000045b800   174 FUNC    GLOBAL DEFAULT    1 runtime.unspillA[...]
  1106: 000000000045b8c0   490 FUNC    GLOBAL DEFAULT    1 runtime.reflectc[...]
  1107: 000000000045bac0   155 FUNC    GLOBAL DEFAULT    1 runtime.call16.abi0
  1108: 000000000045bb60   155 FUNC    GLOBAL DEFAULT    1 runtime.call32.abi0
  1109: 000000000045bc00   155 FUNC    GLOBAL DEFAULT    1 runtime.call64.abi0
  1110: 000000000045bca0   220 FUNC    GLOBAL DEFAULT    1 runtime.call128.abi0
  1111: 000000000045bd80   223 FUNC    GLOBAL DEFAULT    1 runtime.call256.abi0
  1112: 000000000045be60   223 FUNC    GLOBAL DEFAULT    1 runtime.call512.abi0
  1113: 000000000045bf40   223 FUNC    GLOBAL DEFAULT    1 runtime.call1024.abi0
  1114: 000000000045c020   223 FUNC    GLOBAL DEFAULT    1 runtime.call2048.abi0
  1115: 000000000045c100   231 FUNC    GLOBAL DEFAULT    1 runtime.call4096.abi0
  1116: 000000000045c200   231 FUNC    GLOBAL DEFAULT    1 runtime.call8192.abi0
  1117: 000000000045c300   231 FUNC    GLOBAL DEFAULT    1 runtime.call1638[...]
  1118: 000000000045c400   231 FUNC    GLOBAL DEFAULT    1 runtime.call3276[...]
  1119: 000000000045c500   231 FUNC    GLOBAL DEFAULT    1 runtime.call6553[...]
  1120: 000000000045c600   231 FUNC    GLOBAL DEFAULT    1 runtime.call1310[...]
  1121: 000000000045c700   231 FUNC    GLOBAL DEFAULT    1 runtime.call2621[...]
  1122: 000000000045c800   231 FUNC    GLOBAL DEFAULT    1 runtime.call5242[...]
  1123: 000000000045c900   231 FUNC    GLOBAL DEFAULT    1 runtime.call1048[...]
  1124: 000000000045ca00   231 FUNC    GLOBAL DEFAULT    1 runtime.call2097[...]
  1125: 000000000045cb00   231 FUNC    GLOBAL DEFAULT    1 runtime.call4194[...]
  1126: 000000000045cc00   231 FUNC    GLOBAL DEFAULT    1 runtime.call8388[...]
  1127: 000000000045cd00   231 FUNC    GLOBAL DEFAULT    1 runtime.call1677[...]
  1128: 000000000045ce00   231 FUNC    GLOBAL DEFAULT    1 runtime.call3355[...]
  1129: 000000000045cf00   231 FUNC    GLOBAL DEFAULT    1 runtime.call6710[...]
  1130: 000000000045d000   231 FUNC    GLOBAL DEFAULT    1 runtime.call1342[...]
  1131: 000000000045d100   231 FUNC    GLOBAL DEFAULT    1 runtime.call2684[...]
  1132: 000000000045d200   231 FUNC    GLOBAL DEFAULT    1 runtime.call5368[...]
  1133: 000000000045d300   231 FUNC    GLOBAL DEFAULT    1 runtime.call1073[...]
  1134: 000000000045d400    12 FUNC    GLOBAL DEFAULT    1 runtime.procyiel[...]
  1135: 000000000045d420     1 FUNC    GLOBAL DEFAULT    1 runtime.publicat[...]
  1136: 000000000045d440   174 FUNC    GLOBAL DEFAULT    1 runtime.asmcgoca[...]
  1137: 000000000045d500    15 FUNC    GLOBAL DEFAULT    1 runtime.setg.abi0
  1138: 000000000045d520     4 FUNC    GLOBAL DEFAULT    1 runtime.abort.abi0
  1139: 000000000045d540    31 FUNC    GLOBAL DEFAULT    1 runtime.stackche[...]
  1140: 000000000045d560    35 FUNC    GLOBAL DEFAULT    1 runtime.cputicks.abi0
  1141: 000000000045d5a0    19 FUNC    GLOBAL DEFAULT    1 runtime.memhash
  1142: 000000000045d5c0    26 FUNC    GLOBAL DEFAULT    1 runtime.strhash
  1143: 000000000045d5e0    58 FUNC    GLOBAL DEFAULT    1 runtime.memhash32
  1144: 000000000045d620    59 FUNC    GLOBAL DEFAULT    1 runtime.memhash64
  1145: 000000000045d660    29 FUNC    GLOBAL DEFAULT    1 runtime.checkASM.abi0
  1146: 000000000045d680     6 FUNC    GLOBAL DEFAULT    1 runtime.return0.abi0
  1147: 000000000045d6a0     7 FUNC    GLOBAL DEFAULT    1 runtime.goexit.abi0
  1148: 000000000045d6c0    18 FUNC    GLOBAL DEFAULT    1 runtime.sigpanic[...]
  1149: 000000000045d6e0   227 FUNC    GLOBAL DEFAULT    1 runtime.gcWriteB[...]
  1150: 000000000045d7e0    10 FUNC    GLOBAL DEFAULT    1 runtime.gcWriteB[...]
  1151: 000000000045d800    10 FUNC    GLOBAL DEFAULT    1 runtime.gcWriteB[...]
  1152: 000000000045d820    10 FUNC    GLOBAL DEFAULT    1 runtime.gcWriteB[...]
  1153: 000000000045d840    10 FUNC    GLOBAL DEFAULT    1 runtime.gcWriteB[...]
  1154: 000000000045d860    10 FUNC    GLOBAL DEFAULT    1 runtime.gcWriteB[...]
  1155: 000000000045d880    10 FUNC    GLOBAL DEFAULT    1 runtime.gcWriteB[...]
  1156: 000000000045d8a0   649 FUNC    GLOBAL DEFAULT    1 runtime.debugCallV2
  1157: 000000000045db40    51 FUNC    GLOBAL DEFAULT    1 runtime.debugCal[...]
  1158: 000000000045db80     8 FUNC    GLOBAL DEFAULT    1 runtime.panicIndex
  1159: 000000000045dba0     8 FUNC    GLOBAL DEFAULT    1 runtime.panicIndexU
  1160: 000000000045dbc0    11 FUNC    GLOBAL DEFAULT    1 runtime.panicSli[...]
  1161: 000000000045dbe0    11 FUNC    GLOBAL DEFAULT    1 runtime.panicSli[...]
  1162: 000000000045dc00    11 FUNC    GLOBAL DEFAULT    1 runtime.panicSli[...]
  1163: 000000000045dc20    11 FUNC    GLOBAL DEFAULT    1 runtime.panicSli[...]
  1164: 000000000045dc40     8 FUNC    GLOBAL DEFAULT    1 runtime.panicSliceB
  1165: 000000000045dc60     8 FUNC    GLOBAL DEFAULT    1 runtime.panicSliceBU
  1166: 000000000045dc80     8 FUNC    GLOBAL DEFAULT    1 runtime.panicSli[...]
  1167: 000000000045dca0     8 FUNC    GLOBAL DEFAULT    1 runtime.panicSli[...]
  1168: 000000000045dcc0     8 FUNC    GLOBAL DEFAULT    1 runtime.panicSlice3C
  1169: 000000000045dce0   369 FUNC    GLOBAL DEFAULT    1 runtime.duffzero
  1170: 000000000045de60   897 FUNC    GLOBAL DEFAULT    1 runtime.duffcopy
  1171: 000000000045e200   743 FUNC    GLOBAL DEFAULT    1 runtime.memclrNo[...]
  1172: 000000000045e500  1727 FUNC    GLOBAL DEFAULT    1 runtime.memmove
  1173: 000000000045ebc0   431 FUNC    GLOBAL DEFAULT    1 runtime.asyncPre[...]
  1174: 000000000045ed80     5 FUNC    GLOBAL DEFAULT    1 _rt0_amd64_linux
  1175: 000000000045ee20    12 FUNC    GLOBAL DEFAULT    1 runtime.exit.abi0
  1176: 000000000045ee40    27 FUNC    GLOBAL DEFAULT    1 runtime.exitThre[...]
  1177: 000000000045ee60    44 FUNC    GLOBAL DEFAULT    1 runtime.open.abi0
  1178: 000000000045eea0    29 FUNC    GLOBAL DEFAULT    1 runtime.closefd.abi0
  1179: 000000000045eec0    26 FUNC    GLOBAL DEFAULT    1 runtime.write1.abi0
  1180: 000000000045eee0    25 FUNC    GLOBAL DEFAULT    1 runtime.read.abi0
  1181: 000000000045ef00    21 FUNC    GLOBAL DEFAULT    1 runtime.pipe2.abi0
  1182: 000000000045ef20    71 FUNC    GLOBAL DEFAULT    1 runtime.usleep.abi0
  1183: 000000000045ef80    12 FUNC    GLOBAL DEFAULT    1 runtime.gettid.abi0
  1184: 000000000045efa0    34 FUNC    GLOBAL DEFAULT    1 runtime.raise.abi0
  1185: 000000000045efe0    21 FUNC    GLOBAL DEFAULT    1 runtime.raisepro[...]
  1186: 000000000045f000    13 FUNC    GLOBAL DEFAULT    1 runtime.getpid.abi0
  1187: 000000000045f020    23 FUNC    GLOBAL DEFAULT    1 runtime.tgkill.abi0
  1188: 000000000045f040    26 FUNC    GLOBAL DEFAULT    1 runtime.timer_cr[...]
  1189: 000000000045f060    30 FUNC    GLOBAL DEFAULT    1 runtime.timer_se[...]
  1190: 000000000045f080    16 FUNC    GLOBAL DEFAULT    1 runtime.timer_de[...]
  1191: 000000000045f0a0    27 FUNC    GLOBAL DEFAULT    1 runtime.mincore.abi0
  1192: 000000000045f0c0   186 FUNC    GLOBAL DEFAULT    1 runtime.nanotime[...]
  1193: 000000000045f180    46 FUNC    GLOBAL DEFAULT    1 runtime.rtsigpro[...]
  1194: 000000000045f1c0    32 FUNC    GLOBAL DEFAULT    1 runtime.rt_sigac[...]
  1195: 000000000045f1e0    62 FUNC    GLOBAL DEFAULT    1 runtime.callCgoS[...]
  1196: 000000000045f220    34 FUNC    GLOBAL DEFAULT    1 runtime.sigfwd.abi0
  1197: 000000000045f260   108 FUNC    GLOBAL DEFAULT    1 runtime.sigtramp.abi0
  1198: 000000000045f2e0   187 FUNC    GLOBAL DEFAULT    1 runtime.cgoSigtr[...]
  1199: 000000000045f3a0    11 FUNC    GLOBAL DEFAULT    1 runtime.sigretur[...]
  1200: 000000000045f3c0    80 FUNC    GLOBAL DEFAULT    1 runtime.sysMmap.abi0
  1201: 000000000045f420    81 FUNC    GLOBAL DEFAULT    1 runtime.callCgoM[...]
  1202: 000000000045f480    39 FUNC    GLOBAL DEFAULT    1 runtime.sysMunma[...]
  1203: 000000000045f4c0    58 FUNC    GLOBAL DEFAULT    1 runtime.callCgoM[...]
  1204: 000000000045f500    28 FUNC    GLOBAL DEFAULT    1 runtime.madvise.abi0
  1205: 000000000045f520    40 FUNC    GLOBAL DEFAULT    1 runtime.futex.abi0
  1206: 000000000045f560   157 FUNC    GLOBAL DEFAULT    1 runtime.clone.abi0
  1207: 000000000045f600    39 FUNC    GLOBAL DEFAULT    1 runtime.sigaltst[...]
  1208: 000000000045f640    66 FUNC    GLOBAL DEFAULT    1 runtime.settls.abi0
  1209: 000000000045f6a0     8 FUNC    GLOBAL DEFAULT    1 runtime.osyield.abi0
  1210: 000000000045f6c0    27 FUNC    GLOBAL DEFAULT    1 runtime.sched_ge[...]
  1211: 000000000045f6e0   231 FUNC    GLOBAL DEFAULT    1 time.now
  1212: 000000000045f7e0   167 FUNC    GLOBAL DEFAULT    1 runtime.(*lockRa[...]
  1213: 000000000045f8a0   109 FUNC    GLOBAL DEFAULT    1 runtime.(*waitRe[...]
  1214: 000000000045f920   115 FUNC    GLOBAL DEFAULT    1 runtime.(*errorS[...]
  1215: 000000000045f9a0   115 FUNC    GLOBAL DEFAULT    1 runtime.(*errorA[...]
  1216: 000000000045fa20    68 FUNC    GLOBAL DEFAULT    1 runtime.(*plainE[...]
  1217: 000000000045fa80   110 FUNC    GLOBAL DEFAULT    1 runtime.(*bounds[...]
  1218: 000000000045fb00    93 FUNC    GLOBAL DEFAULT    1 runtime.(*itabTa[...]
  1219: 000000000045fb60    57 FUNC    GLOBAL DEFAULT    1 runtime.debugCal[...]
  1220: 000000000045fba0    47 FUNC    GLOBAL DEFAULT    1 runtime.debugCal[...]
  1221: 000000000045fbe0    67 FUNC    GLOBAL DEFAULT    1 runtime.reflectc[...]
  1222: 000000000045fc40    52 FUNC    GLOBAL DEFAULT    1 runtime.wbBufFlu[...]
  1223: 000000000045fc80    18 FUNC    GLOBAL DEFAULT    1 runtime.osinit.abi0
  1224: 000000000045fca0    19 FUNC    GLOBAL DEFAULT    1 runtime.osyield
  1225: 000000000045fcc0    18 FUNC    GLOBAL DEFAULT    1 runtime.asyncPre[...]
  1226: 000000000045fce0    47 FUNC    GLOBAL DEFAULT    1 runtime.badmcall.abi0
  1227: 000000000045fd20    47 FUNC    GLOBAL DEFAULT    1 runtime.badmcall[...]
  1228: 000000000045fd60    18 FUNC    GLOBAL DEFAULT    1 runtime.badrefle[...]
  1229: 000000000045fd80    18 FUNC    GLOBAL DEFAULT    1 runtime.badmores[...]
  1230: 000000000045fda0    18 FUNC    GLOBAL DEFAULT    1 runtime.badmores[...]
  1231: 000000000045fdc0    18 FUNC    GLOBAL DEFAULT    1 runtime.schedini[...]
  1232: 000000000045fde0    18 FUNC    GLOBAL DEFAULT    1 runtime.mstart0.abi0
  1233: 000000000045fe00    18 FUNC    GLOBAL DEFAULT    1 runtime.goexit1.abi0
  1234: 000000000045fe20    47 FUNC    GLOBAL DEFAULT    1 runtime.newproc.abi0
  1235: 000000000045fe60    51 FUNC    GLOBAL DEFAULT    1 runtime.args.abi0
  1236: 000000000045fea0    18 FUNC    GLOBAL DEFAULT    1 runtime.check.abi0
  1237: 000000000045fec0    18 FUNC    GLOBAL DEFAULT    1 runtime.newstack.abi0
  1238: 000000000045fee0    18 FUNC    GLOBAL DEFAULT    1 runtime.morestac[...]
  1239: 000000000045ff00    18 FUNC    GLOBAL DEFAULT    1 runtime.badsyste[...]
  1240: 000000000045ff20   100 FUNC    GLOBAL DEFAULT    1 runtime.reflectcall
  1241: 000000000045ffa0    55 FUNC    GLOBAL DEFAULT    1 runtime.asmcgocall
  1242: 000000000045ffe0    60 FUNC    GLOBAL DEFAULT    1 runtime.write.abi0
  1243: 0000000000460020    81 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
  1244: 0000000000460080    41 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime.itab
  1245: 00000000004600c0   154 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
  1246: 0000000000460160    21 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
  1247: 0000000000460180   212 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
  1248: 0000000000460260   148 FUNC    GLOBAL DEFAULT    1 type:.eq.[2]runt[...]
  1249: 0000000000460300   119 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
  1250: 0000000000460380   198 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
  1251: 0000000000460460    23 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
  1252: 0000000000460480     9 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
  1253: 00000000004604a0    33 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
  1254: 00000000004604e0    67 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
  1255: 0000000000460540    41 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
  1256: 0000000000460580    39 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
  1257: 00000000004605c0    67 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
  1258: 0000000000460620    21 FUNC    GLOBAL DEFAULT    1 type:.eq.struct [...]
  1259: 0000000000460640   117 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
  1260: 00000000004606c0   113 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
  1261: 0000000000460740     6 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
  1262: 0000000000460760    10 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
  1263: 0000000000460780    67 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
  1264: 00000000004607e0    67 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
  1265: 0000000000460840    29 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime.mOS
  1266: 0000000000460860   100 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
  1267: 00000000004608e0   132 FUNC    GLOBAL DEFAULT    1 type:.eq.runtime[...]
  1268: 0000000000460980    15 FUNC    GLOBAL DEFAULT    1 sync/atomic.Comp[...]
  1269: 00000000004609a0     3 FUNC    GLOBAL DEFAULT    1 sync/atomic.Stor[...]
  1270: 00000000004609c0     4 FUNC    GLOBAL DEFAULT    1 sync/atomic.Stor[...]
  1271: 00000000004609e0   201 FUNC    GLOBAL DEFAULT    1 internal/reflect[...]
  1272: 0000000000460ac0   107 FUNC    GLOBAL DEFAULT    1 internal/reflect[...]
  1273: 0000000000460b40   137 FUNC    GLOBAL DEFAULT    1 internal/reflect[...]
  1274: 0000000000460be0   295 FUNC    GLOBAL DEFAULT    1 internal/reflect[...]
  1275: 0000000000460d20    89 FUNC    GLOBAL DEFAULT    1 internal/reflect[...]
  1276: 0000000000460d80    39 FUNC    GLOBAL DEFAULT    1 type:.eq.interna[...]
  1277: 0000000000460dc0   133 FUNC    GLOBAL DEFAULT    1 internal/reflect[...]
  1278: 0000000000460e60    11 FUNC    GLOBAL DEFAULT    1 errors.(*errorSt[...]
  1279: 0000000000460e80   127 FUNC    GLOBAL DEFAULT    1 errors.init
  1280: 0000000000460f00    94 FUNC    GLOBAL DEFAULT    1 sort.Stable
  1281: 0000000000460f60   279 FUNC    GLOBAL DEFAULT    1 sort.insertionSort
  1282: 0000000000461080   335 FUNC    GLOBAL DEFAULT    1 sort.stable
  1283: 00000000004611e0   884 FUNC    GLOBAL DEFAULT    1 sort.symMerge
  1284: 0000000000461560   453 FUNC    GLOBAL DEFAULT    1 sort.rotate
  1285: 0000000000461740    27 FUNC    GLOBAL DEFAULT    1 math.init
  1286: 0000000000461760   422 FUNC    GLOBAL DEFAULT    1 unicode/utf8.Dec[...]
  1287: 0000000000461920   422 FUNC    GLOBAL DEFAULT    1 unicode/utf8.Dec[...]
  1288: 0000000000461ae0   343 FUNC    GLOBAL DEFAULT    1 unicode/utf8.Enc[...]
  1289: 0000000000461c40   381 FUNC    GLOBAL DEFAULT    1 unicode/utf8.app[...]
  1290: 0000000000461dc0   351 FUNC    GLOBAL DEFAULT    1 unicode/utf8.Run[...]
  1291: 0000000000461f20   351 FUNC    GLOBAL DEFAULT    1 unicode/utf8.Run[...]
  1292: 0000000000462080   310 FUNC    GLOBAL DEFAULT    1 strconv.(*decima[...]
  1293: 00000000004621c0   499 FUNC    GLOBAL DEFAULT    1 strconv.rightShift
  1294: 00000000004623c0   536 FUNC    GLOBAL DEFAULT    1 strconv.leftShift
  1295: 00000000004625e0   180 FUNC    GLOBAL DEFAULT    1 strconv.(*decima[...]
  1296: 00000000004626a0   295 FUNC    GLOBAL DEFAULT    1 strconv.(*decima[...]
  1297: 00000000004627e0  1838 FUNC    GLOBAL DEFAULT    1 strconv.genericFtoa
  1298: 0000000000462f20  1035 FUNC    GLOBAL DEFAULT    1 strconv.bigFtoa
  1299: 0000000000463340   570 FUNC    GLOBAL DEFAULT    1 strconv.formatDigits
  1300: 0000000000463580  1279 FUNC    GLOBAL DEFAULT    1 strconv.roundShortest
  1301: 0000000000463a80  1165 FUNC    GLOBAL DEFAULT    1 strconv.fmtE
  1302: 0000000000463f20   727 FUNC    GLOBAL DEFAULT    1 strconv.fmtF
  1303: 0000000000464200   357 FUNC    GLOBAL DEFAULT    1 strconv.fmtB
  1304: 0000000000464380  1711 FUNC    GLOBAL DEFAULT    1 strconv.fmtX
  1305: 0000000000464a40   677 FUNC    GLOBAL DEFAULT    1 strconv.ryuFtoaF[...]
  1306: 0000000000464d00   664 FUNC    GLOBAL DEFAULT    1 strconv.ryuFtoaF[...]
  1307: 0000000000464fa0   727 FUNC    GLOBAL DEFAULT    1 strconv.formatDecimal
  1308: 0000000000465280  1320 FUNC    GLOBAL DEFAULT    1 strconv.ryuFtoaS[...]
  1309: 00000000004657c0   765 FUNC    GLOBAL DEFAULT    1 strconv.ryuDigits
  1310: 0000000000465ac0   501 FUNC    GLOBAL DEFAULT    1 strconv.ryuDigits32
  1311: 0000000000465cc0   216 FUNC    GLOBAL DEFAULT    1 strconv.mult64bi[...]
  1312: 0000000000465da0   285 FUNC    GLOBAL DEFAULT    1 strconv.mult128b[...]
  1313: 0000000000465ec0   233 FUNC    GLOBAL DEFAULT    1 strconv.FormatInt
  1314: 0000000000465fc0  1311 FUNC    GLOBAL DEFAULT    1 strconv.formatBits
  1315: 00000000004664e0  1013 FUNC    GLOBAL DEFAULT    1 strconv.appendQu[...]
  1316: 00000000004668e0   261 FUNC    GLOBAL DEFAULT    1 strconv.appendQu[...]
  1317: 0000000000466a00  1861 FUNC    GLOBAL DEFAULT    1 strconv.appendEs[...]
  1318: 0000000000467160   237 FUNC    GLOBAL DEFAULT    1 strconv.CanBackquote
  1319: 0000000000467260   764 FUNC    GLOBAL DEFAULT    1 strconv.IsPrint
  1320: 0000000000467560   197 FUNC    GLOBAL DEFAULT    1 strconv.init
  1321: 0000000000467640   441 FUNC    GLOBAL DEFAULT    1 sync.(*Map).Load
  1322: 0000000000467800   837 FUNC    GLOBAL DEFAULT    1 sync.(*Map).Load[...]
  1323: 0000000000467b60   338 FUNC    GLOBAL DEFAULT    1 sync.(*entry).tr[...]
  1324: 0000000000467cc0   220 FUNC    GLOBAL DEFAULT    1 sync.(*Map).miss[...]
  1325: 0000000000467da0   379 FUNC    GLOBAL DEFAULT    1 sync.(*Map).dirt[...]
  1326: 0000000000467f20   114 FUNC    GLOBAL DEFAULT    1 sync.(*entry).tr[...]
  1327: 0000000000467fa0   655 FUNC    GLOBAL DEFAULT    1 sync.(*Mutex).lo[...]
  1328: 0000000000468240    68 FUNC    GLOBAL DEFAULT    1 sync.(*Mutex).Unlock
  1329: 00000000004682a0   222 FUNC    GLOBAL DEFAULT    1 sync.(*Mutex).un[...]
  1330: 0000000000468380   308 FUNC    GLOBAL DEFAULT    1 sync.(*Once).doSlow
  1331: 00000000004684c0    44 FUNC    GLOBAL DEFAULT    1 sync.(*Once).doS[...]
  1332: 0000000000468500    76 FUNC    GLOBAL DEFAULT    1 sync.(*Once).doS[...]
  1333: 0000000000468560   197 FUNC    GLOBAL DEFAULT    1 sync.(*Pool).Put
  1334: 0000000000468640   215 FUNC    GLOBAL DEFAULT    1 sync.(*Pool).Get
  1335: 0000000000468720   422 FUNC    GLOBAL DEFAULT    1 sync.(*Pool).getSlow
  1336: 00000000004688e0   108 FUNC    GLOBAL DEFAULT    1 sync.(*Pool).pin
  1337: 0000000000468960   665 FUNC    GLOBAL DEFAULT    1 sync.(*Pool).pinSlow
  1338: 0000000000468c00    76 FUNC    GLOBAL DEFAULT    1 sync.(*Pool).pin[...]
  1339: 0000000000468c60   344 FUNC    GLOBAL DEFAULT    1 sync.poolCleanup
  1340: 0000000000468dc0    54 FUNC    GLOBAL DEFAULT    1 sync.init.0
  1341: 0000000000468e00   249 FUNC    GLOBAL DEFAULT    1 sync.(*poolDeque[...]
  1342: 0000000000468f00   314 FUNC    GLOBAL DEFAULT    1 sync.(*poolDeque[...]
  1343: 0000000000469040   336 FUNC    GLOBAL DEFAULT    1 sync.(*poolDeque[...]
  1344: 00000000004691a0   473 FUNC    GLOBAL DEFAULT    1 sync.(*poolChain[...]
  1345: 0000000000469380   100 FUNC    GLOBAL DEFAULT    1 sync.(*poolChain[...]
  1346: 0000000000469400   199 FUNC    GLOBAL DEFAULT    1 sync.(*poolChain[...]
  1347: 00000000004694e0    47 FUNC    GLOBAL DEFAULT    1 sync.init.1
  1348: 0000000000469520    86 FUNC    GLOBAL DEFAULT    1 sync.init
  1349: 0000000000469580    10 FUNC    GLOBAL DEFAULT    1 type:.eq.sync/at[...]
  1350: 00000000004695a0    10 FUNC    GLOBAL DEFAULT    1 type:.eq.sync.entry
  1351: 00000000004695c0   148 FUNC    GLOBAL DEFAULT    1 type:.eq.sync.po[...]
  1352: 0000000000469660   154 FUNC    GLOBAL DEFAULT    1 type:.eq.sync.po[...]
  1353: 0000000000469700 15321 FUNC    GLOBAL DEFAULT    1 unicode.init
  1354: 000000000046d2e0   479 FUNC    GLOBAL DEFAULT    1 internal/itoa.Itoa
  1355: 000000000046d4c0   756 FUNC    GLOBAL DEFAULT    1 reflect.(*abiSeq[...]
  1356: 000000000046d7c0   588 FUNC    GLOBAL DEFAULT    1 reflect.(*abiSeq[...]
  1357: 000000000046da20  1689 FUNC    GLOBAL DEFAULT    1 reflect.(*abiSeq[...]
  1358: 000000000046e0c0   567 FUNC    GLOBAL DEFAULT    1 reflect.(*abiSeq[...]
  1359: 000000000046e300  2348 FUNC    GLOBAL DEFAULT    1 reflect.newAbiDesc
  1360: 000000000046ec40   158 FUNC    GLOBAL DEFAULT    1 reflect.intFromReg
  1361: 000000000046ece0   158 FUNC    GLOBAL DEFAULT    1 reflect.intToReg
  1362: 000000000046ed80   757 FUNC    GLOBAL DEFAULT    1 reflect.makeMeth[...]
  1363: 000000000046f080   191 FUNC    GLOBAL DEFAULT    1 reflect.moveMake[...]
  1364: 000000000046f140   201 FUNC    GLOBAL DEFAULT    1 reflect.name.name
  1365: 000000000046f220   293 FUNC    GLOBAL DEFAULT    1 reflect.name.tag
  1366: 000000000046f360  1165 FUNC    GLOBAL DEFAULT    1 reflect.newName
  1367: 000000000046f800   133 FUNC    GLOBAL DEFAULT    1 reflect.Kind.String
  1368: 000000000046f8a0   137 FUNC    GLOBAL DEFAULT    1 reflect.(*rtype)[...]
  1369: 000000000046f940     8 FUNC    GLOBAL DEFAULT    1 reflect.(*rtype).Kind
  1370: 000000000046f960   189 FUNC    GLOBAL DEFAULT    1 reflect.(*rtype)[...]
  1371: 000000000046fa20    87 FUNC    GLOBAL DEFAULT    1 reflect.(*rtype)[...]
  1372: 000000000046fa80   329 FUNC    GLOBAL DEFAULT    1 reflect.(*rtype).Elem
  1373: 000000000046fbe0   312 FUNC    GLOBAL DEFAULT    1 reflect.(*rtype)[...]
  1374: 000000000046fd20   120 FUNC    GLOBAL DEFAULT    1 reflect.(*rtype).Len
  1375: 000000000046fda0   175 FUNC    GLOBAL DEFAULT    1 reflect.ChanDir.[...]
  1376: 000000000046fe60   408 FUNC    GLOBAL DEFAULT    1 reflect.(*struct[...]
  1377: 0000000000470000  1701 FUNC    GLOBAL DEFAULT    1 reflect.funcLayout
  1378: 00000000004706c0    56 FUNC    GLOBAL DEFAULT    1 reflect.funcLayo[...]
  1379: 0000000000470700  1453 FUNC    GLOBAL DEFAULT    1 reflect.addTypeBits
  1380: 0000000000470cc0   219 FUNC    GLOBAL DEFAULT    1 reflect.packEface
  1381: 0000000000470da0   279 FUNC    GLOBAL DEFAULT    1 reflect.(*ValueE[...]
  1382: 0000000000470ec0   613 FUNC    GLOBAL DEFAULT    1 reflect.valueMet[...]
  1383: 0000000000471140   190 FUNC    GLOBAL DEFAULT    1 reflect.Value.pa[...]
  1384: 0000000000471200   398 FUNC    GLOBAL DEFAULT    1 reflect.Value.by[...]
  1385: 00000000004713a0   767 FUNC    GLOBAL DEFAULT    1 reflect.methodRe[...]
  1386: 00000000004716a0  4217 FUNC    GLOBAL DEFAULT    1 reflect.callMethod
  1387: 0000000000472720   462 FUNC    GLOBAL DEFAULT    1 reflect.Value.Elem
  1388: 0000000000472900   280 FUNC    GLOBAL DEFAULT    1 reflect.Value.Field
  1389: 0000000000472a20   440 FUNC    GLOBAL DEFAULT    1 reflect.Value.Index
  1390: 0000000000472be0   325 FUNC    GLOBAL DEFAULT    1 reflect.valueInt[...]
  1391: 0000000000472d40    17 FUNC    GLOBAL DEFAULT    1 reflect.Value.Kind
  1392: 0000000000472d60   118 FUNC    GLOBAL DEFAULT    1 reflect.Value.Len
  1393: 0000000000472de0   441 FUNC    GLOBAL DEFAULT    1 reflect.Value.le[...]
  1394: 0000000000472fa0   180 FUNC    GLOBAL DEFAULT    1 reflect.(*MapIte[...]
  1395: 0000000000473060   180 FUNC    GLOBAL DEFAULT    1 reflect.(*MapIte[...]
  1396: 0000000000473120   230 FUNC    GLOBAL DEFAULT    1 reflect.(*MapIte[...]
  1397: 0000000000473220   168 FUNC    GLOBAL DEFAULT    1 reflect.flag.pan[...]
  1398: 00000000004732e0   165 FUNC    GLOBAL DEFAULT    1 reflect.copyVal
  1399: 00000000004733a0   165 FUNC    GLOBAL DEFAULT    1 reflect.Value.Nu[...]
  1400: 0000000000473460   197 FUNC    GLOBAL DEFAULT    1 reflect.Value.Nu[...]
  1401: 0000000000473540   430 FUNC    GLOBAL DEFAULT    1 reflect.Value.Pointer
  1402: 0000000000473700   663 FUNC    GLOBAL DEFAULT    1 reflect.Value.Slice
  1403: 00000000004739a0   118 FUNC    GLOBAL DEFAULT    1 reflect.Value.String
  1404: 0000000000473a20   197 FUNC    GLOBAL DEFAULT    1 reflect.Value.st[...]
  1405: 0000000000473b00   345 FUNC    GLOBAL DEFAULT    1 reflect.Value.ty[...]
  1406: 0000000000473c60   229 FUNC    GLOBAL DEFAULT    1 reflect.init
  1407: 0000000000473d60   168 FUNC    GLOBAL DEFAULT    1 reflect.methodVa[...]
  1408: 0000000000473e20    39 FUNC    GLOBAL DEFAULT    1 type:.eq.reflect[...]
  1409: 0000000000473e60   248 FUNC    GLOBAL DEFAULT    1 type:.eq.reflect[...]
  1410: 0000000000473f60   183 FUNC    GLOBAL DEFAULT    1 reflect.(*Kind).[...]
  1411: 0000000000474020    95 FUNC    GLOBAL DEFAULT    1 reflect.(*ChanDi[...]
  1412: 0000000000474080    66 FUNC    GLOBAL DEFAULT    1 reflect.(*Value).Kind
  1413: 00000000004740e0   130 FUNC    GLOBAL DEFAULT    1 reflect.(*Value).Len
  1414: 0000000000474180   140 FUNC    GLOBAL DEFAULT    1 reflect.(*Value)[...]
  1415: 0000000000474220    83 FUNC    GLOBAL DEFAULT    1 reflect.(*funcTy[...]
  1416: 0000000000474280   293 FUNC    GLOBAL DEFAULT    1 reflect.(*funcTy[...]
  1417: 00000000004743c0    34 FUNC    GLOBAL DEFAULT    1 reflect.(*funcTy[...]
  1418: 0000000000474400    83 FUNC    GLOBAL DEFAULT    1 reflect.(*funcTy[...]
  1419: 0000000000474460    83 FUNC    GLOBAL DEFAULT    1 reflect.(*funcTy[...]
  1420: 00000000004744c0    52 FUNC    GLOBAL DEFAULT    1 reflect.moveMake[...]
  1421: 0000000000474500   108 FUNC    GLOBAL DEFAULT    1 reflect.callMeth[...]
  1422: 0000000000474580   100 FUNC    GLOBAL DEFAULT    1 type:.eq.reflect[...]
  1423: 0000000000474600    43 FUNC    GLOBAL DEFAULT    1 type:.eq.reflect[...]
  1424: 0000000000474640    91 FUNC    GLOBAL DEFAULT    1 type:.eq.reflect[...]
  1425: 00000000004746a0     5 FUNC    GLOBAL DEFAULT    1 internal/fmtsort[...]
  1426: 00000000004746c0   165 FUNC    GLOBAL DEFAULT    1 internal/fmtsort[...]
  1427: 0000000000474780   485 FUNC    GLOBAL DEFAULT    1 internal/fmtsort[...]
  1428: 0000000000474980  1113 FUNC    GLOBAL DEFAULT    1 internal/fmtsort.Sort
  1429: 0000000000474de0  4076 FUNC    GLOBAL DEFAULT    1 internal/fmtsort[...]
  1430: 0000000000475de0   648 FUNC    GLOBAL DEFAULT    1 internal/fmtsort[...]
  1431: 0000000000476080   717 FUNC    GLOBAL DEFAULT    1 io.init
  1432: 0000000000476360   421 FUNC    GLOBAL DEFAULT    1 internal/oserror.init
  1433: 0000000000476520   150 FUNC    GLOBAL DEFAULT    1 syscall.SetNonblock
  1434: 00000000004765c0   133 FUNC    GLOBAL DEFAULT    1 syscall.Errno.Error
  1435: 0000000000476660   178 FUNC    GLOBAL DEFAULT    1 syscall.Close
  1436: 0000000000476720   210 FUNC    GLOBAL DEFAULT    1 syscall.fcntl
  1437: 0000000000476800   245 FUNC    GLOBAL DEFAULT    1 syscall.write
  1438: 0000000000476900   188 FUNC    GLOBAL DEFAULT    1 syscall.munmap
  1439: 00000000004769c0   197 FUNC    GLOBAL DEFAULT    1 syscall.Getrlimit
  1440: 0000000000476aa0   197 FUNC    GLOBAL DEFAULT    1 syscall.Setrlimit
  1441: 0000000000476b80   247 FUNC    GLOBAL DEFAULT    1 syscall.mmap
  1442: 0000000000476c80   120 FUNC    GLOBAL DEFAULT    1 syscall.init
  1443: 0000000000476d00    37 FUNC    GLOBAL DEFAULT    1 syscall.RawSyscall
  1444: 0000000000476d40   117 FUNC    GLOBAL DEFAULT    1 syscall.Syscall
  1445: 0000000000476dc0   139 FUNC    GLOBAL DEFAULT    1 syscall.Syscall6
  1446: 0000000000476e60    95 FUNC    GLOBAL DEFAULT    1 syscall.(*Errno)[...]
  1447: 0000000000476ec0   785 FUNC    GLOBAL DEFAULT    1 time.init
  1448: 00000000004771e0   118 FUNC    GLOBAL DEFAULT    1 path.init
  1449: 0000000000477260   143 FUNC    GLOBAL DEFAULT    1 io/fs.(*PathErro[...]
  1450: 0000000000477300   453 FUNC    GLOBAL DEFAULT    1 io/fs.init
  1451: 00000000004774e0   187 FUNC    GLOBAL DEFAULT    1 type:.eq.io/fs.P[...]
  1452: 00000000004775a0   118 FUNC    GLOBAL DEFAULT    1 internal/syscall[...]
  1453: 0000000000477620    13 FUNC    GLOBAL DEFAULT    1 internal/poll.er[...]
  1454: 0000000000477640    13 FUNC    GLOBAL DEFAULT    1 internal/poll.(*[...]
  1455: 0000000000477660   272 FUNC    GLOBAL DEFAULT    1 internal/poll.(*[...]
  1456: 0000000000477780   332 FUNC    GLOBAL DEFAULT    1 internal/poll.(*[...]
  1457: 00000000004778e0   261 FUNC    GLOBAL DEFAULT    1 internal/poll.(*[...]
  1458: 0000000000477a00   147 FUNC    GLOBAL DEFAULT    1 internal/poll.(*[...]
  1459: 0000000000477aa0    78 FUNC    GLOBAL DEFAULT    1 internal/poll.(*[...]
  1460: 0000000000477b00   244 FUNC    GLOBAL DEFAULT    1 internal/poll.(*[...]
  1461: 0000000000477c00   301 FUNC    GLOBAL DEFAULT    1 internal/poll.(*[...]
  1462: 0000000000477d40   345 FUNC    GLOBAL DEFAULT    1 internal/poll.(*[...]
  1463: 0000000000477ea0   168 FUNC    GLOBAL DEFAULT    1 internal/poll.(*[...]
  1464: 0000000000477f60   156 FUNC    GLOBAL DEFAULT    1 internal/poll.(*[...]
  1465: 0000000000478000   207 FUNC    GLOBAL DEFAULT    1 internal/poll.(*[...]
  1466: 00000000004780e0  1289 FUNC    GLOBAL DEFAULT    1 internal/poll.(*[...]
  1467: 0000000000478600    76 FUNC    GLOBAL DEFAULT    1 internal/poll.(*[...]
  1468: 0000000000478660   269 FUNC    GLOBAL DEFAULT    1 internal/poll.init
  1469: 0000000000478780    76 FUNC    GLOBAL DEFAULT    1 internal/poll.(*[...]
  1470: 00000000004787e0    67 FUNC    GLOBAL DEFAULT    1 type:.eq.interna[...]
  1471: 0000000000478840   118 FUNC    GLOBAL DEFAULT    1 internal/safefil[...]
  1472: 00000000004788c0   138 FUNC    GLOBAL DEFAULT    1 os.glob..func1
  1473: 0000000000478960   662 FUNC    GLOBAL DEFAULT    1 os.(*File).Write
  1474: 0000000000478c00   133 FUNC    GLOBAL DEFAULT    1 os.NewFile
  1475: 0000000000478ca0   549 FUNC    GLOBAL DEFAULT    1 os.newFile
  1476: 0000000000478ee0   478 FUNC    GLOBAL DEFAULT    1 os.(*file).close
  1477: 00000000004790c0    85 FUNC    GLOBAL DEFAULT    1 os.init.0
  1478: 0000000000479120    95 FUNC    GLOBAL DEFAULT    1 os.init.1
  1479: 0000000000479180   793 FUNC    GLOBAL DEFAULT    1 os.init
  1480: 00000000004794a0   165 FUNC    GLOBAL DEFAULT    1 type:.eq.os.file
  1481: 0000000000479560   364 FUNC    GLOBAL DEFAULT    1 fmt.(*fmt).write[...]
  1482: 00000000004796e0   693 FUNC    GLOBAL DEFAULT    1 fmt.(*fmt).pad
  1483: 00000000004799a0   683 FUNC    GLOBAL DEFAULT    1 fmt.(*fmt).padString
  1484: 0000000000479c60    98 FUNC    GLOBAL DEFAULT    1 fmt.(*fmt).fmtBoolean
  1485: 0000000000479ce0   869 FUNC    GLOBAL DEFAULT    1 fmt.(*fmt).fmtUnicode
  1486: 000000000047a060  1428 FUNC    GLOBAL DEFAULT    1 fmt.(*fmt).fmtInteger
  1487: 000000000047a600   293 FUNC    GLOBAL DEFAULT    1 fmt.(*fmt).truncate
  1488: 000000000047a740   229 FUNC    GLOBAL DEFAULT    1 fmt.(*fmt).fmtS
  1489: 000000000047a840   113 FUNC    GLOBAL DEFAULT    1 fmt.(*fmt).fmtBs
  1490: 000000000047a8c0  1298 FUNC    GLOBAL DEFAULT    1 fmt.(*fmt).fmtSbx
  1491: 000000000047ade0   508 FUNC    GLOBAL DEFAULT    1 fmt.(*fmt).fmtQ
  1492: 000000000047afe0   148 FUNC    GLOBAL DEFAULT    1 fmt.(*fmt).fmtC
  1493: 000000000047b080   197 FUNC    GLOBAL DEFAULT    1 fmt.(*fmt).fmtQc
  1494: 000000000047b160  2105 FUNC    GLOBAL DEFAULT    1 fmt.(*fmt).fmtFloat
  1495: 000000000047b9a0    66 FUNC    GLOBAL DEFAULT    1 fmt.glob..func1
  1496: 000000000047ba00   146 FUNC    GLOBAL DEFAULT    1 fmt.newPrinter
  1497: 000000000047baa0   282 FUNC    GLOBAL DEFAULT    1 fmt.(*pp).free
  1498: 000000000047bbc0   266 FUNC    GLOBAL DEFAULT    1 fmt.(*pp).Write
  1499: 000000000047bce0   229 FUNC    GLOBAL DEFAULT    1 fmt.Fprintln
  1500: 000000000047bde0   135 FUNC    GLOBAL DEFAULT    1 fmt.getField
  1501: 000000000047be80   691 FUNC    GLOBAL DEFAULT    1 fmt.(*pp).unknownType
  1502: 000000000047c140  1363 FUNC    GLOBAL DEFAULT    1 fmt.(*pp).badVerb
  1503: 000000000047c6a0    94 FUNC    GLOBAL DEFAULT    1 fmt.(*pp).fmtBool
  1504: 000000000047c700   136 FUNC    GLOBAL DEFAULT    1 fmt.(*pp).fmt0x64
  1505: 000000000047c7a0   509 FUNC    GLOBAL DEFAULT    1 fmt.(*pp).fmtInteger
  1506: 000000000047c9a0   254 FUNC    GLOBAL DEFAULT    1 fmt.(*pp).fmtFloat
  1507: 000000000047caa0   485 FUNC    GLOBAL DEFAULT    1 fmt.(*pp).fmtComplex
  1508: 000000000047cca0   255 FUNC    GLOBAL DEFAULT    1 fmt.(*pp).fmtString
  1509: 000000000047cda0  1721 FUNC    GLOBAL DEFAULT    1 fmt.(*pp).fmtBytes
  1510: 000000000047d460  1113 FUNC    GLOBAL DEFAULT    1 fmt.(*pp).fmtPointer
  1511: 000000000047d8c0  1391 FUNC    GLOBAL DEFAULT    1 fmt.(*pp).catchPanic
  1512: 000000000047de40  1445 FUNC    GLOBAL DEFAULT    1 fmt.(*pp).handle[...]
  1513: 000000000047e400    98 FUNC    GLOBAL DEFAULT    1 fmt.(*pp).handle[...]
  1514: 000000000047e480    98 FUNC    GLOBAL DEFAULT    1 fmt.(*pp).handle[...]
  1515: 000000000047e500    98 FUNC    GLOBAL DEFAULT    1 fmt.(*pp).handle[...]
  1516: 000000000047e580    98 FUNC    GLOBAL DEFAULT    1 fmt.(*pp).handle[...]
  1517: 000000000047e600  1998 FUNC    GLOBAL DEFAULT    1 fmt.(*pp).printArg
  1518: 000000000047ede0  7959 FUNC    GLOBAL DEFAULT    1 fmt.(*pp).printValue
  1519: 0000000000480d00   404 FUNC    GLOBAL DEFAULT    1 fmt.(*pp).doPrintln
  1520: 0000000000480ea0   197 FUNC    GLOBAL DEFAULT    1 fmt.init
  1521: 0000000000480f80   197 FUNC    GLOBAL DEFAULT    1 type:.eq.fmt.fmt
  1522: 0000000000481060   103 FUNC    GLOBAL DEFAULT    1 main.main
  1523: 0000000000515000    32 OBJECT  GLOBAL DEFAULT    9 main..inittask
  1524: 0000000000516f40   104 OBJECT  GLOBAL DEFAULT    9 fmt..inittask
  1525: 0000000000525280    40 OBJECT  GLOBAL DEFAULT   10 fmt.ppFree
  1526: 000000000052c070    16 OBJECT  GLOBAL DEFAULT   11 fmt.complexError
  1527: 000000000052c060    16 OBJECT  GLOBAL DEFAULT   11 fmt.boolError
  1528: 0000000000515680    40 OBJECT  GLOBAL DEFAULT    9 errors..inittask
  1529: 000000000052c050    16 OBJECT  GLOBAL DEFAULT   11 errors.errorType
  1530: 0000000000515040    32 OBJECT  GLOBAL DEFAULT    9 sort..inittask
  1531: 0000000000515d80    56 OBJECT  GLOBAL DEFAULT    9 strconv..inittask
  1532: 0000000000514141     1 OBJECT  GLOBAL DEFAULT    9 strconv.optimize
  1533: 000000000052c2e0    16 OBJECT  GLOBAL DEFAULT   11 strconv.ErrRange
  1534: 000000000052c2f0    16 OBJECT  GLOBAL DEFAULT   11 strconv.ErrSyntax
  1535: 0000000000525230    24 OBJECT  GLOBAL DEFAULT   10 strconv.leftcheats
  1536: 0000000000521bc0 11136 OBJECT  GLOBAL DEFAULT    9 strconv.detailed[...]
  1537: 00000000005148a0    24 OBJECT  GLOBAL DEFAULT    9 strconv.float32info
  1538: 00000000005148c0    24 OBJECT  GLOBAL DEFAULT    9 strconv.float64info
  1539: 0000000000517660   160 OBJECT  GLOBAL DEFAULT    9 strconv.uint64pow10
  1540: 00000000005251f0    24 OBJECT  GLOBAL DEFAULT   10 strconv.isPrint16
  1541: 00000000005251b0    24 OBJECT  GLOBAL DEFAULT   10 strconv.isNotPrint16
  1542: 0000000000525210    24 OBJECT  GLOBAL DEFAULT   10 strconv.isPrint32
  1543: 00000000005251d0    24 OBJECT  GLOBAL DEFAULT   10 strconv.isNotPrint32
  1544: 0000000000525190    24 OBJECT  GLOBAL DEFAULT   10 strconv.isGraphic
  1545: 0000000000518000   256 OBJECT  GLOBAL DEFAULT    9 unicode/utf8.first
  1546: 00000000005150a0    32 OBJECT  GLOBAL DEFAULT    9 unicode/utf8.acc[...]
  1547: 00000000005156c0    40 OBJECT  GLOBAL DEFAULT    9 internal/fmtsort[...]
  1548: 0000000000515a40    48 OBJECT  GLOBAL DEFAULT    9 io..inittask
  1549: 000000000052c150    16 OBJECT  GLOBAL DEFAULT   11 io.ErrShortWrite
  1550: 000000000052c170    16 OBJECT  GLOBAL DEFAULT   11 io.errInvalidWrite
  1551: 000000000052c140    16 OBJECT  GLOBAL DEFAULT   11 io.ErrShortBuffer
  1552: 000000000052c110    16 OBJECT  GLOBAL DEFAULT   11 io.EOF
  1553: 000000000052c160    16 OBJECT  GLOBAL DEFAULT   11 io.ErrUnexpectedEOF
  1554: 000000000052c130    16 OBJECT  GLOBAL DEFAULT   11 io.ErrNoProgress
  1555: 000000000052c190    16 OBJECT  GLOBAL DEFAULT   11 io.errWhence
  1556: 000000000052c180    16 OBJECT  GLOBAL DEFAULT   11 io.errOffset
  1557: 000000000052c120    16 OBJECT  GLOBAL DEFAULT   11 io.ErrClosedPipe
  1558: 0000000000517520   152 OBJECT  GLOBAL DEFAULT    9 os..inittask
  1559: 00000000005252c0    40 OBJECT  GLOBAL DEFAULT   10 os.dirBufPool
  1560: 000000000052c240    16 OBJECT  GLOBAL DEFAULT   11 os.ErrInvalid
  1561: 000000000052c270    16 OBJECT  GLOBAL DEFAULT   11 os.ErrPermission
  1562: 000000000052c230    16 OBJECT  GLOBAL DEFAULT   11 os.ErrExist
  1563: 000000000052c260    16 OBJECT  GLOBAL DEFAULT   11 os.ErrNotExist
  1564: 000000000052c210    16 OBJECT  GLOBAL DEFAULT   11 os.ErrClosed
  1565: 000000000052c250    16 OBJECT  GLOBAL DEFAULT   11 os.ErrNoDeadline
  1566: 000000000052c220    16 OBJECT  GLOBAL DEFAULT   11 os.ErrDeadlineEx[...]
  1567: 000000000052c280    16 OBJECT  GLOBAL DEFAULT   11 os.ErrProcessDone
  1568: 000000000052bf28     8 OBJECT  GLOBAL DEFAULT   11 os.Stdin
  1569: 000000000052bf30     8 OBJECT  GLOBAL DEFAULT   11 os.Stdout
  1570: 000000000052bf20     8 OBJECT  GLOBAL DEFAULT   11 os.Stderr
  1571: 000000000052c2a0    16 OBJECT  GLOBAL DEFAULT   11 os.errWriteAtInA[...]
  1572: 000000000052c370    24 OBJECT  GLOBAL DEFAULT   11 os.Args
  1573: 000000000052c290    16 OBJECT  GLOBAL DEFAULT   11 os.errPatternHas[...]
  1574: 0000000000516bc0    88 OBJECT  GLOBAL DEFAULT    9 reflect..inittask
  1575: 0000000000514240     8 OBJECT  GLOBAL DEFAULT    9 reflect.intArgRegs
  1576: 0000000000514230     8 OBJECT  GLOBAL DEFAULT    9 reflect.floatArgRegs
  1577: 0000000000514238     8 OBJECT  GLOBAL DEFAULT    9 reflect.floatRegSize
  1578: 0000000000525070    24 OBJECT  GLOBAL DEFAULT   10 reflect.kindNames
  1579: 000000000052c580    32 OBJECT  GLOBAL DEFAULT   11 reflect.layoutCache
  1580: 000000000052bf38     8 OBJECT  GLOBAL DEFAULT   11 reflect.bytesType
  1581: 000000000052bf48     8 OBJECT  GLOBAL DEFAULT   11 reflect.uint8Type
  1582: 000000000052bf40     8 OBJECT  GLOBAL DEFAULT   11 reflect.stringType
  1583: 000000000052c390    24 OBJECT  GLOBAL DEFAULT   11 reflect.dummy
  1584: 0000000000515dc0    56 OBJECT  GLOBAL DEFAULT    9 sync..inittask
  1585: 000000000052c018     8 OBJECT  GLOBAL DEFAULT   11 sync.expunged
  1586: 000000000055a060     8 OBJECT  GLOBAL DEFAULT   12 sync.allPoolsMu
  1587: 000000000052c510    24 OBJECT  GLOBAL DEFAULT   11 sync.allPools
  1588: 000000000052c530    24 OBJECT  GLOBAL DEFAULT   11 sync.oldPools
  1589: 0000000000515020    32 OBJECT  GLOBAL DEFAULT    9 math..inittask
  1590: 0000000000559e61     1 OBJECT  GLOBAL DEFAULT   12 math.useFMA
  1591: 0000000000514f80    32 OBJECT  GLOBAL DEFAULT    9 internal/reflect[...]
  1592: 0000000000525050    24 OBJECT  GLOBAL DEFAULT   10 internal/reflect[...]
  1593: 0000000000514f60    32 OBJECT  GLOBAL DEFAULT    9 internal/bytealg[...]
  1594: 0000000000559ed8     8 OBJECT  GLOBAL DEFAULT   12 internal/bytealg[...]
  1595: 0000000000516740    80 OBJECT  GLOBAL DEFAULT    9 io/fs..inittask
  1596: 000000000052c1c0    16 OBJECT  GLOBAL DEFAULT   11 io/fs.ErrInvalid
  1597: 000000000052c1e0    16 OBJECT  GLOBAL DEFAULT   11 io/fs.ErrPermission
  1598: 000000000052c1b0    16 OBJECT  GLOBAL DEFAULT   11 io/fs.ErrExist
  1599: 000000000052c1d0    16 OBJECT  GLOBAL DEFAULT   11 io/fs.ErrNotExist
  1600: 000000000052c1a0    16 OBJECT  GLOBAL DEFAULT   11 io/fs.ErrClosed
  1601: 000000000052c200    16 OBJECT  GLOBAL DEFAULT   11 io/fs.SkipDir
  1602: 000000000052c1f0    16 OBJECT  GLOBAL DEFAULT   11 io/fs.SkipAll
  1603: 00000000005167a0    80 OBJECT  GLOBAL DEFAULT    9 runtime..inittask
  1604: 0000000000559e6f     1 OBJECT  GLOBAL DEFAULT   12 runtime.useAeshash
  1605: 000000000055a420   128 OBJECT  GLOBAL DEFAULT   12 runtime.aeskeysched
  1606: 000000000055a100    32 OBJECT  GLOBAL DEFAULT   12 runtime.hashkey
  1607: 000000000052c640    56 OBJECT  GLOBAL DEFAULT   11 runtime.userAren[...]
  1608: 0000000000559e6a     1 OBJECT  GLOBAL DEFAULT   12 runtime.iscgo
  1609: 0000000000559e64     1 OBJECT  GLOBAL DEFAULT   12 runtime.cgoHasExtraM
  1610: 0000000000524750     8 OBJECT  GLOBAL DEFAULT   10 runtime.cgo_yield
  1611: 0000000000559fa0     8 OBJECT  GLOBAL DEFAULT   12 runtime.ncgocall
  1612: 0000000000559e72     1 OBJECT  GLOBAL DEFAULT   12 runtime.x86HasPOPCNT
  1613: 0000000000559e73     1 OBJECT  GLOBAL DEFAULT   12 runtime.x86HasSSE41
  1614: 0000000000559e71     1 OBJECT  GLOBAL DEFAULT   12 runtime.x86HasFMA
  1615: 0000000000559e6e     1 OBJECT  GLOBAL DEFAULT   12 runtime.useAVXmemmove
  1616: 000000000052d520  8048 OBJECT  GLOBAL DEFAULT   11 runtime.cpuprof
  1617: 000000000052bf50     8 OBJECT  GLOBAL DEFAULT   11 runtime._cgo_setenv
  1618: 000000000052bf58     8 OBJECT  GLOBAL DEFAULT   11 runtime._cgo_unsetenv
  1619: 0000000000528f00   144 OBJECT  GLOBAL DEFAULT   10 runtime.boundsEr[...]
  1620: 0000000000528e80   128 OBJECT  GLOBAL DEFAULT   10 runtime.boundsNe[...]
  1621: 000000000052c5a0    32 OBJECT  GLOBAL DEFAULT   11 runtime.exitHooks
  1622: 0000000000524f50    16 OBJECT  GLOBAL DEFAULT   10 runtime.buildVersion
  1623: 0000000000518220   264 OBJECT  GLOBAL DEFAULT    9 runtime.fastlog2Table
  1624: 0000000000559f70     8 OBJECT  GLOBAL DEFAULT   12 runtime.inf
  1625: 0000000000559f80     8 OBJECT  GLOBAL DEFAULT   12 runtime.itabLock
  1626: 0000000000524758     8 OBJECT  GLOBAL DEFAULT   10 runtime.itabTable
  1627: 000000000052aec0  4112 OBJECT  GLOBAL DEFAULT   10 runtime.itabTableInit
  1628: 0000000000524ff0    16 OBJECT  GLOBAL DEFAULT   10 runtime.uint16Eface
  1629: 0000000000525000    16 OBJECT  GLOBAL DEFAULT   10 runtime.uint32Eface
  1630: 0000000000525010    16 OBJECT  GLOBAL DEFAULT   10 runtime.uint64Eface
  1631: 0000000000524fe0    16 OBJECT  GLOBAL DEFAULT   10 runtime.stringEface
  1632: 0000000000524fd0    16 OBJECT  GLOBAL DEFAULT   10 runtime.sliceEface
  1633: 000000000052bff8     8 OBJECT  GLOBAL DEFAULT   11 runtime.uint16Type
  1634: 000000000052c000     8 OBJECT  GLOBAL DEFAULT   11 runtime.uint32Type
  1635: 000000000052c008     8 OBJECT  GLOBAL DEFAULT   11 runtime.uint64Type
  1636: 000000000052bfe0     8 OBJECT  GLOBAL DEFAULT   11 runtime.stringType
  1637: 000000000052bfd8     8 OBJECT  GLOBAL DEFAULT   11 runtime.sliceType
  1638: 000000000051f7e0  2048 OBJECT  GLOBAL DEFAULT    9 runtime.staticuint64s
  1639: 00000000005250b0    24 OBJECT  GLOBAL DEFAULT   10 runtime.lockNames
  1640: 0000000000559fe0     8 OBJECT  GLOBAL DEFAULT   12 runtime.physPageSize
  1641: 0000000000559fd8     8 OBJECT  GLOBAL DEFAULT   12 runtime.physHuge[...]
  1642: 0000000000559fd0     8 OBJECT  GLOBAL DEFAULT   12 runtime.physHuge[...]
  1643: 000000000055a050     8 OBJECT  GLOBAL DEFAULT   12 runtime.zerobase
  1644: 000000000055a0c0    24 OBJECT  GLOBAL DEFAULT   12 runtime.globalAlloc
  1645: 0000000000559fc8     8 OBJECT  GLOBAL DEFAULT   12 runtime.persiste[...]
  1646: 000000000055b0e0  1024 OBJECT  GLOBAL DEFAULT   12 runtime.zeroVal
  1647: 000000000055a540   160 OBJECT  GLOBAL DEFAULT   12 runtime.emptymspan
  1648: 0000000000559e70     1 OBJECT  GLOBAL DEFAULT   12 runtime.useCheckmark
  1649: 0000000000514144     4 OBJECT  GLOBAL DEFAULT    9 runtime.adviseUnused
  1650: 0000000000559e90     4 OBJECT  GLOBAL DEFAULT   12 runtime.fingStatus
  1651: 0000000000559f58     8 OBJECT  GLOBAL DEFAULT   12 runtime.finlock
  1652: 000000000052bf98     8 OBJECT  GLOBAL DEFAULT   11 runtime.fing
  1653: 0000000000559f60     8 OBJECT  GLOBAL DEFAULT   12 runtime.finq
  1654: 0000000000559f50     8 OBJECT  GLOBAL DEFAULT   12 runtime.finc
  1655: 000000000055a260    64 OBJECT  GLOBAL DEFAULT   12 runtime.finptrmask
  1656: 0000000000559ef0     8 OBJECT  GLOBAL DEFAULT   12 runtime.allfin
  1657: 000000000051415c     5 OBJECT  GLOBAL DEFAULT    9 runtime.finalizer1
  1658: 0000000000559ea0     4 OBJECT  GLOBAL DEFAULT   12 runtime.gcphase
  1659: 000000000055a0b0    16 OBJECT  GLOBAL DEFAULT   12 runtime.writeBarrier
  1660: 0000000000559e98     4 OBJECT  GLOBAL DEFAULT   12 runtime.gcBlacke[...]
  1661: 000000000052c8c0   512 OBJECT  GLOBAL DEFAULT   11 runtime.work
  1662: 0000000000559e9c     4 OBJECT  GLOBAL DEFAULT   12 runtime.gcMarkDo[...]
  1663: 000000000052bfd0     8 OBJECT  GLOBAL DEFAULT   11 runtime.poolcleanup
  1664: 000000000052c430    24 OBJECT  GLOBAL DEFAULT   11 runtime.boringCaches
  1665: 000000000055a3c0    88 OBJECT  GLOBAL DEFAULT   12 runtime.gcCPULimiter
  1666: 0000000000514140     1 OBJECT  GLOBAL DEFAULT    9 runtime.oneptrmask
  1667: 000000000055a900   368 OBJECT  GLOBAL DEFAULT   12 runtime.gcController
  1668: 000000000055a140    32 OBJECT  GLOBAL DEFAULT   12 runtime.scavenge
  1669: 000000000052c680   160 OBJECT  GLOBAL DEFAULT   11 runtime.scavenger
  1670: 000000000052c600    40 OBJECT  GLOBAL DEFAULT   11 runtime.sweep
  1671: 0000000000543460 92672 OBJECT  GLOBAL DEFAULT   11 runtime.mheap_
  1672: 00000000005250d0    24 OBJECT  GLOBAL DEFAULT   10 runtime.mSpanSta[...]
  1673: 000000000055a160    40 OBJECT  GLOBAL DEFAULT   12 runtime.gcBitsArenas
  1674: 0000000000515740    40 OBJECT  GLOBAL DEFAULT    9 runtime.levelBits
  1675: 00000000005157c0    40 OBJECT  GLOBAL DEFAULT    9 runtime.levelShift
  1676: 0000000000515780    40 OBJECT  GLOBAL DEFAULT    9 runtime.levelLogPages
  1677: 000000000055a000     8 OBJECT  GLOBAL DEFAULT   12 runtime.profInse[...]
  1678: 0000000000559ff8     8 OBJECT  GLOBAL DEFAULT   12 runtime.profBlockLock
  1679: 000000000055a008     8 OBJECT  GLOBAL DEFAULT   12 runtime.profMemA[...]
  1680: 000000000055a0e0    24 OBJECT  GLOBAL DEFAULT   12 runtime.profMemF[...]
  1681: 000000000052bfb0     8 OBJECT  GLOBAL DEFAULT   11 runtime.mbuckets
  1682: 000000000052bf78     8 OBJECT  GLOBAL DEFAULT   11 runtime.bbuckets
  1683: 000000000052c010     8 OBJECT  GLOBAL DEFAULT   11 runtime.xbuckets
  1684: 000000000052bf80     8 OBJECT  GLOBAL DEFAULT   11 runtime.buckhash
  1685: 0000000000559ea8     4 OBJECT  GLOBAL DEFAULT   12 runtime.mProfCycle
  1686: 0000000000559f10     8 OBJECT  GLOBAL DEFAULT   12 runtime.blockpro[...]
  1687: 0000000000559f98     8 OBJECT  GLOBAL DEFAULT   12 runtime.mutexpro[...]
  1688: 0000000000514248     8 OBJECT  GLOBAL DEFAULT    9 runtime.MemProfi[...]
  1689: 0000000000524740     1 OBJECT  GLOBAL DEFAULT   10 runtime.disableM[...]
  1690: 0000000000528e40    64 OBJECT  GLOBAL DEFAULT   10 runtime.goroutin[...]
  1691: 000000000055a038     8 OBJECT  GLOBAL DEFAULT   12 runtime.tracelock
  1692: 0000000000514288     8 OBJECT  GLOBAL DEFAULT    9 runtime.minOffAddr
  1693: 0000000000514270     8 OBJECT  GLOBAL DEFAULT    9 runtime.maxOffAddr
  1694: 000000000055a018     8 OBJECT  GLOBAL DEFAULT   12 runtime.spanSetB[...]
  1695: 000000000055b4e0  9040 OBJECT  GLOBAL DEFAULT   12 runtime.memstats
  1696: 0000000000559fb8     8 OBJECT  GLOBAL DEFAULT   12 runtime.netpollI[...]
  1697: 0000000000559eb0     4 OBJECT  GLOBAL DEFAULT   12 runtime.netpollInited
  1698: 000000000055a0a0    16 OBJECT  GLOBAL DEFAULT   12 runtime.pollcache
  1699: 0000000000559eb4     4 OBJECT  GLOBAL DEFAULT   12 runtime.netpollW[...]
  1700: 0000000000524fb0    16 OBJECT  GLOBAL DEFAULT   10 runtime.pdEface
  1701: 000000000052bfc8     8 OBJECT  GLOBAL DEFAULT   11 runtime.pdType
  1702: 0000000000514148     4 OBJECT  GLOBAL DEFAULT    9 runtime.epfd
  1703: 0000000000559fa8     8 OBJECT  GLOBAL DEFAULT   12 runtime.netpollB[...]
  1704: 0000000000559fb0     8 OBJECT  GLOBAL DEFAULT   12 runtime.netpollB[...]
  1705: 0000000000559eb8     4 OBJECT  GLOBAL DEFAULT   12 runtime.netpollW[...]
  1706: 00000000005250f0    24 OBJECT  GLOBAL DEFAULT   10 runtime.procAuxv
  1707: 0000000000559e62     1 OBJECT  GLOBAL DEFAULT   12 runtime.addrspace_vec
  1708: 000000000052c4d0    24 OBJECT  GLOBAL DEFAULT   11 runtime.startupR[...]
  1709: 0000000000525110    24 OBJECT  GLOBAL DEFAULT   10 runtime.sysTHPSi[...]
  1710: 0000000000525130    24 OBJECT  GLOBAL DEFAULT   10 runtime.urandom_dev
  1711: 000000000055a2a0    72 OBJECT  GLOBAL DEFAULT   12 runtime.perThrea[...]
  1712: 0000000000514298     8 OBJECT  GLOBAL DEFAULT    9 runtime.sigset_all
  1713: 0000000000524fc0    16 OBJECT  GLOBAL DEFAULT   10 runtime.shiftError
  1714: 0000000000524f60    16 OBJECT  GLOBAL DEFAULT   10 runtime.divideError
  1715: 0000000000524fa0    16 OBJECT  GLOBAL DEFAULT   10 runtime.overflowError
  1716: 0000000000524f70    16 OBJECT  GLOBAL DEFAULT   10 runtime.floatError
  1717: 0000000000524f80    16 OBJECT  GLOBAL DEFAULT   10 runtime.memoryError
  1718: 0000000000559ec8     4 OBJECT  GLOBAL DEFAULT   12 runtime.runningP[...]
  1719: 0000000000559ec0     4 OBJECT  GLOBAL DEFAULT   12 runtime.panicking
  1720: 0000000000559fc0     8 OBJECT  GLOBAL DEFAULT   12 runtime.paniclk
  1721: 0000000000559e65     1 OBJECT  GLOBAL DEFAULT   12 runtime.didothers
  1722: 0000000000559f28     8 OBJECT  GLOBAL DEFAULT   12 runtime.deadlock
  1723: 0000000000514250     8 OBJECT  GLOBAL DEFAULT    9 runtime.asyncPre[...]
  1724: 000000000055aa80   512 OBJECT  GLOBAL DEFAULT   12 runtime.printBacklog
  1725: 0000000000559fe8     8 OBJECT  GLOBAL DEFAULT   12 runtime.printBac[...]
  1726: 0000000000559f30     8 OBJECT  GLOBAL DEFAULT   12 runtime.debuglock
  1727: 0000000000559f90     8 OBJECT  GLOBAL DEFAULT   12 runtime.minhexdigits
  1728: 0000000000524f90    16 OBJECT  GLOBAL DEFAULT   10 runtime.modinfo
  1729: 000000000052cac0  1000 OBJECT  GLOBAL DEFAULT   11 runtime.m0
  1730: 000000000052c720   392 OBJECT  GLOBAL DEFAULT   11 runtime.g0
  1731: 0000000000559f88     8 OBJECT  GLOBAL DEFAULT   12 runtime.mcache0
  1732: 000000000052bfa8     8 OBJECT  GLOBAL DEFAULT   11 runtime.main_ini[...]
  1733: 0000000000559e6c     1 OBJECT  GLOBAL DEFAULT   12 runtime.mainStarted
  1734: 000000000055a010     8 OBJECT  GLOBAL DEFAULT   12 runtime.runtimeI[...]
  1735: 0000000000559f78     8 OBJECT  GLOBAL DEFAULT   12 runtime.initSigmask
  1736: 0000000000559f00     8 OBJECT  GLOBAL DEFAULT   12 runtime.allglock
  1737: 000000000052c3d0    24 OBJECT  GLOBAL DEFAULT   11 runtime.allgs
  1738: 0000000000559ef8     8 OBJECT  GLOBAL DEFAULT   12 runtime.allglen
  1739: 000000000052bf60     8 OBJECT  GLOBAL DEFAULT   11 runtime.allgptr
  1740: 0000000000559f48     8 OBJECT  GLOBAL DEFAULT   12 runtime.fastrandseed
  1741: 0000000000559e66     1 OBJECT  GLOBAL DEFAULT   12 runtime.freezing
  1742: 0000000000559e63     1 OBJECT  GLOBAL DEFAULT   12 runtime.casgstat[...]
  1743: 0000000000514158     4 OBJECT  GLOBAL DEFAULT    9 runtime.worldsema
  1744: 000000000051414c     4 OBJECT  GLOBAL DEFAULT    9 runtime.gcsema
  1745: 0000000000559f38     8 OBJECT  GLOBAL DEFAULT   12 runtime.extram
  1746: 0000000000559e88     4 OBJECT  GLOBAL DEFAULT   12 runtime.extraMCount
  1747: 0000000000559e8c     4 OBJECT  GLOBAL DEFAULT   12 runtime.extraMWaiters
  1748: 000000000055a1e0    48 OBJECT  GLOBAL DEFAULT   12 runtime.allocmLock
  1749: 000000000055a220    48 OBJECT  GLOBAL DEFAULT   12 runtime.execLock
  1750: 000000000055a1a0    40 OBJECT  GLOBAL DEFAULT   12 runtime.newmHandoff
  1751: 0000000000559e67     1 OBJECT  GLOBAL DEFAULT   12 runtime.inForkedChild
  1752: 0000000000559ff0     8 OBJECT  GLOBAL DEFAULT   12 runtime.prof
  1753: 0000000000514258     8 OBJECT  GLOBAL DEFAULT    9 runtime.forcegcperiod
  1754: 000000000055a020     8 OBJECT  GLOBAL DEFAULT   12 runtime.starttime
  1755: 000000000052c5e0    32 OBJECT  GLOBAL DEFAULT   11 runtime.stealOrder
  1756: 000000000055a120    32 OBJECT  GLOBAL DEFAULT   12 runtime.inittrace
  1757: 000000000052c450    24 OBJECT  GLOBAL DEFAULT   11 runtime.envs
  1758: 000000000052c410    24 OBJECT  GLOBAL DEFAULT   11 runtime.argslice
  1759: 000000000052bfa0     8 OBJECT  GLOBAL DEFAULT   11 runtime.godebugEnv
  1760: 0000000000514154     4 OBJECT  GLOBAL DEFAULT    9 runtime.tracebac[...]
  1761: 0000000000559ed0     4 OBJECT  GLOBAL DEFAULT   12 runtime.traceback_env
  1762: 0000000000559e80     4 OBJECT  GLOBAL DEFAULT   12 runtime.argc
  1763: 000000000052bf70     8 OBJECT  GLOBAL DEFAULT   11 runtime.argv
  1764: 000000000055a030     8 OBJECT  GLOBAL DEFAULT   12 runtime.test_z64
  1765: 000000000055a028     8 OBJECT  GLOBAL DEFAULT   12 runtime.test_x64
  1766: 000000000055a360    84 OBJECT  GLOBAL DEFAULT   12 runtime.debug
  1767: 0000000000525090    24 OBJECT  GLOBAL DEFAULT   10 runtime.dbgvars
  1768: 000000000052c2d0    16 OBJECT  GLOBAL DEFAULT   11 runtime.globalGODEBUG
  1769: 00000000005295a0   512 OBJECT  GLOBAL DEFAULT   10 runtime.waitReas[...]
  1770: 000000000052bf68     8 OBJECT  GLOBAL DEFAULT   11 runtime.allm
  1771: 0000000000559ea4     4 OBJECT  GLOBAL DEFAULT   12 runtime.gomaxprocs
  1772: 0000000000559eac     4 OBJECT  GLOBAL DEFAULT   12 runtime.ncpu
  1773: 000000000052c470    24 OBJECT  GLOBAL DEFAULT   11 runtime.forcegc
  1774: 000000000052cec0  1632 OBJECT  GLOBAL DEFAULT   11 runtime.sched
  1775: 0000000000559ebc     4 OBJECT  GLOBAL DEFAULT   12 runtime.newprocs
  1776: 0000000000559f08     8 OBJECT  GLOBAL DEFAULT   12 runtime.allpLock
  1777: 000000000052c3f0    24 OBJECT  GLOBAL DEFAULT   11 runtime.allp
  1778: 000000000052c490    24 OBJECT  GLOBAL DEFAULT   11 runtime.idlepMask
  1779: 000000000052c4f0    24 OBJECT  GLOBAL DEFAULT   11 runtime.timerpMask
  1780: 0000000000559f68     8 OBJECT  GLOBAL DEFAULT   12 runtime.gcBgMark[...]
  1781: 0000000000559e94     4 OBJECT  GLOBAL DEFAULT   12 runtime.gcBgMark[...]
  1782: 0000000000559ec4     4 OBJECT  GLOBAL DEFAULT   12 runtime.processo[...]
  1783: 0000000000559e68     1 OBJECT  GLOBAL DEFAULT   12 runtime.isIntel
  1784: 0000000000559e6b     1 OBJECT  GLOBAL DEFAULT   12 runtime.islibrary
  1785: 0000000000559e69     1 OBJECT  GLOBAL DEFAULT   12 runtime.isarchive
  1786: 0000000000559f20     8 OBJECT  GLOBAL DEFAULT   12 runtime.chansendpc
  1787: 0000000000559f18     8 OBJECT  GLOBAL DEFAULT   12 runtime.chanrecvpc
  1788: 000000000052f4a0 16064 OBJECT  GLOBAL DEFAULT   11 runtime.semtable
  1789: 000000000055ac80   520 OBJECT  GLOBAL DEFAULT   12 runtime.fwdSig
  1790: 000000000055a7e0   260 OBJECT  GLOBAL DEFAULT   12 runtime.handlingSig
  1791: 0000000000559e6d     1 OBJECT  GLOBAL DEFAULT   12 runtime.signalsOK
  1792: 000000000055a5e0   256 OBJECT  GLOBAL DEFAULT   12 runtime.sigprofC[...]
  1793: 0000000000559ecc     4 OBJECT  GLOBAL DEFAULT   12 runtime.sigprofC[...]
  1794: 0000000000559e84     4 OBJECT  GLOBAL DEFAULT   12 runtime.crashing
  1795: 000000000052bfe8     8 OBJECT  GLOBAL DEFAULT   11 runtime.testSigtrap
  1796: 000000000052bff0     8 OBJECT  GLOBAL DEFAULT   11 runtime.testSigusr1
  1797: 0000000000514290     8 OBJECT  GLOBAL DEFAULT    9 runtime.sigsetAl[...]
  1798: 000000000055a300    72 OBJECT  GLOBAL DEFAULT   12 runtime.sig
  1799: 000000000052a040  1560 OBJECT  GLOBAL DEFAULT   10 runtime.sigtable
  1800: 0000000000517200   136 OBJECT  GLOBAL DEFAULT    9 runtime.class_to_size
  1801: 00000000005161a0    68 OBJECT  GLOBAL DEFAULT    9 runtime.class_to[...]
  1802: 0000000000518460   272 OBJECT  GLOBAL DEFAULT    9 runtime.class_to[...]
  1803: 00000000005170c0   129 OBJECT  GLOBAL DEFAULT    9 runtime.size_to_[...]
  1804: 0000000000517f00   249 OBJECT  GLOBAL DEFAULT    9 runtime.size_to_[...]
  1805: 000000000055a6e0   256 OBJECT  GLOBAL DEFAULT   12 runtime.stackpool
  1806: 000000000055aea0   568 OBJECT  GLOBAL DEFAULT   12 runtime.stackLarge
  1807: 0000000000514280     8 OBJECT  GLOBAL DEFAULT    9 runtime.maxstacksize
  1808: 0000000000514278     8 OBJECT  GLOBAL DEFAULT    9 runtime.maxstack[...]
  1809: 0000000000514150     4 OBJECT  GLOBAL DEFAULT    9 runtime.starting[...]
  1810: 000000000055a090    16 OBJECT  GLOBAL DEFAULT   12 runtime.methodVa[...]
  1811: 0000000000514260     8 OBJECT  GLOBAL DEFAULT    9 runtime.intArgRegs
  1812: 000000000052c4b0    24 OBJECT  GLOBAL DEFAULT   11 runtime.pinnedTy[...]
  1813: 0000000000519d00   568 OBJECT  GLOBAL DEFAULT    9 runtime.firstmod[...]
  1814: 0000000000514268     8 OBJECT  GLOBAL DEFAULT    9 runtime.lastmodu[...]
  1815: 000000000052bfb8     8 OBJECT  GLOBAL DEFAULT   11 runtime.modulesSlice
  1816: 0000000000559f40     8 OBJECT  GLOBAL DEFAULT   12 runtime.faketime
  1817: 000000000052bfc0     8 OBJECT  GLOBAL DEFAULT   11 runtime.overrideWrite
  1818: 0000000000533360 65784 OBJECT  GLOBAL DEFAULT   11 runtime.trace
  1819: 0000000000528fa0   160 OBJECT  GLOBAL DEFAULT   10 runtime.gStatusS[...]
  1820: 000000000052bf90     8 OBJECT  GLOBAL DEFAULT   11 runtime.cgoTraceback
  1821: 000000000052bf88     8 OBJECT  GLOBAL DEFAULT   11 runtime.cgoSymbolizer
  1822: 000000000052c5c0    32 OBJECT  GLOBAL DEFAULT   11 runtime.reflectOffs
  1823: 0000000000525150    24 OBJECT  GLOBAL DEFAULT   10 runtime.vdsoLinu[...]
  1824: 0000000000525170    24 OBJECT  GLOBAL DEFAULT   10 runtime.vdsoSymb[...]
  1825: 000000000055a048     8 OBJECT  GLOBAL DEFAULT   12 runtime.vdsoGett[...]
  1826: 000000000055a040     8 OBJECT  GLOBAL DEFAULT   12 runtime.vdsoCloc[...]
  1827: 0000000000516200    72 OBJECT  GLOBAL DEFAULT    9 syscall..inittask
  1828: 000000000052c550    24 OBJECT  GLOBAL DEFAULT   11 syscall.envs
  1829: 000000000055a070     8 OBJECT  GLOBAL DEFAULT   12 syscall._zero
  1830: 000000000055a068     8 OBJECT  GLOBAL DEFAULT   12 syscall.Stdin
  1831: 00000000005142c0     8 OBJECT  GLOBAL DEFAULT    9 syscall.Stdout
  1832: 00000000005142b8     8 OBJECT  GLOBAL DEFAULT    9 syscall.Stderr
  1833: 0000000000525020    16 OBJECT  GLOBAL DEFAULT   10 syscall.errEAGAIN
  1834: 0000000000525030    16 OBJECT  GLOBAL DEFAULT   10 syscall.errEINVAL
  1835: 0000000000525040    16 OBJECT  GLOBAL DEFAULT   10 syscall.errENOENT
  1836: 000000000052a660  2128 OBJECT  GLOBAL DEFAULT   10 syscall.errors
  1837: 0000000000514fe0    32 OBJECT  GLOBAL DEFAULT    9 internal/testlog[...]
  1838: 0000000000516b60    88 OBJECT  GLOBAL DEFAULT    9 internal/poll..i[...]
  1839: 0000000000524ed0    16 OBJECT  GLOBAL DEFAULT   10 internal/poll.er[...]
  1840: 0000000000524ee0    16 OBJECT  GLOBAL DEFAULT   10 internal/poll.er[...]
  1841: 0000000000524ef0    16 OBJECT  GLOBAL DEFAULT   10 internal/poll.er[...]
  1842: 000000000052c0d0    16 OBJECT  GLOBAL DEFAULT   11 internal/poll.Er[...]
  1843: 000000000052c0e0    16 OBJECT  GLOBAL DEFAULT   11 internal/poll.Er[...]
  1844: 0000000000524ec0    16 OBJECT  GLOBAL DEFAULT   10 internal/poll.Er[...]
  1845: 000000000052c0f0    16 OBJECT  GLOBAL DEFAULT   11 internal/poll.Er[...]
  1846: 000000000055a080    12 OBJECT  GLOBAL DEFAULT   12 internal/poll.se[...]
  1847: 0000000000524748     8 OBJECT  GLOBAL DEFAULT   10 internal/poll.Cl[...]
  1848: 0000000000516040    64 OBJECT  GLOBAL DEFAULT    9 time..inittask
  1849: 000000000052c300    16 OBJECT  GLOBAL DEFAULT   11 time.atoiError
  1850: 000000000052c320    16 OBJECT  GLOBAL DEFAULT   11 time.errBad
  1851: 000000000052c330    16 OBJECT  GLOBAL DEFAULT   11 time.errLeadingInt
  1852: 000000000052c020     8 OBJECT  GLOBAL DEFAULT   11 time.unitMap
  1853: 000000000055a078     8 OBJECT  GLOBAL DEFAULT   12 time.startNano
  1854: 000000000052c340    16 OBJECT  GLOBAL DEFAULT   11 time.errLocation
  1855: 000000000052c310    16 OBJECT  GLOBAL DEFAULT   11 time.badData
  1856: 0000000000514fa0    32 OBJECT  GLOBAL DEFAULT    9 internal/syscall[...]
  1857: 0000000000515a00    48 OBJECT  GLOBAL DEFAULT    9 internal/safefil[...]
  1858: 000000000052c100    16 OBJECT  GLOBAL DEFAULT   11 internal/safefil[...]
  1859: 0000000000514fc0    32 OBJECT  GLOBAL DEFAULT    9 internal/syscall[...]
  1860: 0000000000514228     8 OBJECT  GLOBAL DEFAULT    9 internal/syscall[...]
  1861: 0000000000515080    32 OBJECT  GLOBAL DEFAULT    9 unicode..inittask
  1862: 000000000052c028     8 OBJECT  GLOBAL DEFAULT   11 unicode.Categories
  1863: 0000000000524818     8 OBJECT  GLOBAL DEFAULT   10 unicode.Cc
  1864: 0000000000524820     8 OBJECT  GLOBAL DEFAULT   10 unicode.Cf
  1865: 0000000000524848     8 OBJECT  GLOBAL DEFAULT   10 unicode.Co
  1866: 0000000000524860     8 OBJECT  GLOBAL DEFAULT   10 unicode.Cs
  1867: 0000000000524b48     8 OBJECT  GLOBAL DEFAULT   10 unicode.Nd
  1868: 0000000000524a00     8 OBJECT  GLOBAL DEFAULT   10 unicode.L
  1869: 0000000000524a48     8 OBJECT  GLOBAL DEFAULT   10 unicode.Lm
  1870: 0000000000524a50     8 OBJECT  GLOBAL DEFAULT   10 unicode.Lo
  1871: 0000000000524a40     8 OBJECT  GLOBAL DEFAULT   10 unicode.Ll
  1872: 0000000000524a80     8 OBJECT  GLOBAL DEFAULT   10 unicode.M
  1873: 0000000000524ac0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Mc
  1874: 0000000000524ac8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Me
  1875: 0000000000524b00     8 OBJECT  GLOBAL DEFAULT   10 unicode.Mn
  1876: 0000000000524b68     8 OBJECT  GLOBAL DEFAULT   10 unicode.Nl
  1877: 0000000000524b70     8 OBJECT  GLOBAL DEFAULT   10 unicode.No
  1878: 0000000000524b30     8 OBJECT  GLOBAL DEFAULT   10 unicode.N
  1879: 00000000005247f8     8 OBJECT  GLOBAL DEFAULT   10 unicode.C
  1880: 0000000000524c68     8 OBJECT  GLOBAL DEFAULT   10 unicode.Pc
  1881: 0000000000524c70     8 OBJECT  GLOBAL DEFAULT   10 unicode.Pd
  1882: 0000000000524c78     8 OBJECT  GLOBAL DEFAULT   10 unicode.Pe
  1883: 0000000000524c80     8 OBJECT  GLOBAL DEFAULT   10 unicode.Pf
  1884: 0000000000524c98     8 OBJECT  GLOBAL DEFAULT   10 unicode.Pi
  1885: 0000000000524ca0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Po
  1886: 0000000000524cb0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Ps
  1887: 0000000000524c38     8 OBJECT  GLOBAL DEFAULT   10 unicode.P
  1888: 0000000000524d00     8 OBJECT  GLOBAL DEFAULT   10 unicode.Sc
  1889: 0000000000524d38     8 OBJECT  GLOBAL DEFAULT   10 unicode.Sk
  1890: 0000000000524d40     8 OBJECT  GLOBAL DEFAULT   10 unicode.Sm
  1891: 0000000000524d48     8 OBJECT  GLOBAL DEFAULT   10 unicode.So
  1892: 0000000000524e48     8 OBJECT  GLOBAL DEFAULT   10 unicode.Z
  1893: 0000000000524ce8     8 OBJECT  GLOBAL DEFAULT   10 unicode.S
  1894: 0000000000524a60     8 OBJECT  GLOBAL DEFAULT   10 unicode.Lt
  1895: 0000000000524a68     8 OBJECT  GLOBAL DEFAULT   10 unicode.Lu
  1896: 0000000000524e58     8 OBJECT  GLOBAL DEFAULT   10 unicode.Zl
  1897: 0000000000524e60     8 OBJECT  GLOBAL DEFAULT   10 unicode.Zp
  1898: 0000000000524e68     8 OBJECT  GLOBAL DEFAULT   10 unicode.Zs
  1899: 000000000052c048     8 OBJECT  GLOBAL DEFAULT   11 unicode.Scripts
  1900: 0000000000524768     8 OBJECT  GLOBAL DEFAULT   10 unicode.Adlam
  1901: 0000000000524770     8 OBJECT  GLOBAL DEFAULT   10 unicode.Ahom
  1902: 0000000000524778     8 OBJECT  GLOBAL DEFAULT   10 unicode.Anatolia[...]
  1903: 0000000000524780     8 OBJECT  GLOBAL DEFAULT   10 unicode.Arabic
  1904: 0000000000524788     8 OBJECT  GLOBAL DEFAULT   10 unicode.Armenian
  1905: 0000000000524790     8 OBJECT  GLOBAL DEFAULT   10 unicode.Avestan
  1906: 0000000000524798     8 OBJECT  GLOBAL DEFAULT   10 unicode.Balinese
  1907: 00000000005247a0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Bamum
  1908: 00000000005247a8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Bassa_Vah
  1909: 00000000005247b0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Batak
  1910: 00000000005247b8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Bengali
  1911: 00000000005247c0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Bhaiksuki
  1912: 00000000005247d0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Bopomofo
  1913: 00000000005247d8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Brahmi
  1914: 00000000005247e0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Braille
  1915: 00000000005247e8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Buginese
  1916: 00000000005247f0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Buhid
  1917: 0000000000524800     8 OBJECT  GLOBAL DEFAULT   10 unicode.Canadian[...]
  1918: 0000000000524808     8 OBJECT  GLOBAL DEFAULT   10 unicode.Carian
  1919: 0000000000524810     8 OBJECT  GLOBAL DEFAULT   10 unicode.Caucasia[...]
  1920: 0000000000524828     8 OBJECT  GLOBAL DEFAULT   10 unicode.Chakma
  1921: 0000000000524830     8 OBJECT  GLOBAL DEFAULT   10 unicode.Cham
  1922: 0000000000524838     8 OBJECT  GLOBAL DEFAULT   10 unicode.Cherokee
  1923: 0000000000524840     8 OBJECT  GLOBAL DEFAULT   10 unicode.Chorasmian
  1924: 0000000000524850     8 OBJECT  GLOBAL DEFAULT   10 unicode.Common
  1925: 0000000000524858     8 OBJECT  GLOBAL DEFAULT   10 unicode.Coptic
  1926: 0000000000524868     8 OBJECT  GLOBAL DEFAULT   10 unicode.Cuneiform
  1927: 0000000000524870     8 OBJECT  GLOBAL DEFAULT   10 unicode.Cypriot
  1928: 0000000000524878     8 OBJECT  GLOBAL DEFAULT   10 unicode.Cyrillic
  1929: 0000000000524890     8 OBJECT  GLOBAL DEFAULT   10 unicode.Deseret
  1930: 0000000000524898     8 OBJECT  GLOBAL DEFAULT   10 unicode.Devanagari
  1931: 00000000005248a8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Dives_Akuru
  1932: 00000000005248b0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Dogra
  1933: 00000000005248b8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Duployan
  1934: 00000000005248c0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Egyptian[...]
  1935: 00000000005248c8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Elbasan
  1936: 00000000005248d0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Elymaic
  1937: 00000000005248d8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Ethiopic
  1938: 00000000005248e8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Georgian
  1939: 00000000005248f0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Glagolitic
  1940: 00000000005248f8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Gothic
  1941: 0000000000524900     8 OBJECT  GLOBAL DEFAULT   10 unicode.Grantha
  1942: 0000000000524908     8 OBJECT  GLOBAL DEFAULT   10 unicode.Greek
  1943: 0000000000524910     8 OBJECT  GLOBAL DEFAULT   10 unicode.Gujarati
  1944: 0000000000524918     8 OBJECT  GLOBAL DEFAULT   10 unicode.Gunjala_Gondi
  1945: 0000000000524920     8 OBJECT  GLOBAL DEFAULT   10 unicode.Gurmukhi
  1946: 0000000000524928     8 OBJECT  GLOBAL DEFAULT   10 unicode.Han
  1947: 0000000000524930     8 OBJECT  GLOBAL DEFAULT   10 unicode.Hangul
  1948: 0000000000524938     8 OBJECT  GLOBAL DEFAULT   10 unicode.Hanifi_R[...]
  1949: 0000000000524940     8 OBJECT  GLOBAL DEFAULT   10 unicode.Hanunoo
  1950: 0000000000524948     8 OBJECT  GLOBAL DEFAULT   10 unicode.Hatran
  1951: 0000000000524950     8 OBJECT  GLOBAL DEFAULT   10 unicode.Hebrew
  1952: 0000000000524960     8 OBJECT  GLOBAL DEFAULT   10 unicode.Hiragana
  1953: 0000000000524988     8 OBJECT  GLOBAL DEFAULT   10 unicode.Imperial[...]
  1954: 0000000000524990     8 OBJECT  GLOBAL DEFAULT   10 unicode.Inherited
  1955: 0000000000524998     8 OBJECT  GLOBAL DEFAULT   10 unicode.Inscript[...]
  1956: 00000000005249a0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Inscript[...]
  1957: 00000000005249a8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Javanese
  1958: 00000000005249b8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Kaithi
  1959: 00000000005249c0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Kannada
  1960: 00000000005249c8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Katakana
  1961: 00000000005249d0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Kayah_Li
  1962: 00000000005249d8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Kharoshthi
  1963: 00000000005249e0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Khitan_S[...]
  1964: 00000000005249e8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Khmer
  1965: 00000000005249f0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Khojki
  1966: 00000000005249f8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Khudawadi
  1967: 0000000000524a08     8 OBJECT  GLOBAL DEFAULT   10 unicode.Lao
  1968: 0000000000524a10     8 OBJECT  GLOBAL DEFAULT   10 unicode.Latin
  1969: 0000000000524a18     8 OBJECT  GLOBAL DEFAULT   10 unicode.Lepcha
  1970: 0000000000524a20     8 OBJECT  GLOBAL DEFAULT   10 unicode.Limbu
  1971: 0000000000524a28     8 OBJECT  GLOBAL DEFAULT   10 unicode.Linear_A
  1972: 0000000000524a30     8 OBJECT  GLOBAL DEFAULT   10 unicode.Linear_B
  1973: 0000000000524a38     8 OBJECT  GLOBAL DEFAULT   10 unicode.Lisu
  1974: 0000000000524a70     8 OBJECT  GLOBAL DEFAULT   10 unicode.Lycian
  1975: 0000000000524a78     8 OBJECT  GLOBAL DEFAULT   10 unicode.Lydian
  1976: 0000000000524a88     8 OBJECT  GLOBAL DEFAULT   10 unicode.Mahajani
  1977: 0000000000524a90     8 OBJECT  GLOBAL DEFAULT   10 unicode.Makasar
  1978: 0000000000524a98     8 OBJECT  GLOBAL DEFAULT   10 unicode.Malayalam
  1979: 0000000000524aa0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Mandaic
  1980: 0000000000524aa8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Manichaean
  1981: 0000000000524ab0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Marchen
  1982: 0000000000524ab8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Masaram_Gondi
  1983: 0000000000524ad0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Medefaidrin
  1984: 0000000000524ad8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Meetei_Mayek
  1985: 0000000000524ae0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Mende_Kikakui
  1986: 0000000000524ae8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Meroitic[...]
  1987: 0000000000524af0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Meroitic[...]
  1988: 0000000000524af8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Miao
  1989: 0000000000524b08     8 OBJECT  GLOBAL DEFAULT   10 unicode.Modi
  1990: 0000000000524b10     8 OBJECT  GLOBAL DEFAULT   10 unicode.Mongolian
  1991: 0000000000524b18     8 OBJECT  GLOBAL DEFAULT   10 unicode.Mro
  1992: 0000000000524b20     8 OBJECT  GLOBAL DEFAULT   10 unicode.Multani
  1993: 0000000000524b28     8 OBJECT  GLOBAL DEFAULT   10 unicode.Myanmar
  1994: 0000000000524b38     8 OBJECT  GLOBAL DEFAULT   10 unicode.Nabataean
  1995: 0000000000524b40     8 OBJECT  GLOBAL DEFAULT   10 unicode.Nandinagari
  1996: 0000000000524b50     8 OBJECT  GLOBAL DEFAULT   10 unicode.New_Tai_Lue
  1997: 0000000000524b58     8 OBJECT  GLOBAL DEFAULT   10 unicode.Newa
  1998: 0000000000524b60     8 OBJECT  GLOBAL DEFAULT   10 unicode.Nko
  1999: 0000000000524b80     8 OBJECT  GLOBAL DEFAULT   10 unicode.Nushu
  2000: 0000000000524b88     8 OBJECT  GLOBAL DEFAULT   10 unicode.Nyiakeng[...]
  2001: 0000000000524b90     8 OBJECT  GLOBAL DEFAULT   10 unicode.Ogham
  2002: 0000000000524b98     8 OBJECT  GLOBAL DEFAULT   10 unicode.Ol_Chiki
  2003: 0000000000524ba0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Old_Hungarian
  2004: 0000000000524ba8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Old_Italic
  2005: 0000000000524bb0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Old_Nort[...]
  2006: 0000000000524bb8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Old_Permic
  2007: 0000000000524bc0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Old_Persian
  2008: 0000000000524bc8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Old_Sogdian
  2009: 0000000000524bd0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Old_Sout[...]
  2010: 0000000000524bd8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Old_Turkic
  2011: 0000000000524be0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Oriya
  2012: 0000000000524be8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Osage
  2013: 0000000000524bf0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Osmanya
  2014: 0000000000524c40     8 OBJECT  GLOBAL DEFAULT   10 unicode.Pahawh_Hmong
  2015: 0000000000524c48     8 OBJECT  GLOBAL DEFAULT   10 unicode.Palmyrene
  2016: 0000000000524c60     8 OBJECT  GLOBAL DEFAULT   10 unicode.Pau_Cin_Hau
  2017: 0000000000524c88     8 OBJECT  GLOBAL DEFAULT   10 unicode.Phags_Pa
  2018: 0000000000524c90     8 OBJECT  GLOBAL DEFAULT   10 unicode.Phoenician
  2019: 0000000000524cb8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Psalter_[...]
  2020: 0000000000524cd8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Rejang
  2021: 0000000000524ce0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Runic
  2022: 0000000000524cf0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Samaritan
  2023: 0000000000524cf8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Saurashtra
  2024: 0000000000524d10     8 OBJECT  GLOBAL DEFAULT   10 unicode.Sharada
  2025: 0000000000524d18     8 OBJECT  GLOBAL DEFAULT   10 unicode.Shavian
  2026: 0000000000524d20     8 OBJECT  GLOBAL DEFAULT   10 unicode.Siddham
  2027: 0000000000524d28     8 OBJECT  GLOBAL DEFAULT   10 unicode.SignWriting
  2028: 0000000000524d30     8 OBJECT  GLOBAL DEFAULT   10 unicode.Sinhala
  2029: 0000000000524d58     8 OBJECT  GLOBAL DEFAULT   10 unicode.Sogdian
  2030: 0000000000524d60     8 OBJECT  GLOBAL DEFAULT   10 unicode.Sora_Sompeng
  2031: 0000000000524d68     8 OBJECT  GLOBAL DEFAULT   10 unicode.Soyombo
  2032: 0000000000524d70     8 OBJECT  GLOBAL DEFAULT   10 unicode.Sundanese
  2033: 0000000000524d78     8 OBJECT  GLOBAL DEFAULT   10 unicode.Syloti_Nagri
  2034: 0000000000524d80     8 OBJECT  GLOBAL DEFAULT   10 unicode.Syriac
  2035: 0000000000524d88     8 OBJECT  GLOBAL DEFAULT   10 unicode.Tagalog
  2036: 0000000000524d90     8 OBJECT  GLOBAL DEFAULT   10 unicode.Tagbanwa
  2037: 0000000000524d98     8 OBJECT  GLOBAL DEFAULT   10 unicode.Tai_Le
  2038: 0000000000524da0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Tai_Tham
  2039: 0000000000524da8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Tai_Viet
  2040: 0000000000524db0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Takri
  2041: 0000000000524db8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Tamil
  2042: 0000000000524dc0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Tangut
  2043: 0000000000524dc8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Telugu
  2044: 0000000000524dd8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Thaana
  2045: 0000000000524de0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Thai
  2046: 0000000000524de8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Tibetan
  2047: 0000000000524df0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Tifinagh
  2048: 0000000000524df8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Tirhuta
  2049: 0000000000524e00     8 OBJECT  GLOBAL DEFAULT   10 unicode.Ugaritic
  2050: 0000000000524e10     8 OBJECT  GLOBAL DEFAULT   10 unicode.Vai
  2051: 0000000000524e20     8 OBJECT  GLOBAL DEFAULT   10 unicode.Wancho
  2052: 0000000000524e28     8 OBJECT  GLOBAL DEFAULT   10 unicode.Warang_Citi
  2053: 0000000000524e38     8 OBJECT  GLOBAL DEFAULT   10 unicode.Yezidi
  2054: 0000000000524e40     8 OBJECT  GLOBAL DEFAULT   10 unicode.Yi
  2055: 0000000000524e50     8 OBJECT  GLOBAL DEFAULT   10 unicode.Zanabaza[...]
  2056: 000000000052c040     8 OBJECT  GLOBAL DEFAULT   11 unicode.Properties
  2057: 0000000000524760     8 OBJECT  GLOBAL DEFAULT   10 unicode.ASCII_He[...]
  2058: 00000000005247c8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Bidi_Control
  2059: 0000000000524880     8 OBJECT  GLOBAL DEFAULT   10 unicode.Dash
  2060: 0000000000524888     8 OBJECT  GLOBAL DEFAULT   10 unicode.Deprecated
  2061: 00000000005248a0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Diacritic
  2062: 00000000005248e0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Extender
  2063: 0000000000524958     8 OBJECT  GLOBAL DEFAULT   10 unicode.Hex_Digit
  2064: 0000000000524968     8 OBJECT  GLOBAL DEFAULT   10 unicode.Hyphen
  2065: 0000000000524970     8 OBJECT  GLOBAL DEFAULT   10 unicode.IDS_Bina[...]
  2066: 0000000000524978     8 OBJECT  GLOBAL DEFAULT   10 unicode.IDS_Trin[...]
  2067: 0000000000524980     8 OBJECT  GLOBAL DEFAULT   10 unicode.Ideographic
  2068: 00000000005249b0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Join_Control
  2069: 0000000000524a58     8 OBJECT  GLOBAL DEFAULT   10 unicode.Logical_[...]
  2070: 0000000000524b78     8 OBJECT  GLOBAL DEFAULT   10 unicode.Nonchara[...]
  2071: 0000000000524bf8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Other_Al[...]
  2072: 0000000000524c00     8 OBJECT  GLOBAL DEFAULT   10 unicode.Other_De[...]
  2073: 0000000000524c08     8 OBJECT  GLOBAL DEFAULT   10 unicode.Other_Gr[...]
  2074: 0000000000524c10     8 OBJECT  GLOBAL DEFAULT   10 unicode.Other_ID[...]
  2075: 0000000000524c18     8 OBJECT  GLOBAL DEFAULT   10 unicode.Other_ID[...]
  2076: 0000000000524c20     8 OBJECT  GLOBAL DEFAULT   10 unicode.Other_Lo[...]
  2077: 0000000000524c28     8 OBJECT  GLOBAL DEFAULT   10 unicode.Other_Math
  2078: 0000000000524c30     8 OBJECT  GLOBAL DEFAULT   10 unicode.Other_Up[...]
  2079: 0000000000524c50     8 OBJECT  GLOBAL DEFAULT   10 unicode.Pattern_[...]
  2080: 0000000000524c58     8 OBJECT  GLOBAL DEFAULT   10 unicode.Pattern_[...]
  2081: 0000000000524ca8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Prepende[...]
  2082: 0000000000524cc0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Quotatio[...]
  2083: 0000000000524cc8     8 OBJECT  GLOBAL DEFAULT   10 unicode.Radical
  2084: 0000000000524cd0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Regional[...]
  2085: 0000000000524d08     8 OBJECT  GLOBAL DEFAULT   10 unicode.Sentence[...]
  2086: 0000000000524d50     8 OBJECT  GLOBAL DEFAULT   10 unicode.Soft_Dotted
  2087: 0000000000524dd0     8 OBJECT  GLOBAL DEFAULT   10 unicode.Terminal[...]
  2088: 0000000000524e08     8 OBJECT  GLOBAL DEFAULT   10 unicode.Unified_[...]
  2089: 0000000000524e18     8 OBJECT  GLOBAL DEFAULT   10 unicode.Variatio[...]
  2090: 0000000000524e30     8 OBJECT  GLOBAL DEFAULT   10 unicode.White_Space
  2091: 000000000052c030     8 OBJECT  GLOBAL DEFAULT   11 unicode.FoldCategory
  2092: 0000000000524e88     8 OBJECT  GLOBAL DEFAULT   10 unicode.foldL
  2093: 0000000000524e90     8 OBJECT  GLOBAL DEFAULT   10 unicode.foldLl
  2094: 0000000000524e98     8 OBJECT  GLOBAL DEFAULT   10 unicode.foldLt
  2095: 0000000000524ea0     8 OBJECT  GLOBAL DEFAULT   10 unicode.foldLu
  2096: 0000000000524ea8     8 OBJECT  GLOBAL DEFAULT   10 unicode.foldM
  2097: 0000000000524eb0     8 OBJECT  GLOBAL DEFAULT   10 unicode.foldMn
  2098: 000000000052c038     8 OBJECT  GLOBAL DEFAULT   11 unicode.FoldScript
  2099: 0000000000524e70     8 OBJECT  GLOBAL DEFAULT   10 unicode.foldCommon
  2100: 0000000000524e78     8 OBJECT  GLOBAL DEFAULT   10 unicode.foldGreek
  2101: 0000000000524e80     8 OBJECT  GLOBAL DEFAULT   10 unicode.foldInherited
  2102: 0000000000559e60     1 OBJECT  GLOBAL DEFAULT   12 internal/cpu.Deb[...]
  2103: 0000000000514208     8 OBJECT  GLOBAL DEFAULT    9 internal/cpu.Cac[...]
  2104: 000000000055a4a0   145 OBJECT  GLOBAL DEFAULT   12 internal/cpu.X86
  2105: 000000000052c350    24 OBJECT  GLOBAL DEFAULT   11 internal/cpu.options
  2106: 0000000000559e78     4 OBJECT  GLOBAL DEFAULT   12 internal/cpu.max[...]
  2107: 0000000000515700    40 OBJECT  GLOBAL DEFAULT    9 internal/oserror[...]
  2108: 000000000052c0a0    16 OBJECT  GLOBAL DEFAULT   11 internal/oserror[...]
  2109: 000000000052c0c0    16 OBJECT  GLOBAL DEFAULT   11 internal/oserror[...]
  2110: 000000000052c090    16 OBJECT  GLOBAL DEFAULT   11 internal/oserror[...]
  2111: 000000000052c0b0    16 OBJECT  GLOBAL DEFAULT   11 internal/oserror[...]
  2112: 000000000052c080    16 OBJECT  GLOBAL DEFAULT   11 internal/oserror[...]
  2113: 0000000000515a80    48 OBJECT  GLOBAL DEFAULT    9 path..inittask
  2114: 000000000052c2b0    16 OBJECT  GLOBAL DEFAULT   11 path.ErrBadPattern
  2115: 000000000055a058     8 OBJECT  GLOBAL DEFAULT   12 runtime/internal[...]
  2116: 00000000004b7658    32 OBJECT  GLOBAL DEFAULT    2 go:itab.*os.File[...]
  2117: 00000000004b7708    56 OBJECT  GLOBAL DEFAULT    2 go:itab.*fmt.pp,[...]
  2118: 00000000004b8420   272 OBJECT  GLOBAL DEFAULT    2 go:itab.*reflect[...]
  2119: 00000000004b75f8    32 OBJECT  GLOBAL DEFAULT    2 go:itab.*errors.[...]
  2120: 00000000004b7d48   112 OBJECT  GLOBAL DEFAULT    2 go:itab.*interna[...]
  2121: 00000000004b76d8    48 OBJECT  GLOBAL DEFAULT    2 go:itab.*interna[...]
  2122: 00000000004b7638    32 OBJECT  GLOBAL DEFAULT    2 go:itab.*io/fs.P[...]
  2123: 00000000004b76b8    32 OBJECT  GLOBAL DEFAULT    2 go:itab.syscall.[...]
  2124: 00000000004b7698    32 OBJECT  GLOBAL DEFAULT    2 go:itab.runtime.[...]
  2125: 000000000052bee8     8 OBJECT  GLOBAL DEFAULT   11 _cgo_init
  2126: 000000000052bf10     8 OBJECT  GLOBAL DEFAULT   11 _cgo_thread_start
  2127: 000000000052bf00     8 OBJECT  GLOBAL DEFAULT   11 _cgo_notify_runt[...]
  2128: 000000000052bee0     8 OBJECT  GLOBAL DEFAULT   11 _cgo_callers
  2129: 000000000052bf18     8 OBJECT  GLOBAL DEFAULT   11 _cgo_yield
  2130: 000000000052bef0     8 OBJECT  GLOBAL DEFAULT   11 _cgo_mmap
  2131: 000000000052bef8     8 OBJECT  GLOBAL DEFAULT   11 _cgo_munmap
  2132: 000000000052bf08     8 OBJECT  GLOBAL DEFAULT   11 _cgo_sigaction
  2133: 00000000004b7148     8 OBJECT  GLOBAL DEFAULT    2 runtime.mainPC
  2134: 00000000004b7678    32 OBJECT  GLOBAL DEFAULT    2 go:itab.internal[...]
  2135: 00000000004b7618    32 OBJECT  GLOBAL DEFAULT    2 go:itab.*interna[...]
  2136: 00000000004b7158     9 OBJECT  GLOBAL DEFAULT    2 runtime.buildVer[...]
  2137: 00000000004b8120   251 OBJECT  GLOBAL DEFAULT    2 runtime.modinfo.str
  2138: 0000000000514000   304 OBJECT  GLOBAL DEFAULT    8 go:buildinfo
  2139: 00000000004b7128     8 OBJECT  GLOBAL DEFAULT    2 go:buildinfo.ref
  2140: 0000000000482000     0 OBJECT  GLOBAL DEFAULT    2 type:*
  2141: 00000000004b75e0    24 OBJECT  GLOBAL DEFAULT    2 runtime.textsect[...]

No version information found in this file.

Displaying notes found in: .note.go.buildid
  Owner                Data size    Description
  Go                   0x00000053   GO BUILDID
   description data: 34 4e 38 73 50 2d 55 53 6c 30 4c 4f 64 48 7a 73 50 58 50 72 2f 59 32 55 77 73 71 4e 78 78 75 76 42 71 72 65 68 76 5f 2d 4d 2f 54 50 35 63 77 4c 77 6c 71 35 48 4d 5a 46 66 71 7a 4e 50 34 2f 6f 38 58 49 55 4a 65 4d 33 58 46 44 4b 74 4c 74 56 44 6a 68
```

Any go binary includes `cgo` related symbols?

```
  2125: 000000000052bee8     8 OBJECT  GLOBAL DEFAULT   11 _cgo_init
  2126: 000000000052bf10     8 OBJECT  GLOBAL DEFAULT   11 _cgo_thread_start
  2127: 000000000052bf00     8 OBJECT  GLOBAL DEFAULT   11 _cgo_notify_runt[...]
  2128: 000000000052bee0     8 OBJECT  GLOBAL DEFAULT   11 _cgo_callers
  2129: 000000000052bf18     8 OBJECT  GLOBAL DEFAULT   11 _cgo_yield
  2130: 000000000052bef0     8 OBJECT  GLOBAL DEFAULT   11 _cgo_mmap
  2131: 000000000052bef8     8 OBJECT  GLOBAL DEFAULT   11 _cgo_munmap
  2132: 000000000052bf08     8 OBJECT  GLOBAL DEFAULT   11 _cgo_sigaction
```

More on the various sections and section types:

* [https://refspecs.linuxbase.org/elf/gabi4+/ch4.sheader.html](https://refspecs.linuxbase.org/elf/gabi4+/ch4.sheader.html)
* [https://www.linuxjournal.com/article/1060](https://www.linuxjournal.com/article/1060)
