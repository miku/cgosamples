package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 1_000_000_000; i++ {
		if rand.Float64() > 0.5 {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
	}
}
