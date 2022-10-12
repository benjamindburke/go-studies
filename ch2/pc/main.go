// Pc computes the number of set bits (population count) in numbers 0-255.
package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/benjamindburke/go-studies/ch2/popcount"
)

func main() {
	for _, arg := range os.Args[1:] {
		num, err := strconv.ParseUint(arg, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "pc: %v", err)
			os.Exit(1)
		}

		fmt.Printf("pc(%d) = %d\n", num, popcount.PopCount(num))
	}
}