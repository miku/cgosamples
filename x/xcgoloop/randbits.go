package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	bw := bufio.NewWriter(os.Stdout)
	defer bw.Flush()
	for i := 0; i < 1_000_000_000; i++ {
		if rand.Float64() > 0.5 {
			// fmt.Println(0)
			fmt.Fprintln(bw, 0)
		} else {
			// fmt.Println(1)   // 2M
			fmt.Fprintln(bw, 1) // 11M
		}
	}
}
