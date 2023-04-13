package main

/*
typedef struct {
  int a, b, c;
} someStruct;

someStruct* someReference;

void storeReference(someStruct* s) {
  someReference = s;
}

int retrieveReference() {
    return someReference->a;
}
*/
import "C"
import "runtime"

func main() {
	passSomeMemory()
	runtime.GC()
	// It's better now: https://speakerdeck.com/filosottile/from-cgo-back-to-go-gophercon-2016?slide=22
	println(getSomeMemory()) // 1
}

func passSomeMemory() {
	someGoMemory := &C.someStruct{
		a: 1, b: 2, c: 3,
	}
	C.storeReference(someGoMemory)
}

func getSomeMemory() int {
	return int(C.retrieveReference())
}
