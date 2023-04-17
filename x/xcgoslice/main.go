package main

/*

#include <stdio.h>

void f(int vs[], int c) {
	int i;
	for (i = 0; i < c; i++) {
		printf("[% 2d/% 2d] %d\n", i, c, vs[i]);
	}
}

*/
import "C"
import (
	"unsafe"
)

func main() {
	s := make([]int, 10)
	for i := 0; i < len(s); i++ {
		s[i] = 8
	}
	C.f((*C.int)(unsafe.Pointer(&s[0])), C.int(len(s)))
}
