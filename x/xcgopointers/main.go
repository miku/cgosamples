package main

/*

#include <stdio.h>
#include <stdlib.h>

void print_i(void *v) {
	printf("%d\n", *(int*)v);
	// free(v); // free(): invalid pointer
}


*/
import "C"
import (
	"log"
	"unsafe"
)

func main() {
	i := 42
	p := unsafe.Pointer(&i) // Go pointer
	C.print_i(p)
	log.Println(i)
}
