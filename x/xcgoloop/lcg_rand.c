// This program generates 1B random bits, printed out as 0 or 1 per line. Uses
// 2772 bytes of virtual memory. About 11M lines/s.
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

uint32_t lcg_parkmiller(uint32_t* state) {
  return *state = (uint64_t)*state * 48271 % 0x7fffffff;
}

uint32_t lcg_rand(uint32_t* state) {
  return *state = (uint64_t)*state * 279470273u % 0xfffffffb;
}

int_fast8_t flip(uint32_t r) {
  return r & 1;
}

int main() {
  srand(time(NULL));
  int i;
  uint32_t state = rand();
  for (i = 0; i < 1000000000; i++) {
    // printf("%d\n", flip(lcg_rand(&state)));
    printf("%d\n", flip(lcg_parkmiller(&state)));
  }
}
