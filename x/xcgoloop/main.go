// Two random walks, one in C, one in Go; Go about 3-5 slower than C.
// Cf. https://go.dev/play/p/lYRJFW_iby_c
//
//  C               1               1 34.095µs
// GO               1               1 78.826µs
//  C            1000               8 40.951µs
// GO            1000             -26 123.964µs
//  C         1000000            -830 25.799065ms
// GO         1000000            -482 85.781652ms
//  C      1000000000           43980 6.811447918s
// GO      1000000000          -68244 27.529384394s
//

package main

/*

#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

uint32_t lcg_parkmiller(uint32_t* state) {
  return *state = (uint64_t)*state * 48271 % 0x7fffffff;
}

int_fast8_t flip(uint32_t r) {
  return r & 1;
}

// rwalk implements a random walk on n steps, starting at 0.
uint64_t rwalk(uint64_t n) {
  srand(time(NULL));
  uint32_t state = rand();
  int64_t i = 0, j = 0;
  for (i = 0; i < n; i++) {
    if (flip(lcg_parkmiller(&state)) == 0) {
      j--;
    } else {
      j++;
    }
  }
  return j;
}

*/
import "C"
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sizes := []int64{
		1,
		1_000,
		1_000_000,
		1_000_000_000,
		// 10_000_000_000,
	}
	for _, sz := range sizes {
		var (
			csz     = C.ulong(sz)
			started = time.Now()
			result  = int(C.rwalk(csz))
			done    = time.Since(started)
		)
		fmt.Printf(" C % 15d % 15d %v\n", sz, result, done)

		started = time.Now()
		result = rwalk(sz)
		done = time.Since(started)
		fmt.Printf("GO % 15d % 15d %v\n", sz, result, done)
	}
}

func rwalk(n int64) int {
	var i, j int64
	for i = 0; i < n; i++ {
		if rand.Float32() > 0.5 {
			j++
		} else {
			j--
		}
	}
	return int(j)
}

// parkmiller, https://go.dev/play/p/lYRJFW_iby_c, do we care about the uint64
// https://en.wikipedia.org/wiki/Lehmer_random_number_generator#Sample_C99_code
func parkmiller(state *uint32) uint32 {
	*state = *state * 48271 % 0x7fffffff
	return *state
}
