package main

/*

#include <stdlib.h>

typedef struct {
	int a, b, c;
	char *str;
} someStruct;

void processStruct(someStruct *s) {}

*/
import "C"
import (
	"unsafe"
)

func ProcessStruct() {
	s := (*C.someStruct)(C.calloc(1, C.sizeof_someStruct)) // T: *main._Ctype_struct___0
	defer C.free(unsafe.Pointer(s))

	s.str = C.CString(str)
	defer C.free(unsafe.Pointer(s.str))

	s := (*C.someStruct)(someCMemory)
	C.processStruct(s)
}

func main() {
	ProcessStruct()
}
