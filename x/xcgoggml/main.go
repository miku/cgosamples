// Note, run with LD_LIBRARY_PATH=. ...
package main

/*

#cgo CFLAGS: -g
#cgo LDFLAGS: -Wl,-R -Wl,${SRCDIR} -L${SRCDIR} -lm -lggml

#include "ggml.h"

*/
import "C"
import (
	"log"
)

func main() {
	// For example, here we define the function: f(x) = a*x^2 + b
	params := C.struct_ggml_init_params{
		mem_size:   16 * 1024 * 1024,
		mem_buffer: nil,
	}
	C.ggml_time_init()

	ctx := C.ggml_init(params)
	x := C.ggml_new_tensor_1d(ctx, C.GGML_TYPE_F32, C.int(1))

	C.ggml_set_param(ctx, x)
	var (
		a  = C.ggml_new_tensor_1d(ctx, C.GGML_TYPE_F32, C.int(1))
		b  = C.ggml_new_tensor_1d(ctx, C.GGML_TYPE_F32, C.int(1))
		x2 = C.ggml_mul(ctx, x, x)
		f  = C.ggml_add(ctx, C.ggml_mul(ctx, a, x2), b)
	)

	log.Printf("%v %T\n", f, f) // *main._Ctype_struct_ggml_tensor

	gf := C.ggml_build_forward(f)
	C.ggml_set_f32(x, C.float(2.0))
	C.ggml_set_f32(a, C.float(3.0))
	C.ggml_set_f32(b, C.float(4.0))

	// 3.0 * 2.0^2 + 4.0

	C.ggml_graph_compute(ctx, &gf)
	result := C.ggml_get_f32_1d(f, C.int(0))
	log.Printf("%v", result)
}
