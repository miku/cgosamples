package main

/*
#include <rect.h>

*/
import "C"
import "fmt"

func main() {
	r := C.Rectangle(2, 3)
	fmt.Println(r)
}
