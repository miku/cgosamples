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

func main() {
	s := make([]int, 10) // TODO: need to change something here
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
	// Cannot do this.
	// C.f((*C.int)(&s[0]), C.int(len(s))) // invalid conversion

	// One would might expect this to work?  https://pkg.go.dev/reflect#SliceHeader
	// C.f((*C.int)(unsafe.Pointer(&s)), C.int(len(s)))

	// Actually, we want to point to the first element.
	// C.f((*C.int)(unsafe.Pointer(&s[0])), C.int(len(s)))
}
