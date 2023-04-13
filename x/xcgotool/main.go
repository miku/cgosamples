package main

/*

#include <stdio.h>
#include <stdlib.h>

void hello(char* s)
{
    printf("hello %s\n", s);
}

*/
import "C"
import "unsafe"

func main() {
	cs := C.CString("gopher")
	C.hello(cs)
	C.free(unsafe.Pointer(cs))
}
