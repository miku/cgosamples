package main

/*

#include <stdio.h>
#include <stdlib.h>

void print_i(void *v) {
	printf("*v = %d %p (C)\n", *(int*)v, v);
	int j = 84;

	// Try to change the value of v, will not be visible to Go.
	v = &j;
	printf("&j = %d %p (C)\n", j, v);
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	i := 42
	p := unsafe.Pointer(&i) // Go pointer
	fmt.Printf("&p = %p (Go)\n", p)
	C.print_i(p)
	fmt.Printf("back: %v (Go)\n", i)
}
