// $ ./main | jq -r '.tid' | sort | uniq -c
//
//	 70 1059689
//	134 1059694
//	207 1059695
//	138 1059696
//	109 1059698
//	150 1059699
//	 50 1059700
//	143 1059701
//
// So far, so good, but why do we have more than NumCPU threads when we increase N?
//
// # For 10K
//
// $ for i in $(seq 100); do ./main -n 10000 | jq -r '.tid' | sort | uniq -c | wc -l; done | sort | uniq -c
//
//	99 10
//	 1 11
//
// For 100K, we cannot get beyond 10 threads, seemingly.
// For 500K, -- "" --
// For 1M, we've seen: 11, 10, 10, 10
package main

/*

#include <sys/types.h>
#include <unistd.h>
#include <sys/syscall.h>

pid_t gettid() {
	pid_t x = syscall(__NR_gettid);
	return x;
}

*/
import "C"
import (
	"flag"
	"fmt"
	"os"
	"sync"
)

var N = flag.Int("n", 1000, "number of goroutines to start")

func main() {
	flag.Parse()
	var wg sync.WaitGroup
	for i := 0; i < *N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("{\"i\": %d, \"tid\": %d, \"pid\": %d}\n",
				i, C.gettid(), os.Getpid())
		}(i)
	}
	wg.Wait()
	fmt.Printf("{\"tid\": %d, \"pid\": %d}\n", C.gettid(), os.Getpid())
}
