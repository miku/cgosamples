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
