// Just including "rect" won't work, as Go uses a C compiler.
//
// Run with:
//
//	LD_LIBRARY_PATH=. ./main
package main

/*

#cgo LDFLAGS: -L . -lrect -lstdc++
#include <rect.h>

*/
import "C"
import "fmt"

func main() {
	r := C.rect_new(C.int(2), C.int(3))
	fmt.Println(r)
	fmt.Println(C.rect_area(r)) // 6
}
