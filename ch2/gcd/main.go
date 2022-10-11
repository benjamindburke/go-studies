package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Fprintf(os.Stderr, "exactly 2 arguments required, received: %s", strings.Join(os.Args[1:], " "))
		os.Exit(1)
	}
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "gcd: %v", err)
		os.Exit(1)
	}
	y, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "gcd: %v", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "gcd(%d, %d) = %d\n", x, y, gcd(x, y))
}

func gcd(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
