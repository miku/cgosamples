// Note, run with LD_LIBRARY_PATH=. ...
package main

/*

#cgo CFLAGS: -g
#cgo LDFLAGS: -Wl,-R -Wl,${SRCDIR} -L${SRCDIR} -lm -lggml

#include "ggml.h"

*/
import "C"
import "fmt"

func main() {
	params := C.struct_ggml_init_params{
		mem_size:   16 * 1024 * 1024,
		mem_buffer: nil,
	}
	fmt.Printf("%v %T\n", params, params) // main._Ctype_struct_ggml_init_params

	C.ggml_time_init()

	// Initialize, will allocate memory.
	ctx := C.ggml_init(params)
	fmt.Printf("%v %T\n", ctx, ctx) // main._Ctype_struct_ggml_init_params
}
