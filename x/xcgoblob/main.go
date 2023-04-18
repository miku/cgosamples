package main

/*

#include <stdlib.h>

void * blob() {
	void * b = malloc(8);
	return b;
}

*/
import "C"
import (
	"log"
	"unsafe"
)

func main() {
	b := C.GoBytes(unsafe.Pointer(C.blob()), C.int(32))
	// $ go run main.go
	// 2023/04/18 16:57:07 32 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 209 7 2 0 0 0 0 0] []uint8

	log.Printf("%d %v %T", len(b), b, b) // uint8
	b[31] = 123
	log.Printf("%d %v %T", len(b), b, b) // uint8
}
