package main

/*

#include <stdlib.h>

typedef struct {
	int a, b, c;
} someStruct;

void processStruct(someStruct *s) {}

*/
import "C"

func ProcessStruct() {
	someCMemory := C.calloc(1, C.sizeof_someStruct)
	defer C.free(someCMemory)
	s := (*C.someStruct)(someCMemory)
	C.processStruct(s)
}

func main() {
	ProcessStruct()
}
